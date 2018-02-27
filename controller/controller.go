package controller

import (
	"context"

	"github.com/Sharykhin/gl-mail-api/contract"
	"github.com/Sharykhin/gl-mail-api/entity"
	"github.com/Sharykhin/gl-mail-api/grpc"
)

// FailMail is a reference to a private struct that implements all necessary methods
var FailMail failMail

type failMail struct {
	storage contract.StorageProvider
}

// GetList returns limiter number of rows with a total count
func (c failMail) GetList(ctx context.Context, limit, offset int64) ([]entity.FailMail, int, error) {
	fm, err := grpc.Server.GetList(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return fm, 0, nil
}

func init() {
	FailMail = failMail{storage: grpc.Server}
}
