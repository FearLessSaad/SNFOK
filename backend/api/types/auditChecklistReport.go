package types

import "time"

type AuditChecklistReport struct {
	ACRID     int       `json:"acr_id"`
	AuditID   string    `json:"audit_id"`
	ACID      int       `json:"ac_id"`
	Status    string    `json:"status"`
	Results   string    `json:"results"`
	CreatedAt time.Time `json:"created_at"`
}
