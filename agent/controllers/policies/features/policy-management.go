package features

import (
	"os"
	"os/exec"
	"path/filepath"

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

func DeletePolicy(policy_path string) (string, error) {
	applied_policies_dir := os.Getenv("APPLIED_POLICIES_DIR")

	path := filepath.Join(applied_policies_dir, policy_path)
	cmd := exec.Command("kubectl", "delete", "-f", path)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	err = os.Remove(path)
	if err != nil {
		return err.Error(), err
	}

	return policy_path, nil
}
