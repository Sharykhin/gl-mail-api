package controller

import (
	"context"

	"github.com/Sharykhin/gl-mail-api/entity"
	"github.com/Sharykhin/gl-mail-api/grpc"
)

// GetList returns limiter number of rows with count value
func GetList(ctx context.Context, limit, offset int64) ([]entity.FailMail, int, error) {
	fm, err := grpc.Server.GetList(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return fm, 0, nil
}
