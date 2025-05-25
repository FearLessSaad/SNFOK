package agent_consts

const (
	HEALTH_BEAT_PATH     = "/api/health/beat"
	HEALTH_GET_INTO_PATH = "/api/health/get"
)

// Ploicies Template
const (
	NAMESPACE_TEMPLATE   = "{{.Namespace}}"
	APP_LABEL_TEMPLATE   = "{{.AppLabel}}"
	POLICY_NAME_TEMPLATE = "{{.PolicyName}}"
)
