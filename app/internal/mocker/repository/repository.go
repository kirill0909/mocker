package repository

import (
	"github.com/jmoiron/sqlx"
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
