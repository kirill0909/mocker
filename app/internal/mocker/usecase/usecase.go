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

func (u *MockerUC) Mock(ctx context.Context) error {
	log.Println("Hello from mocker usecase")

	tableNames, err := u.pgRepo.GetTableNames(ctx)
	if err != nil {
		return err
	}

	log.Println(tableNames)
	log.Println(len(tableNames))

	return nil
}
