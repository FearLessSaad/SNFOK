package features

import (
	"os/exec"

	"github.com/FearLessSaad/SNFOK/agent/tooling/templates"
)

func DeployPolicy(policy_file string, namespace string, app_label string) (string, error) {

	policy_path, _ := templates.GeneratePolicy(policy_file, namespace, app_label)

	cmd := exec.Command("kubectl", "apply", "-f", policy_path)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return policy_path, nil
}
