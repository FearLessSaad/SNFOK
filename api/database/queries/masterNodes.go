package queries

import (
	"SNFOK/api/database"
	"SNFOK/api/database/models"
)

func MasterNodeCheckIp(ipAddress string) bool {
	query := `SELECT COUNT(*) FROM inventory.master_nodes WHERE ip_address = $1`
	var count int

	err := database.DB.QueryRow(query, ipAddress).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func GetAllIPsAndHostnames() ([]models.MasterNodeIpAndHostname, error) {
	query := `SELECT ip_address, hostname FROM inventory.master_nodes`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nodes []models.MasterNodeIpAndHostname

	for rows.Next() {
		var node models.MasterNodeIpAndHostname
		err := rows.Scan(&node.IPAddress, &node.Hostname)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nodes, nil
}
