package dto

type CountAllStats struct {
	RunningPods         int `json:"running_pods"`
	Alerts              int `json:"alerts"`
	ImplimentedPolicies int `json:"implimented_policies"`
}
