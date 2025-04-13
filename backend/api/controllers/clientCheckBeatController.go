package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"SNFOK/api/database/commands"
	"SNFOK/api/database/queries"
)

type ClientInfo struct {
	CPU                     string `json:"cpu"`
	Hostname                string `json:"hostname"`
	KubernetesClientVersion string `json:"kubernetes_client_version"`
	KubernetesServerVersion string `json:"kubernetes_server_version"`
	Memory                  string `json:"memory"`
	OsName                  string `json:"os_name"`
	OsVersion               string `json:"os_version"`
}

func ClientCheckBeatWithInfo(w http.ResponseWriter, r *http.Request) {
	var data ClientInfo
	// Get and Serialze Data From Request
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//Get IP Address Of Client
	ip := strings.Split(strings.TrimSpace(r.RemoteAddr), ":")[0]

	//Check If IP Address is Exists In The Inventory
	if queries.MasterNodeCheckIp(ip) {
		up_error := commands.UpdateMasterNodeByIP(ip, data.OsName, data.OsVersion, data.KubernetesClientVersion, data.KubernetesServerVersion, data.CPU, data.Memory, data.Hostname)
		if up_error != nil {
			fmt.Println("[+] Failed To Update Data Into Database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		return
	}

	//Generate New UUID-4
	newUUID := uuid.New().String()

	//Insert Data Into Database
	insertErr := commands.InsertMasterNode(newUUID, ip, data.OsName, data.OsVersion, data.KubernetesClientVersion, data.KubernetesServerVersion, data.CPU, data.Memory, data.Hostname)

	if insertErr != nil {
		fmt.Println("[+] Failed To Instet Data Into Database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
