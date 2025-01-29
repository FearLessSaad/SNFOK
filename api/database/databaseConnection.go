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

var tables = map[string]string{
	"cluster_audits": `CREATE TABLE IF NOT EXISTS autdit.cluster_audits
	(
		audit_id varchar(36) NOT NULL,
		audit_name varchar(120),
		audit_description text COLLATE pg_catalog."default",
		machine_ip_address varchar(40),
		operating_system varchar(120),
		os_version varchar(40),
		kubernetes_version varchar(40),
		audit_started_at timestamp without time zone DEFAULT now(),
		audit_ended_at timestamp without time zone,
		audit_created_by varchar(26),
		CONSTRAINT cluster_audits_pkey PRIMARY KEY (audit_id)
	);`,
	"audit_checklist": `CREATE TABLE IF NOT EXISTS autdit.audit_checklist
	(
		ac_id SERIAL NOT NULL,
		check_name varchar(120),
		check_description text COLLATE pg_catalog."default",
		check_command text COLLATE pg_catalog."default",
		profile_ability text COLLATE pg_catalog."default",
		rational text COLLATE pg_catalog."default",
		impact text COLLATE pg_catalog."default",
		verification text COLLATE pg_catalog."default",
		IG1 BOOLEAN,
		IG2 BOOLEAN,
		IG3 BOOLEAN,
		created_at timestamp without time zone DEFAULT now(),
		updated_at timestamp without time zone,
		CONSTRAINT audit_checklist_pkey PRIMARY KEY (ac_id)
	)`,
	"audit_checklist_report": `CREATE TABLE IF NOT EXISTS autdit.audit_checklist_report
	(
		acr_id SERIAL NOT NULL,
		audit_id varchar(36),
		ac_id int,
		status varchar(50),
		results text COLLATE pg_catalog."default",
		created_at timestamp without time zone DEFAULT now(),
		CONSTRAINT audit_checklist_report_pkey PRIMARY KEY (acr_id)
	)`,
}

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

	// Check, If required tables and schemas are not exists then create them
	preFlightChecks()
}

func preFlightChecks() {
	// Table Names
	table_names := []string{"audit_checklist_report", "cluster_audits", "audit_checklist"}
	schema_names := []string{"autdit"}

	//check schemas
	for i := 0; i < len(schema_names); i++ {
		exists := schemaExists(schema_names[i])
		if !exists {
			fmt.Printf("[!] Schema %s is not exists!\n", schema_names[i])
			// Create Schema
			createSchema(schema_names[i])
			fmt.Printf("[+] Schema '%s' is created successfully!\n", schema_names[i])
		} else {
			fmt.Printf("[+] Schema '%s' is exists!\n", schema_names[i])
		}
	}

	// Checking tables
	for i := 0; i < len(table_names); i++ {
		exists := tableExists(table_names[i])
		if !exists {
			fmt.Printf("[!] Table '%s' is not exists!\n", table_names[i])
			// Create Table
			createTable(tables[table_names[i]], table_names[i])
			fmt.Printf("[+] Table '%s' is created successfully!\n", table_names[i])
		} else {
			fmt.Printf("[+] Table '%s' is exists!\n", table_names[i])
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

func schemaExists(schemaName string) bool {
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM information_schema.schemata
            WHERE schema_name = $1
        );
    `

	var exists bool
	err := DB.QueryRow(query, schemaName).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}

func createSchema(schemaName string) {
	query := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", schemaName)

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("[!] Error while creating schema '%s'!", schemaName)
	}
}

func createTable(query string, table_name string) {
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("[!] Failed to create table '%s'", table_name)
	}
}
