package agent_consts

const (
	HEALTH_BEAT_PATH                  = "/api/health/beat"
	HEALTH_GET_INTO_PATH              = "/api/health/get"
	KUBERNETES_GET_ALL_NAMESPACES     = "/api/kubernetes/namespaces/all"
	KUBERNETES_COUNT_ALL_RUNNING_PODS = "/api/kubernetes/count/pods"
)

// Ploicies Template
const (
	POLICY_NAMESPACE_TEMPLATE = "{{.Namespace}}"
	POLICY_APP_LABEL_TEMPLATE = "{{.AppLabel}}"
	POLICY_ID_TEMPLATE        = "{{.PolicyID}}"
)
