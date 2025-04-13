package tooling

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

func SendInitialBeatWithInfo() {
	var data ClientInfo

	// Get Operating System
	data.OsName = RunCommand("lsb_release -d | awk '{print $2}'")

	// Get Operating System Version
	data.OsVersion = RunCommand("grep '^VERSION=' /etc/os-release | cut -d= -f2 | tr -d '\"'")

	//Get Kubernetes Client Version
	data.KubernetesClientVersion = RunCommand("kubectl version | grep Client | awk '{print $3}'")

	//Get Kubernetes Server Version
	data.KubernetesServerVersion = RunCommand("kubectl version | grep Server | awk '{print $3}'")

	//Get CPU Model
	data.CPU = RunCommand("lscpu | grep \"Model name:\" | cut -d ':' -f2")

	//Get Memory
	data.Memory = RunCommand("lsmem | grep \"Total online memory\" | cut -d ':' -f2 | tr -d ' ' | cut -d 'G' -f1") + "GB"

	//Get Hostname
	data.Hostname = RunCommand("hostname")

	req, _ := json.Marshal(data)

	URL := "http://192.168.2.100:6444/client/beat/info"

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("[+] Successfully Connected To The SNFOK Server. Server Response:", resp.Status)

}
