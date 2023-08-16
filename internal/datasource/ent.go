package datasource

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"subflow-core-go/internal/api/constants"
	"subflow-core-go/internal/config"
	"subflow-core-go/pkg/ent"
)

func NewEntClient(cfg *config.Config) (*ent.Client, error) {
	var dialectName string
	var client *ent.Client
	switch cfg.Datasource.Type {
	case config.DatasourceTypeSQLite:
		dialectName = "sqlite3"
	case config.DatasourceTypeMySQL:
		dialectName = "mysql"
	default:
		zap.S().Fatalw(
			"Unsupported datasource type",
			"type", cfg.Datasource.Type,
		)
	}
	drv, err := sql.Open(dialectName, cfg.Datasource.DSN)
	if err != nil {
		return nil, err
	}

	db := drv.DB()
	db.SetMaxOpenConns(cfg.Datasource.MaxIdleConn)
	db.SetMaxOpenConns(cfg.Datasource.MaxOpenConn)
	db.SetConnMaxLifetime(time.Hour * time.Duration(cfg.Datasource.ConnMaxLifeTime))
	client = ent.NewClient(ent.Driver(drv))

	zap.S().Infow(
		"Datasource initialized successfully",
		"type", cfg.Datasource.Type,
	)

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		zap.S().Fatalw(
			"Failed to create schema resources",
			"err", err,
		)
		return nil, err
	}
	zap.S().Info("Datasource auto migration finished")

	// Predefined data initialize (if not exist)
	if err := createPredefinedRoles(client); err != nil {
		zap.S().Fatalw(
			"Failed to initialize predefined roles",
			"err", err,
		)
		return nil, err
	}
	zap.S().Info("Predefined roles initialized")
	return client, nil
}

func createPredefinedRoles(client *ent.Client) error {
	if client == nil {
		return errors.New("invalid ent client")
	}
	roles := constants.GetAllRoles()
	bulk := make([]*ent.RoleCreate, len(roles))
	for i, role := range roles {
		bulk[i] = client.Role.
			Create().
			SetID(role.ID).
			SetName(role.Name).
			SetDesc(role.Desc)
	}

	err := client.Role.
		CreateBulk(bulk...).
		OnConflict().
		UpdateNewValues().
		Exec(context.Background())
	return err
}
