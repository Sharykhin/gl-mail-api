package contract

import (
	"context"

	"github.com/Sharykhin/gl-mail-api/entity"
)

// StorageProvider is an interface for getting data
type StorageProvider interface {
	GetList(ctx context.Context, limit, offset int64) ([]entity.FailMail, error)
}
