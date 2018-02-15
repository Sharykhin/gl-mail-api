package controller

import (
	"context"

	//db "github.com/Sharykhin/gl-mail-api/database"
	"github.com/Sharykhin/gl-mail-api/entity"
)

//TODO: experimental way to make code testable
type StorageKeeper interface {
	Create(ctx context.Context, m entity.MessageRequest) (*entity.Message, error)
}

// Create creates a new failed mail entity
func Create(ctx context.Context, mr entity.MessageRequest, db StorageKeeper) (*entity.Message, error) {
	// there might be some other stuff ...
	return db.Create(ctx, mr)
}
