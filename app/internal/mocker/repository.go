package mocker

import (
	"context"
	"mocker/internal/models"
)

type PGRepo interface {
	GetTableNames(context.Context) ([]models.TableData, error)
	GetColumns(context.Context, string) ([]models.ColumnData, error)
	GetRowsNum(context.Context, string) (int, error)
	Mock(context.Context, string) error
}
