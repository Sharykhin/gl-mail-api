package controller

import (
	"context"

	db "github.com/Sharykhin/gl-mail-api/database"
	"github.com/Sharykhin/gl-mail-api/entity"
)

// Creates creates a new failed mail entity
func Create(ctx context.Context, m entity.MessageRequest) (*entity.Message, error) {
	return db.Create(ctx, m)
}
