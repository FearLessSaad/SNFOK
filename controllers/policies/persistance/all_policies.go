package persistance

import (
	"context"

	"github.com/FearLessSaad/SNFOK/db"
	"github.com/FearLessSaad/SNFOK/db/models/k8s"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
)

func GetAllPolices() ([]k8s.AllPolicies, error) {

	conn := db.GetDB()
	ctx := context.Background()

	all_policies := new([]k8s.AllPolicies)
	err := conn.NewSelect().Model(all_policies).Scan(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to execute select query on 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return []k8s.AllPolicies{}, err
	}

	return *all_policies, nil
}

func GetPlicysById(id string) (k8s.AllPolicies, error) {

	conn := db.GetDB()
	ctx := context.Background()

	alert := new(k8s.AllPolicies)
	err := conn.NewSelect().Model(alert).Where("id = ?", id).Limit(1).Scan(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to execute select query on 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return k8s.AllPolicies{}, err
	}

	return *alert, nil
}
