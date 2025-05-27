package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/FearLessSaad/SNFOK/constants/agent_consts"
	"github.com/FearLessSaad/SNFOK/constants/message"
	"github.com/FearLessSaad/SNFOK/constants/response"
	"github.com/FearLessSaad/SNFOK/controllers/policies/persistance"
	"github.com/FearLessSaad/SNFOK/db/models/k8s"
	"github.com/FearLessSaad/SNFOK/shared/agent_dto"
	"github.com/FearLessSaad/SNFOK/tooling/global_dto"
	"github.com/FearLessSaad/SNFOK/tooling/httpclient"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
	"github.com/gofiber/fiber"

	cluster "github.com/FearLessSaad/SNFOK/controllers/clusters/persistance"
)

type DeployedPolicyResponse struct {
	PolicyPath string `json:"policy_path"`
}

func DeployPolicy(namespace string, app_label string, policy string) (global_dto.Response[DeployedPolicyResponse], int) {

	get_policy, _ := persistance.GetPlicysById(policy)
	get_Master, _ := cluster.GetAllClusters()
	ip := get_Master[0].MasterIP
	port := get_Master[0].AgentPort

	client := httpclient.NewClient(0)

	res, err := client.Post("http://"+ip+":"+fmt.Sprintf("%d", port)+agent_consts.POLICIES_DEPLOY_POLICY, agent_dto.DeployPolicy{
		Namespace: namespace,
		AppLabel:  app_label,
		FilePath:  get_policy.PolicyFilePath,
	}, map[string]string{
		"Content-Type": "application/json",
	})
	if err != nil {
		logger.Log(logger.DEBUG, "HTTP Request Error", logger.Field{Key: "error", Value: err.Error()})
		return global_dto.Response[DeployedPolicyResponse]{
			Status:  "error",
			Message: message.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			},
		}, fiber.StatusBadRequest
	}

	var res_data DeployedPolicyResponse
	if err := json.Unmarshal(res.Body, &res_data); err != nil {
		logger.Log(logger.DEBUG, "Unmarshal Response", logger.Field{Key: "error", Value: err.Error()})
		return global_dto.Response[DeployedPolicyResponse]{
			Status:  "error",
			Message: message.SOMETING_WRONG,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.EXECUTION_ERROR,
			},
		}, fiber.StatusInternalServerError
	}

	i_policy := k8s.ImplimentedPolicies{
		PolicyTitle:    get_policy.PolicyTitle,
		Description:    get_policy.Description,
		AppLabel:       app_label,
		Namespace:      namespace,
		PolicyFilePath: res_data.PolicyPath,
		AuditFields: k8s.AuditFields{
			CreatedBy: "SNFOK:USER",
			CreatedAt: time.Now(),
		},
	}

	persistance.CreateImplimentedPolicy(i_policy)

	return global_dto.Response[DeployedPolicyResponse]{
		Status:  "error",
		Message: message.SOMETING_WRONG,
		Data:    &res_data,
		Meta: &global_dto.Meta{
			Code: response.POLICY_DEPLOYED,
		},
	}, fiber.StatusOK

}

func DeletePolicy(id string) (global_dto.Response[[]string], int) {
	policy, _ := persistance.GetImplimentedPolicyById(id)
	cluster, _ := cluster.GetAllClusters()

	ip := cluster[0].MasterIP
	port := cluster[0].AgentPort
	client := httpclient.NewClient(0)

	_, err := client.Get("http://"+ip+":"+fmt.Sprintf("%d", port)+agent_consts.DELETE_TETRAGON_POLICY(policy.PolicyFilePath), map[string]string{})
	if err != nil {
		logger.Log(logger.DEBUG, "HTTP Request Error", logger.Field{Key: "error", Value: err.Error()})
		return global_dto.Response[[]string]{
			Status:  "error",
			Message: message.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			},
		}, fiber.StatusInternalServerError
	}

	err = persistance.DeleteImplimentedPolicyById(policy.ID)

	if err != nil {
		return global_dto.Response[[]string]{
			Status:  "error",
			Message: message.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			},
		}, fiber.StatusInternalServerError
	}

	return global_dto.Response[[]string]{
		Status:  "success",
		Message: "Policy Is Deleted Successfully!",
		Data:    nil,
		Meta: &global_dto.Meta{
			Code: response.NAMESPACES_RESPONSE,
		},
	}, fiber.StatusOK

}

func GetAllImplimentedPolicies() (global_dto.Response[[]k8s.ImplimentedPolicies], int) {

	policies, _ := persistance.GetAllImplimentedPolicies()

	return global_dto.Response[[]k8s.ImplimentedPolicies]{
		Status:  "success",
		Message: "",
		Data:    &policies,
		Meta: &global_dto.Meta{
			Code: 5,
		},
	}, fiber.StatusOK

}
