package k8s

import (
	"github.com/uptrace/bun"
)

type Alerts struct {
	bun.BaseModel `bun:"table:k8s.alerts,alias:h"`

	ID          string `bun:",pk,type:uuid,default:gen_random_uuid()"`
	ElasticId   string
	AlertTitle  string
	Description string
	Pod         string
	Namespace   string
	Severity    string

	AuditFields
}

const AlertsTableName = "k8s.alerts"
