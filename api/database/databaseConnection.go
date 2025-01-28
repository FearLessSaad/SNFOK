package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "saad1234"
	dbname   = "SNFOK"
)

var DB *sql.DB

func InitDB() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}
	fmt.Println("[+] Connected to the database successfully!")

	// Check, If required tables are not exists then create them
	checkTables()
}

func checkTables() {
	// Table Names
	table_names := []string{"audit_checklist_report", "cluster_audits", "audit_checklist"}

	for i := 0; i < len(table_names); i++ {
		exists := tableExists(table_names[i])
		if !exists {
			fmt.Printf("[!] Table %s is not exists!\n", table_names[i])
		}
	}
}

func tableExists(tableName string) bool {
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM information_schema.tables
            WHERE table_name = $1
        );
    `
	var exists bool
	err := DB.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
