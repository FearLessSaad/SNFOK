package controllers

import (
	"SNFOK/client/tooling"
	"encoding/json"
	"net/http"
)

func HealthCheckController(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "online",
	}

	jsonRes, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func HealthCheckAndInfoCOntroller(w http.ResponseWriter, r *http.Request) {
	var responseData map[string]string = make(map[string]string)

	// Get Operating System
	responseData["os_name"] = tooling.RunCommand("lsb_release -d | awk '{print $2}'")

	// Get Operating System Version
	responseData["os_version"] = tooling.RunCommand("grep '^VERSION=' /etc/os-release | cut -d= -f2 | tr -d '\"'")

	//Get Kubernetes Client Version
	responseData["kubernetes_client_version"] = tooling.RunCommand("kubectl version | grep Client | awk '{print $3}'")

	//Get Kubernetes Server Version
	responseData["kubernetes_server_version"] = tooling.RunCommand("kubectl version | grep Server | awk '{print $3}'")

	//Get CPU Model
	responseData["cpu"] = tooling.RunCommand("lscpu | grep \"Model name:\" | cut -d ':' -f2")

	//Get Memory
	responseData["memory"] = tooling.RunCommand("lsmem | grep \"Total online memory\" | cut -d ':' -f2 | tr -d ' ' | cut -d 'G' -f1") + "GB"

	//Get Hostname
	responseData["hostname"] = tooling.RunCommand("hostname")

	data, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(data)
}
