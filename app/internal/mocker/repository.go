package mocker

import (
	"context"
)

type PGRepo interface {
	GetTableNames(context.Context) ([]string, error)
}
