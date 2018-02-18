package contract

import (
	"context"

	"github.com/Sharykhin/gl-mail-api/entity"
)

// InputValidation - an interface for all request structs
type InputValidation interface {
	Validate() error
}

// StorageKeeper provides general interface for managing basis entity
type StorageKeeper interface {
	Create(ctx context.Context, fmr entity.FailMailRequest) (*entity.FailMail, error)
	GetList(ctx context.Context, limit, offset int) ([]entity.FailMail, error)
	Count(ctx context.Context) (int, error)
}
