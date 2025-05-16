package initializer

import (
	"context"

	"github.com/FearLessSaad/SNFOK/db"
	"github.com/FearLessSaad/SNFOK/db/models/auth"
	"github.com/FearLessSaad/SNFOK/db/utils"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
)

func InitializeAuth() {
	ctx := context.Background()
	conn := db.GetDB()

	logger.Log(logger.INFO, "Initializing 'auth' schema!")

	utils.SchemaInitializer(ctx, conn, "auth")
	utils.InitializeTable(ctx, conn, auth.UsersTableName, (*auth.Users)(nil))

	logger.Log(logger.INFO, "The 'auth' schema initialized successfully!")
}
