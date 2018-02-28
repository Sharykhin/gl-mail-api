package controller

import (
	"context"

	"fmt"

	"github.com/Sharykhin/gl-mail-api/contract"
	"github.com/Sharykhin/gl-mail-api/entity"
	"github.com/Sharykhin/gl-mail-api/grpc"
)

// FailMail is a reference to a private struct that implements all necessary methods
var FailMail failMail

type failMail struct {
	storage contract.StorageProvider
}

func (c failMail) GetList(ctx context.Context, limit, offset int64) ([]entity.FailMail, int64, error) {
	fm, err := getList(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	count, err := getCount(ctx)
	if err != nil {
		return nil, 0, err
	}
	return fm, count, nil
}

func init() {
	FailMail = failMail{storage: grpc.Server}
}

func getList(ctx context.Context, limit, offset int64) ([]entity.FailMail, error) {
	fms, err := grpc.Server.GetList(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("could not get list of failed mails: %v", err)
	}

	return fms, nil
}

func getCount(ctx context.Context) (int64, error) {
	count, err := grpc.Server.Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
