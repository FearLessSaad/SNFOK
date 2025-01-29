package commands

import (
	"SNFOK/api/database"
	"fmt"
)

func InsertClusterAudit(auditID, auditName, auditDescription, mNodeID, auditCreatedBy string) error {
	query := `
		INSERT INTO autdit.cluster_audits
		(audit_id, audit_name, audit_description, m_node_id, audit_started_at, audit_created_by)
		VALUES ($1, $2, $3, $4, now(), $5)`

	var endedAt interface{}

	_, err := database.DB.Exec(query, auditID, auditName, auditDescription, mNodeID, endedAt, auditCreatedBy)
	if err != nil {
		return fmt.Errorf("failed to insert cluster audit: %w", err)
	}

	return nil
}
