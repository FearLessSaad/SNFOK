package persistance

import (
	"context"

	"github.com/FearLessSaad/SNFOK/db"
	"github.com/FearLessSaad/SNFOK/db/models/k8s"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
)

func GetAllImplimentedPolicies() ([]k8s.ImplimentedPolicies, error) {

	conn := db.GetDB()
	ctx := context.Background()

	i_policies := new([]k8s.ImplimentedPolicies)
	err := conn.NewSelect().Model(i_policies).Scan(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to execute select query on 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return []k8s.ImplimentedPolicies{}, err
	}

	return *i_policies, nil
}

func CreateImplimentedPolicy(data k8s.ImplimentedPolicies) error {
	conn := db.GetDB()
	ctx := context.Background()

	_, err := conn.NewInsert().Model(&data).Exec(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to execute insert query on 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return err
	}

	return nil
}
