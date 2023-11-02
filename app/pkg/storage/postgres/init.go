package postgres

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"mocker/config"
)

func InitPGDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	_, span := otel.Tracer("").Start(ctx, "storage.InitPsqlDB")
	defer span.End()

	connectionURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	database, err := sqlx.Open("pgx", connectionURL)
	if err != nil {
		err = errors.Wrap(err, "storage.postgres.InitPGDB.Open()")
		return nil, err
	}

	if err = database.Ping(); err != nil {
		err = errors.Wrap(err, "storage.postgres.InitPGDB.Ping()")
		return nil, err
	}

	return database, nil
}
