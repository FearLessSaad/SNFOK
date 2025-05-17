package k8s

import (
	"time"

	"github.com/uptrace/bun"
)

type AuditFields struct {
	CreatedBy string       `bun:"type:varchar(150)"`
	UpdatedBy string       `bun:"type:varchar(150)"`
	CreatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt bun.NullTime `bun:",nullzero"`
}

type Clusters struct {
	bun.BaseModel `bun:"table:k8s.clusters,alias:h"`

	ID          string `bun:",pk,type:uuid,default:gen_random_uuid()"`
	ClusterName string
	MasterIP    string
	AgentPort   int
	Description string

	AuditFields
}

const ClustersTableName = "k8s.clusters"
