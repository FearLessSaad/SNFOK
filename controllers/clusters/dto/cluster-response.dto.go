package dto

type ClusterResponse struct {
	ID          string
	ClusterName string
	MasterIP    string
	AgentPort   int
	Description string
}
