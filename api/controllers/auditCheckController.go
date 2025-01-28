package controllers

import (
	"encoding/json"
	"net/http"
)

func AuditCheckSuccess(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"status": "Working Correctly!",
	}
	w.Header().Set("Content-Type", "application/json")
	marshalResponse, _ := json.Marshal(res)
	w.Write(marshalResponse)
}
