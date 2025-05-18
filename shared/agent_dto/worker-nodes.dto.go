package agent_dto

// NodeInfo holds the details of a worker node
type WorkerNodeInfo struct {
	Name       string `json:"name"`
	Hostname   string `json:"hostname"`
	InternalIP string `json:"internal_ip"`
	ExternalIP string `json:"external_ip"`
}
