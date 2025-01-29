package commands

import (
	"SNFOK/api/database"
	"fmt"
)

func InsertMasterNode(mNodeID, ipAddress, os, osVersion, k8sClientVersion, k8sServerVersion, cpu, ram, hostname string) error {
	query := `
		INSERT INTO inventory.master_nodes
		(m_node_id, ip_address, operating_system, os_version, kubernetes_client_version, kubernetes_server_version, cpu, ram, hostname) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := database.DB.Exec(query, mNodeID, ipAddress, os, osVersion, k8sClientVersion, k8sServerVersion, cpu, ram, hostname)
	if err != nil {
		return err
	}
	fmt.Println("Inserted Successfully!")
	return nil
}

func UpdateMasterNodeByIP(ipAddress, os, osVersion, k8sClientVersion, k8sServerVersion, cpu, ram, hostname string) error {
	query := `
		UPDATE inventory.master_nodes
		SET operating_system = $1,
		os_version = $2,
		kubernetes_client_version = $3,
		kubernetes_server_version = $4,
		cpu = $5,
		ram = $6,
		hostname = $7,
		updated_at = now()
		WHERE ip_address = $8`

	result, err := database.DB.Exec(query, os, osVersion, k8sClientVersion, k8sServerVersion, cpu, ram, hostname, ipAddress)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no record found with IP address: %s", ipAddress)
	}

	return nil
}
