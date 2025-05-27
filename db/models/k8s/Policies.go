package k8s

import (
	"github.com/uptrace/bun"
)

type ImplimentedPolicies struct {
	bun.BaseModel `bun:"table:k8s.implimented_policies,alias:h"`

	ID             string `bun:",pk,type:uuid,default:gen_random_uuid()"`
	PolicyTitle    string
	Description    string
	AppLabel       string
	Namespace      string
	PolicyFilePath string

	AuditFields
}

const ImplimentedPoliciesTableName = "k8s.implimented_policies"

type AllPolicies struct {
	bun.BaseModel `bun:"table:k8s.all_policies,alias:h"`

	ID             string `bun:",pk,type:uuid,default:gen_random_uuid()"`
	PolicyTitle    string
	Description    string
	PolicyType     string
	PolicyFilePath string

	AuditFields
}

const AllPoliciesTableName = "k8s.all_policies"
