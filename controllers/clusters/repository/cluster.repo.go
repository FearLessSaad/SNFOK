package repository

import (
	"encoding/json"
	"fmt"

	"github.com/FearLessSaad/SNFOK/constants/agent_consts"
	"github.com/FearLessSaad/SNFOK/constants/message"
	"github.com/FearLessSaad/SNFOK/constants/response"
	"github.com/FearLessSaad/SNFOK/controllers/clusters/dto"
	"github.com/FearLessSaad/SNFOK/controllers/clusters/persistance"
	"github.com/FearLessSaad/SNFOK/db/models/k8s"
	"github.com/FearLessSaad/SNFOK/shared/agent_dto"
	"github.com/FearLessSaad/SNFOK/tooling/global_dto"
	"github.com/FearLessSaad/SNFOK/tooling/httpclient"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
	"github.com/gofiber/fiber/v2"
)

func GetAllClusters() (global_dto.Response[[]dto.ClusterResponse], int) {
	clusters, err := persistance.GetAllClusters()

	if err != nil {
		return global_dto.Response[[]dto.ClusterResponse]{
			Status:  "error",
			Message: message.SOMETING_WRONG,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.EXECUTION_ERROR,
			},
		}, fiber.StatusInternalServerError
	}

	if len(clusters) == 0 {
		return global_dto.Response[[]dto.ClusterResponse]{
			Status:  "success",
			Message: message.NO_REGISTERED_CLUSTER_AVAILABLE,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.NO_CLUSTER_AVAILABLE,
			},
		}, fiber.StatusNoContent
	}

	res := []dto.ClusterResponse{}

	for i := 0; i < len(clusters); i++ {
		res = append(res, dto.ClusterResponse{
			ID:          clusters[i].ID,
			ClusterName: clusters[i].ClusterName,
			MasterIP:    clusters[i].MasterIP,
			AgentPort:   clusters[i].AgentPort,
			Description: clusters[i].Description,
		})
	}

	return global_dto.Response[[]dto.ClusterResponse]{
		Status:  "success",
		Message: message.NO_REGISTERED_CLUSTER_AVAILABLE,
		Data:    &res,
		Meta: &global_dto.Meta{
			Code: response.NO_CLUSTER_AVAILABLE,
		},
	}, fiber.StatusNoContent
}

func AddNewCluster(data dto.ClusterRequest, uid string) (global_dto.Response[string], int) {

	exists, err := persistance.CheckMasterIPExists(data.MasterIP)
	if err != nil {
		return global_dto.Response[string]{
			Status:  "error",
			Message: message.SOMETING_WRONG,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.EXECUTION_ERROR,
			},
		}, fiber.StatusInternalServerError
	}

	if exists {
		return global_dto.Response[string]{
			Status:  "error",
			Message: message.CLUSTER_ALREADY_REGISTERED,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.CLUSTER_ALREADY_REGISTERED,
			},
		}, fiber.StatusOK
	}

	client := httpclient.NewClient(0)

	res, err := client.Get("http://"+data.MasterIP+":"+fmt.Sprintf("%d", data.AgentPort)+agent_consts.HEALTH_GET_INTO_PATH, map[string]string{})
	if err != nil {
		logger.Log(logger.DEBUG, "HTTP Request Error", logger.Field{Key: "error", Value: err.Error()})
		return global_dto.Response[string]{
			Status:  "error",
			Message: message.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.SNFOK_AGENT_IS_NOT_ACCESSABLE,
			},
		}, fiber.StatusOK
	}

	var res_data agent_dto.HealthResponse
	if err := json.Unmarshal(res.Body, &res_data); err != nil {
		logger.Log(logger.DEBUG, "Unmarshal Response", logger.Field{Key: "error", Value: err.Error()})
		return global_dto.Response[string]{
			Status:  "error",
			Message: message.SOMETING_WRONG,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.EXECUTION_ERROR,
			},
		}, fiber.StatusInternalServerError
	}

	cluster := k8s.Clusters{
		ClusterName: res_data.K8sInfo.ClusterName,
		MasterIP:    data.MasterIP,
		AgentPort:   data.AgentPort,
		Description: data.Description,
	}
	cluster.AuditFields.CreatedBy = uid

	err = persistance.CreateCluster(cluster)

	if err != nil {
		return global_dto.Response[string]{
			Status:  "error",
			Message: message.SOMETING_WRONG,
			Data:    nil,
			Meta: &global_dto.Meta{
				Code: response.EXECUTION_ERROR,
			},
		}, fiber.StatusOK
	}

	return global_dto.Response[string]{
		Status:  "success",
		Message: message.CLUSTER_REGISTERED,
		Data:    nil,
		Meta: &global_dto.Meta{
			Code: response.CLUSTER_REGISTERED,
		},
	}, fiber.StatusOK
}
