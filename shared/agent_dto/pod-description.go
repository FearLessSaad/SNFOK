package agent_dto

import "time"

// PodDescription defines the structured data for a pod's description
type PodDescription struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
	Created     time.Time
	Status      string
	Reason      string
	Message     string
	StartTime   *time.Time
	NodeName    string
	PodIP       string
	Containers  []ContainerDescription
	Conditions  []PodCondition
	Events      []PodEvent
}

// ContainerDescription holds container details
type ContainerDescription struct {
	Name            string
	Image           string
	ImagePullPolicy string
	State           string
	LastState       string
	Ready           bool
	RestartCount    int32
}

// PodCondition holds pod condition details
type PodCondition struct {
	Type               string
	Status             string
	LastTransitionTime time.Time
}

// PodEvent holds pod event details
type PodEvent struct {
	Type    string
	Reason  string
	Age     string
	Source  string
	Message string
}
