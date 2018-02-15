package controller

import (
	"context"

	//db "github.com/Sharykhin/gl-mail-api/database"
	"github.com/Sharykhin/gl-mail-api/entity"
)

type StorageKeeper interface {
	Create(ctx context.Context, m entity.MessageRequest) (*entity.Message, error)
}

// Create creates a new failed mail entity
func Create(ctx context.Context, m entity.MessageRequest, db StorageKeeper) (*entity.Message, error) {
	return db.Create(ctx, m)
}
