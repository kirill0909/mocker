package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"mocker/config"
	"mocker/internal/mocker"
	"mocker/internal/models"
)

type MockerPGRepo struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewMockerPGRepo(cfg *config.Config, db *sqlx.DB) mocker.PGRepo {
	return &MockerPGRepo{cfg: cfg, db: db}
}

func (r *MockerPGRepo) GetTableNames(ctx context.Context) ([]models.TableData, error) {
	rows, err := r.db.QueryContext(ctx, queryGetTableNames)
	if err != nil {
		err = errors.Wrap(err, "MockerPGRepo.GetTableNames.queryGetTableNames")
		return []models.TableData{}, err
	}
	defer rows.Close()

	var tables []models.TableData
	var table models.TableData
	for rows.Next() {
		if err = rows.Scan(&table.SchemaName, &table.Name); err != nil {
			err = errors.Wrapf(err, "MockerPGRepo.GetTableNames.Scan(%s)", table)
			return []models.TableData{}, err
		}

		tables = append(tables, table)
	}

	if err = rows.Err(); err != nil {
		err = errors.Wrap(err, "MockerPGRepo.GetTableNames.Err()")
		return []models.TableData{}, err
	}

	return tables, nil
}

func (r *MockerPGRepo) GetColumns(ctx context.Context, tableName string) ([]models.ColumnData, error) {
	rows, err := r.db.QueryContext(ctx, queryGetColumns, tableName)
	if err != nil {
		err = errors.Wrap(err, "MockerPGRepo.GetColumns.queryGetColumns")
		return []models.ColumnData{}, err
	}
	defer rows.Close()

	var columns []models.ColumnData
	var column models.ColumnData
	for rows.Next() {
		if err = rows.Scan(&column.Name, &column.Type); err != nil {
			err = errors.Wrapf(err, "MockerPGRepo.GetColumns.Scan(%v)", column)
			return []models.ColumnData{}, err
		}

		columns = append(columns, column)
	}

	if err = rows.Err(); err != nil {
		err = errors.Wrap(err, "MockerPGRepo.GetColumns.Err")
		return []models.ColumnData{}, err
	}

	return columns, nil
}

func (r *MockerPGRepo) Mock(ctx context.Context, query string) error {
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "MockerPGRepo.Mock")
	}

	return nil
}

func (r *MockerPGRepo) GetRowsNum(ctx context.Context, tableName string) (int, error) {
	var result int
	if err := r.db.GetContext(ctx, &result, fmt.Sprintf(queryGetRowsNum, tableName)); err != nil {
		return 0, errors.Wrapf(err, "MockerPGRepo.GetRowsNum.queryGetRowsNum. TableName: %s", tableName)
	}

	return result, nil
}
