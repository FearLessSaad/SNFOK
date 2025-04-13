package beats

import (
	"SNFOK/api/database/models"
	"SNFOK/api/database/queries"
	"net/http"
	"time"
)

func GetAllLiveHosts() []models.MasterNodeIpAndHostname {
	host, _ := queries.GetAllIPsAndHostnames()
	var liveHosts []models.MasterNodeIpAndHostname
	for _, host := range host {
		if CheckTargetOnline("http://" + host.IPAddress + ":45667/health/check") {
			liveHosts = append(liveHosts, host)
		}
	}
	return liveHosts
}

func CheckTargetOnline(url string) bool {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode >= 200 && resp.StatusCode < 300
}
