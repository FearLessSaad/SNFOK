package persistance

import (
	"context"

	"github.com/FearLessSaad/SNFOK/db"
	"github.com/FearLessSaad/SNFOK/db/models/k8s"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
)

func GetAllAlerts() ([]k8s.Alerts, error) {

	conn := db.GetDB()
	ctx := context.Background()

	alerts := new([]k8s.Alerts)
	err := conn.NewSelect().Model(alerts).Scan(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to execute select query on 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return []k8s.Alerts{}, err
	}

	return *alerts, nil
}

func CountAlerts() (int, error) {

	conn := db.GetDB()
	ctx := context.Background()

	alerts := new([]k8s.Alerts)
	count, err := conn.NewSelect().Model(alerts).Count(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to execute select query on 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return 0, err
	}

	return count, nil
}

func GetAlertsById(id string) (k8s.Alerts, error) {

	conn := db.GetDB()
	ctx := context.Background()

	alerts := new(k8s.Alerts)
	err := conn.NewSelect().Model(alerts).Where("id = ?", id).Limit(1).Scan(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to execute select query on 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return k8s.Alerts{}, err
	}

	return *alerts, nil
}

func CreateAlert(data k8s.Alerts) error {
	conn := db.GetDB()
	ctx := context.Background()

	_, err := conn.NewInsert().Model(&data).Exec(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to execute insert query on 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return err
	}

	return nil
}
