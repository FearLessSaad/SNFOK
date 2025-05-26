package templates

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/FearLessSaad/SNFOK/constants/agent_consts"
	"github.com/google/uuid"
)

func GeneratePolicy(policy_file string, namespace string, app_label string) (string, error) {
	policy_templates_dir := os.Getenv("POLICIES_TEMPLATES_DIR")
	policy_template_path := filepath.Join(policy_templates_dir, policy_file)
	applied_policies_dir := os.Getenv("APPLIED_POLICIES_DIR")

	if _, err := os.Stat(policy_template_path); os.IsNotExist(err) {
		return "Policy file is not exists.", err
	}

	err := os.MkdirAll(applied_policies_dir, 0755)
	if err != nil {
		return "Unable to create directory for applied policies. Please check permissions.", err
	}

	content, err := os.ReadFile(policy_template_path)
	if err != nil {
		return "Unable to read policy file. Please check permissions.", err
	}

	newUUID := uuid.NewString()
	id := strings.Split(newUUID, "-")[4]
	policy := strings.ReplaceAll(string(content), agent_consts.POLICY_ID_TEMPLATE, id)
	policy = strings.ReplaceAll(string(policy), agent_consts.POLICY_NAMESPACE_TEMPLATE, namespace)
	policy = strings.ReplaceAll(string(policy), agent_consts.POLICY_APP_LABEL_TEMPLATE, app_label)

	dst_path := filepath.Join(applied_policies_dir, newUUID+".yaml")

	file, err := os.Create(dst_path)
	if err != nil {
		return "Unable to create new policy yaml. Please Check permissions.", err
	}
	file.WriteString(policy)
	defer file.Close()

	return dst_path, nil
}
