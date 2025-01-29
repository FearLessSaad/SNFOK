package controllers

import (
	"SNFOK/api/database/queries"
	"SNFOK/api/tooling/beats"
	"encoding/json"
	"net/http"
)

func GetAllHosts(w http.ResponseWriter, r *http.Request) {
	data, err := queries.GetAllIPsAndHostnames()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	resData, _ := json.Marshal(data)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resData)
}

func GetAllLiveHosts(w http.ResponseWriter, r *http.Request) {
	liveHosts := beats.GetAllLiveHosts()
	if liveHosts == nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	response, _ := json.Marshal(liveHosts)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
