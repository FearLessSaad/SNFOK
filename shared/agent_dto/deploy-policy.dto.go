package agent_dto

type DeployPolicy struct {
	AppLabel  string `json:"app_label"`
	Namespace string `json:"namespace"`
	FilePath  string `json:"file_path"`
}
