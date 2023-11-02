package usecase

import (
	"context"
	"fmt"
	"log"
	"mocker/config"
	"mocker/internal/mocker"
	"mocker/pkg/utils"
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

	tables, err := u.pgRepo.GetTableNames(ctx)
	if err != nil {
		return err
	}

	for _, table := range tables {
		columns, err := u.pgRepo.GetColumns(ctx, table.Name)
		if err != nil {
			return err
		}

		// uuid, text, integer, bigint time, timestamp with time zone, timestamp without time zone, numeric, boolean, jsonb, ARRAY
		// for _, column := range columnNames {
		// 	log.Println(column.Type)
		// }

		var query string
		for _, column := range columns {
			switch column.Type {
			case "uuid":
				// log.Println(table, column)
				query = fmt.Sprintf("UPDATE %s.%s SET %s = %s", table.SchemaName, table.Name, column.Name, utils.GetRandomUUID())
				log.Println(query)
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
