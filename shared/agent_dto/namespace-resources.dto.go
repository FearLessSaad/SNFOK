package agent_dto

// NamespaceResources holds details of resources running in a namespace
type NamespaceResources struct {
	Pods         []PodInfo         `json:"pods,omitempty"`
	Deployments  []DeploymentInfo  `json:"deployments,omitempty"`
	Services     []ServiceInfo     `json:"services,omitempty"`
	StatefulSets []StatefulSetInfo `json:"stateful_sets,omitempty"`
	DaemonSets   []DaemonSetInfo   `json:"daemon_sets,omitempty"`
}

// PodInfo contains pod details
type PodInfo struct {
	Name       string          `json:"name,omitempty"`
	Phase      string          `json:"phase,omitempty"`
	Containers []ContainerInfo `json:"containers,omitempty"`
}

// ContainerInfo contains container details
type ContainerInfo struct {
	Name   string `json:"name,omitempty"`
	Image  string `json:"image,omitempty"`
	Status string `json:"status,omitempty"` // e.g., Running, Waiting, Terminated
}

// DeploymentInfo contains deployment details
type DeploymentInfo struct {
	Name              string `json:"name,omitempty"`
	Replicas          int32  `json:"replicas,omitempty"`
	AvailableReplicas int32  `json:"available_replicas,omitempty"`
}

// ServiceInfo contains service details
type ServiceInfo struct {
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	ClusterIP string `json:"cluster_ip,omitempty"`
}

// StatefulSetInfo contains StatefulSet details
type StatefulSetInfo struct {
	Name              string `json:"name,omitempty"`
	Replicas          int32  `json:"replicas,omitempty"`
	AvailableReplicas int32  `json:"available_replicas,omitempty"`
}

// DaemonSetInfo contains DaemonSet details
type DaemonSetInfo struct {
	Name                   string `json:"name,omitempty"`
	DesiredNumberScheduled int32  `json:"desired_number_scheduled,omitempty"`
	NumberReady            int32  `json:"number_ready,omitempty"`
}
