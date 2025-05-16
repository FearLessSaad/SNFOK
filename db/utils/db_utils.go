package utils

import (
	"github.com/FearLessSaad/SNFOK/tooling/logger"
	"context"
	"fmt"
	"strings"

	"github.com/uptrace/bun"
)

func CheckTableExists(ctx context.Context, db *bun.DB, tableName string) (bool, error) {
	schema := "public"
	name := tableName
	if parts := strings.Split(tableName, "."); len(parts) == 2 {
		schema = parts[0]
		name = parts[1]
	}

	// Query information_schema.tables
	count, err := db.NewSelect().
		Table("information_schema.tables").
		Where("table_schema = ?", schema).
		Where("table_name = ?", name).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func SchemaInitializer(ctx context.Context, conn *bun.DB, schemaName string) error {
	query := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", schemaName)
	_, err := conn.ExecContext(ctx, query)
	if err != nil {
		logger.Log(logger.ERROR, fmt.Sprintf("Failed to create '%s' schema", schemaName), logger.Field{Key: logger.ERROR_MESSAGE, Value: err.Error()})
		panic(err.Error())
	}

	logger.Log(logger.INFO, fmt.Sprintf("Schema '%s' initialized", schemaName))
	return nil
}

func InitializeTable(ctx context.Context, conn *bun.DB, tableName string, model interface{}) {
	exists, err := CheckTableExists(ctx, conn, tableName)
	if err != nil {
		logger.Log(logger.ERROR, "Failed to check if table exists.", logger.Field{Key: logger.ERROR_MESSAGE, Value: err.Error()})
		panic(err)
	}
	if exists {
		logger.Log(logger.INFO, "Table '"+tableName+"' already exists.")
		return
	}

	_, err = conn.NewCreateTable().Model(model).IfNotExists().Exec(ctx)
	if err == nil {
		logger.Log(logger.INFO, "Table '"+tableName+"' created successfully.")
		return
	}

	logger.Log(logger.ERROR, "Failed to create '"+tableName+"' table.", logger.Field{Key: logger.ERROR_MESSAGE, Value: err.Error()})
	panic(err)
}
