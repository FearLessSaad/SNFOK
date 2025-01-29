package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"SNFOK/api/database"
	"SNFOK/api/types"
)

func AuditCheckSuccess(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"status": "Working Correctly!",
	}
	w.Header().Set("Content-Type", "application/json")
	marshalResponse, _ := json.Marshal(res)
	w.Write(marshalResponse)
}

func GetAllAudits(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM autdit.cluster_audits;")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var audits []types.ClusterAudit
	for rows.Next() {
		var audit types.ClusterAudit
		err := rows.Scan(
			&audit.AuditID,
			&audit.AuditName,
			&audit.AuditDescription,
			&audit.MachineIPAddress,
			&audit.OperatingSystem,
			&audit.OSVersion,
			&audit.KubernetesVersion,
			&audit.AuditStartedAt,
			&audit.AuditEndedAt,
			&audit.AuditCreatedBy,
		)
		if err != nil {
			log.Fatal(err)
		}
		audits = append(audits, audit)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(audits)
	w.Write(res)
}
