package repository

import (
	"encoding/json"
	"fmt"

	"github.com/FearLessSaad/SNFOK/constants/agent_consts"
	"github.com/FearLessSaad/SNFOK/constants/message"
	"github.com/FearLessSaad/SNFOK/constants/response"
	"github.com/FearLessSaad/SNFOK/controllers/clusters/persistance"
	"github.com/FearLessSaad/SNFOK/shared/agent_dto"
	"github.com/FearLessSaad/SNFOK/tooling/global_dto"
	"github.com/FearLessSaad/SNFOK/tooling/httpclient"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
	"github.com/gofiber/fiber"
)

type AllResorcesByNamespaces struct {
	Namespace string                       `json:"namespace"`
	Resources agent_dto.NamespaceResources `json:"resources,omitempty"`
}

func GetAllResources() (global_dto.Response[[]AllResorcesByNamespaces], int) {

	clusters, _ := persistance.GetAllClusters()
	ip := clusters[0].MasterIP
	port := clusters[0].AgentPort

	client := httpclient.NewClient(0)

	res, err := client.Get("http://"+ip+":"+fmt.Sprintf("%d", port)+agent_consts.KUBERNETES_GET_ALL_NAMESPACES, map[string]string{})
	if err != nil {
		logger.Log(logger.DEBUG, "HTTP Request Error", logger.Field{Key: "error", Value: err.Error()})
		return global_dto.Response[[]AllResorcesByNamespaces]{
			Status:  "error",
			Message: message.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			},
		}, fiber.StatusOK
	}

	var res_data []string
	if err := json.Unmarshal(res.Body, &res_data); err != nil {
		logger.Log(logger.DEBUG, "Unmarshal Response", logger.Field{Key: "error", Value: err.Error()})
		return global_dto.Response[[]AllResorcesByNamespaces]{
			Status:  "error",
			Message: message.SOMETING_WRONG,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.EXECUTION_ERROR,
			},
		}, fiber.StatusInternalServerError
	}

	resources := []AllResorcesByNamespaces{}

	for i := 0; i < len(res_data); i++ {
		path := agent_consts.GET_ALL_NAMESPACE_RESOURCES(res_data[i])
		res, err := client.Get("http://"+ip+":"+fmt.Sprintf("%d", port)+path, map[string]string{})
		if err != nil {
			logger.Log(logger.DEBUG, "HTTP Request Error", logger.Field{Key: "error", Value: err.Error()})
			return global_dto.Response[[]AllResorcesByNamespaces]{
				Status:  "error",
				Message: message.SNFOK_AGENT_IS_NOT_ACCESSABLE,
				Data:    nil,
				Meta: &global_dto.Meta{
					Code: response.SNFOK_AGENT_IS_NOT_ACCESSABLE,
				},
			}, fiber.StatusOK
		}

		var namespace_resource agent_dto.NamespaceResources
		if err := json.Unmarshal(res.Body, &namespace_resource); err != nil {
			logger.Log(logger.DEBUG, "Unmarshal Response", logger.Field{Key: "error", Value: err.Error()})
			return global_dto.Response[[]AllResorcesByNamespaces]{
				Status:  "error",
				Message: message.SOMETING_WRONG,
				Data:    nil,
				Meta: &global_dto.Meta{
					Code: response.EXECUTION_ERROR,
				},
			}, fiber.StatusInternalServerError
		}

		data := AllResorcesByNamespaces{
			Namespace: res_data[i],
			Resources: namespace_resource,
		}

		resources = append(resources, data)
	}

	return global_dto.Response[[]AllResorcesByNamespaces]{
		Status:  "success",
		Message: "",
		Data:    &resources,
		Meta: &global_dto.Meta{
			Code: response.NAMESPACES_RESPONSE,
		},
	}, fiber.StatusOK

}

func GetAllNamespaces() (global_dto.Response[[]string], int) {

	clusters, _ := persistance.GetAllClusters()
	ip := clusters[0].MasterIP
	port := clusters[0].AgentPort

	client := httpclient.NewClient(0)

	res, err := client.Get("http://"+ip+":"+fmt.Sprintf("%d", port)+agent_consts.KUBERNETES_GET_ALL_NAMESPACES, map[string]string{})
	if err != nil {
		logger.Log(logger.DEBUG, "HTTP Request Error", logger.Field{Key: "error", Value: err.Error()})
		return global_dto.Response[[]string]{
			Status:  "error",
			Message: message.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			},
		}, fiber.StatusOK
	}

	var res_data []string
	if err := json.Unmarshal(res.Body, &res_data); err != nil {
		logger.Log(logger.DEBUG, "Unmarshal Response", logger.Field{Key: "error", Value: err.Error()})
		return global_dto.Response[[]string]{
			Status:  "error",
			Message: message.SOMETING_WRONG,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.EXECUTION_ERROR,
			},
		}, fiber.StatusInternalServerError
	}

	return global_dto.Response[[]string]{
		Status:  "success",
		Message: "",
		Data:    &res_data,
		Meta: &global_dto.Meta{
			Code: response.NAMESPACES_RESPONSE,
		},
	}, fiber.StatusOK

}
