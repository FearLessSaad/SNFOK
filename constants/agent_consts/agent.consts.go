package agent_consts

const (
	HEALTH_BEAT_PATH     = "/api/health/beat"
	HEALTH_GET_INTO_PATH = "/api/health/get"
)

// Ploicies Template
const (
	POLICY_NAMESPACE_TEMPLATE = "{{.Namespace}}"
	POLICY_APP_LABEL_TEMPLATE = "{{.AppLabel}}"
	POLICY_ID_TEMPLATE        = "{{.PolicyID}}"
)
