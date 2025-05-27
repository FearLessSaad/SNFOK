package agent_consts

import "fmt"

const (
	HEALTH_BEAT_PATH                  = "/api/health/beat"
	HEALTH_GET_INTO_PATH              = "/api/health/get"
	KUBERNETES_GET_ALL_NAMESPACES     = "/api/kubernetes/namespaces/all"
	KUBERNETES_COUNT_ALL_RUNNING_PODS = "/api/kubernetes/count/pods"
	GET_ALL_APP_LABELS                = "/api/kubernetes/get/all/labels"
)

func GET_ALL_NAMESPACE_RESOURCES(namespace string) string {
	return fmt.Sprintf("/api/kubernetes/namespaces/%s/resources", namespace)
}

const (
	POLICIES_DEPLOY_POLICY = "/api/policies/deplye/policy"
)

// Ploicies Template
const (
	POLICY_NAMESPACE_TEMPLATE = "{{.Namespace}}"
	POLICY_APP_LABEL_TEMPLATE = "{{.AppLabel}}"
	POLICY_ID_TEMPLATE        = "{{.PolicyID}}"
)
