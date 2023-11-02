package mocker

import (
	"context"
)

type Usecase interface {
	Mock(ctx context.Context) error
}
