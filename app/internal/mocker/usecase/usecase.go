package usecase

import (
	"context"
	"fmt"
	"log"
	"mocker/config"
	"mocker/internal/mocker"
	"mocker/internal/models"
	"mocker/pkg/utils"
	"strings"
)

type MockerUC struct {
	cfg    *config.Config
	pgRepo mocker.PGRepo
}

func NewMockerUC(cfg *config.Config, pgRepo mocker.PGRepo) mocker.Usecase {
	return &MockerUC{cfg: cfg, pgRepo: pgRepo}
}

func (u *MockerUC) Mock(ctx context.Context) error {
	tables, err := u.pgRepo.GetTableNames(ctx)
	if err != nil {
		return err
	}

	for _, table := range tables {
		columns, err := u.pgRepo.GetColumns(ctx, table.Name)
		if err != nil {
			return err
		}

		// var query string
		for _, column := range columns {
			switch column.Type {
			case "uuid":
				u.handleUUIDCase(ctx, table, column)
				log.Println("Update ", table, column)
			case "text":
				// log.Println(column)
			case "integer":
				// log.Println(column)
			case "bigint":
				// log.Println(column)
			case "timestamp with time zone":
				// log.Println(column)
			case "timestamp without time zone":
				// log.Println(column)
			case "numeric":
				// log.Println(column)
			case "boolean":
				// log.Println(column)
			case "jsonb":
				// log.Println(column)
			case "ARRAY":
				// log.Println(column)
			}
		}
	}

	return nil
}

func (u *MockerUC) handleUUIDCase(ctx context.Context, table models.TableData, column models.ColumnData) {
	rowsNum, err := u.pgRepo.GetRowsNum(ctx, fmt.Sprintf("%s.%s", table.SchemaName, table.Name))
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < rowsNum; i++ {
		query := fmt.Sprintf("UPDATE %s.%s SET %s = '%s'::UUID", table.SchemaName, table.Name, column.Name, utils.GetRandomUUID())
		err = u.pgRepo.Mock(ctx, query)
		if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			continue
		}
	}
}
