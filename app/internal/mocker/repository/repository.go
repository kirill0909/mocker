package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"mocker/config"
	"mocker/internal/mocker"
)

type MockerPGRepo struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewMockerPGRepo(cfg *config.Config, db *sqlx.DB) mocker.PGRepo {
	return &MockerPGRepo{cfg: cfg, db: db}
}

func (r *MockerPGRepo) GetTableNames(ctx context.Context) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, queryGetTableNames)
	if err != nil {
		err = errors.Wrap(err, "MockerPGRepo.GetTableNames.queryGetTableNames")
		return []string{}, err
	}
	defer rows.Close()

	var names []string
	var name string
	for rows.Next() {
		if err = rows.Scan(&name); err != nil {
			err = errors.Wrapf(err, "MockerPGRepo.GetTableNames.Scan(%s)", name)
			return []string{}, err
		}

		names = append(names, name)
	}

	if err = rows.Err(); err != nil {
		err = errors.Wrap(err, "MockerPGRepo.GetTableNames.Err()")
		return []string{}, err
	}

	return names, nil
}
