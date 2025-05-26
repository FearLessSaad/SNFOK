package repository

import (
	"encoding/json"
	"fmt"

	"github.com/FearLessSaad/SNFOK/constants/agent_consts"
	"github.com/FearLessSaad/SNFOK/constants/message"
	"github.com/FearLessSaad/SNFOK/constants/response"
	"github.com/FearLessSaad/SNFOK/controllers/clusters/persistance"
	"github.com/FearLessSaad/SNFOK/tooling/global_dto"
	"github.com/FearLessSaad/SNFOK/tooling/httpclient"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
	"github.com/gofiber/fiber"
)

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
		Status:  "error",
		Message: message.SNFOK_AGENT_IS_NOT_ACCESSABLE,
		Data:    &res_data,
		Meta: &global_dto.Meta{
			Code: response.SNFOK_AGENT_IS_NOT_ACCESSABLE,
		},
	}, fiber.StatusOK

}
