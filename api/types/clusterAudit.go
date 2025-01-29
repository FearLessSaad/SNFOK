package types

import "time"

type ClusterAudit struct {
	AuditID           string    `json:"audit_id"`
	AuditName         string    `json:"audit_name"`
	AuditDescription  string    `json:"audit_description"`
	MachineIPAddress  string    `json:"machine_ip_address"`
	OperatingSystem   string    `json:"operating_system"`
	OSVersion         string    `json:"os_version"`
	KubernetesVersion string    `json:"kubernetes_version"`
	AuditStartedAt    time.Time `json:"audit_started_at"`
	AuditEndedAt      time.Time `json:"audit_ended_at"`
	AuditCreatedBy    string    `json:"audit_created_by"`
}
