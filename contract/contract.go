package contract

import (
	"context"

	"github.com/Sharykhin/gl-mail-api/entity"
)

type (
	// StorageProvider is an interface for getting data
	StorageProvider interface {
		GetList(ctx context.Context, limit, offset int64) ([]entity.FailMail, error)
	}
)
