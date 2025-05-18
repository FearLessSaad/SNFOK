package agent_dto

// NamespaceResources holds details of resources running in a namespace
type NamespaceResources struct {
	Pods         []PodInfo         `json:"pods"`
	Deployments  []DeploymentInfo  `json:"deployments"`
	Services     []ServiceInfo     `json:"services"`
	StatefulSets []StatefulSetInfo `json:"stateful_sets"`
	DaemonSets   []DaemonSetInfo   `json:"daemon_sets"`
}

// PodInfo contains pod details
type PodInfo struct {
	Name       string          `json:"name"`
	Phase      string          `json:"phase"`
	Containers []ContainerInfo `json:"containers"`
}

// ContainerInfo contains container details
type ContainerInfo struct {
	Name   string `json:"name"`
	Image  string `json:"image"`
	Status string `json:"status"` // e.g., Running, Waiting, Terminated
}

// DeploymentInfo contains deployment details
type DeploymentInfo struct {
	Name              string `json:"name"`
	Replicas          int32  `json:"replicas"`
	AvailableReplicas int32  `json:"available_replicas"`
}

// ServiceInfo contains service details
type ServiceInfo struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	ClusterIP string `json:"cluster_ip"`
}

// StatefulSetInfo contains StatefulSet details
type StatefulSetInfo struct {
	Name              string `json:"name"`
	Replicas          int32  `json:"replicas"`
	AvailableReplicas int32  `json:"available_replicas"`
}

// DaemonSetInfo contains DaemonSet details
type DaemonSetInfo struct {
	Name                   string `json:"name"`
	DesiredNumberScheduled int32  `json:"desired_number_scheduled"`
	NumberReady            int32  `json:"number_ready"`
}
