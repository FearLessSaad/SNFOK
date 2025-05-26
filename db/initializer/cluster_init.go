package initializer

import (
	"context"

	"github.com/FearLessSaad/SNFOK/db"
	"github.com/FearLessSaad/SNFOK/db/models/k8s"
	"github.com/FearLessSaad/SNFOK/db/utils"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
)

func InitializeCluster() {
	ctx := context.Background()
	conn := db.GetDB()

	logger.Log(logger.INFO, "Initializing 'k8s' schema!")

	utils.SchemaInitializer(ctx, conn, "k8s")
	utils.InitializeTable(ctx, conn, k8s.ClustersTableName, (*k8s.Clusters)(nil))
	utils.InitializeTable(ctx, conn, k8s.AlertsTableName, (*k8s.Alerts)(nil))
	utils.InitializeTable(ctx, conn, k8s.ImplimentedPoliciesTableName, (*k8s.ImplimentedPolicies)(nil))
	utils.InitializeTable(ctx, conn, k8s.AllPoliciesTableName, (*k8s.AllPolicies)(nil))
	logger.Log(logger.INFO, "The 'k8s' schema initialized successfully!")
}
