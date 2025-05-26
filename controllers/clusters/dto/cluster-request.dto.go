package dto

type ClusterRequest struct {
	MasterIP    string `json:"master_ip" validate:"required"`
	AgentPort   string `json:"agent_port" validate:"required"`
	Description string `json:"description" validate:"required"`
}
