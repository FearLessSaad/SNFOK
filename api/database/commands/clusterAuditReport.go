package commands

import (
	"SNFOK/api/database"
	"fmt"
)

func InsertAuditChecklistReport(auditID string, acID int, status, results string) error {
	query := `
		INSERT INTO autdit.audit_checklist_report
		(audit_id, ac_id, status, results, created_at)
		VALUES ($1, $2, $3, $4, now())`

	_, err := database.DB.Exec(query, auditID, acID, status, results)
	if err != nil {
		return fmt.Errorf("failed to insert audit checklist report: %w", err)
	}

	return nil
}
