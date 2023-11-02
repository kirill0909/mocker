package usecase

import (
	"context"
	"log"
	"mocker/config"
	"mocker/internal/mocker"
)

type MockerUC struct {
	cfg    *config.Config
	pgRepo mocker.PGRepo
}

func NewMockerUC(cfg *config.Config, pgRepo mocker.PGRepo) mocker.Usecase {
	return &MockerUC{cfg: cfg, pgRepo: pgRepo}
}

func (m *MockerUC) Mock(ctx context.Context) error {
	log.Println("Hello from mocker usecase")
	return nil
}
