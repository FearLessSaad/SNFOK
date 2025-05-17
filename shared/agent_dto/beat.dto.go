package agent_dto

type Health_K8s struct {
	ClusterName      string `json:"cluster_name"`
	KubernetesStatus bool   `json:"kubernetes_status"`
	HealthyNodes     int    `json:"healthy_nodes"`
}

type Health_System struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type HealthResponse struct {
	K8sInfo    Health_K8s    `json:"k8s"`
	SystemInfo Health_System `json:"system"`
}
