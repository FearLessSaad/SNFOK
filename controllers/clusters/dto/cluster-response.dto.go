package dto

type ClusterResponse struct {
	ID          string `json:"id"`
	ClusterName string `json:"cluster_name"`
	MasterIP    string `json:"master_ip"`
	AgentPort   int    `json:"agent_port"`
	Description string `json:"description"`
}
