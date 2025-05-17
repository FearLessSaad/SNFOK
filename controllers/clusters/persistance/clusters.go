package persistance

import (
	"context"

	"github.com/FearLessSaad/SNFOK/db"
	"github.com/FearLessSaad/SNFOK/db/models/k8s"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
)

func GetAllClusters() ([]k8s.Clusters, error) {

	conn := db.GetDB()
	ctx := context.Background()

	clusters := new([]k8s.Clusters)
	err := conn.NewSelect().Model(clusters).Scan(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to execute select query on 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return []k8s.Clusters{}, err
	}

	return *clusters, nil
}

func CreateCluster(data k8s.Clusters) error {
	conn := db.GetDB()
	ctx := context.Background()

	_, err := conn.NewInsert().Model(&data).Exec(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to execute insert query on 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return err
	}

	return nil
}

func CheckMasterIPExists(masterIP string) (bool, error) {
	conn := db.GetDB()
	ctx := context.Background()

	exists, err := conn.NewSelect().
		Model((*k8s.Clusters)(nil)).
		Where("master_ip = ?", masterIP).
		Exists(ctx)

	if err != nil {
		logger.Log(logger.ERROR, "Failed to check if master IP exists in 'k8s.clusters'.", logger.Field{Key: "error", Value: err.Error()})
		return false, err
	}

	return exists, nil
}
