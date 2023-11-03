package usecase

import (
	"context"
	"fmt"
	"log"
	"mocker/config"
	"mocker/internal/mocker"
	"mocker/internal/models"
	"strings"
	"sync"
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
	log.Println(tables)

	wg := sync.WaitGroup{}
	for _, table := range tables {
		columns, err := u.pgRepo.GetColumns(ctx, table.Name)
		if err != nil {
			return err
		}

		for _, column := range columns {
			switch column.Type {
			case "uuid":
				wg.Add(1)
				go func(table models.TableData, column models.ColumnData) {
					defer wg.Done()
					u.handleUUIDCase(ctx, table, column)
					log.Printf("Updated: Table: %s.%s Column: %s", table.SchemaName, table.Name, column.Name)
				}(table, column)
			case "text":
				// wg.Add(1)
				// go func(table models.TableData, column models.ColumnData) {
				// 	defer wg.Done()
				// 	u.handleTextCase(ctx, table, column)
				// 	log.Printf("Updated: Table: %s.%s Column: %s", table.SchemaName, table.Name, column.Name)
				// }(table, column)
			case "integer", "bigint", "smallint", "numeric":
				// log.Println(column)
			case "timestamp with time zone", "timestamp without time zone":
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

	wg.Wait()

	return nil
}

func (u *MockerUC) handleTextCase(ctx context.Context, table models.TableData, column models.ColumnData) {
	query := fmt.Sprintf("UPDATE %s.%s SET %s = 'My Best Mock'", table.SchemaName, table.Name, column.Name)
	if err := u.pgRepo.Mock(ctx, query); err != nil {
		log.Println(err)
	}
}

func (u *MockerUC) handleUUIDCase(ctx context.Context, table models.TableData, column models.ColumnData) {
	rowsNum, err := u.pgRepo.GetRowsNum(ctx, fmt.Sprintf("%s.%s", table.SchemaName, table.Name))
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < rowsNum; i++ {
		query := fmt.Sprintf("UPDATE %s.%s SET %s = gen_random_uuid()", table.SchemaName, table.Name, column.Name)
		err = u.pgRepo.Mock(ctx, query)
		if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			continue
		}
	}
}
