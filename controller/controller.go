package controller

import (
	"context"

	"github.com/Sharykhin/gl-mail-api/contract"
	"github.com/Sharykhin/gl-mail-api/entity"
	"github.com/Sharykhin/gl-mail-api/grpc"
)

var (
	// FailMail exports method for getting list of failed mails
	FailMail = failMail{storage: grpc.Server}
)

type failMail struct {
	storage contract.StorageProvider
}

func (c failMail) GetList(ctx context.Context, limit, offset int64) ([]entity.FailMail, int64, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	chFms := make(chan []entity.FailMail)
	chFmc := make(chan int64)
	chErr := make(chan error)
	defer close(chErr)

	var fms []entity.FailMail
	var count int64

	go getList(ctx, chFms, chErr, limit, offset)

	go getCount(ctx, chFmc, chErr)

	for {
		if chFmc == nil && chFms == nil {
			return fms, count, nil
		}
		select {
		case receivedFms, ok := <-chFms:
			if !ok {
				chFms = nil
				continue
			}
			fms = receivedFms
		case receivedFmc, ok := <-chFmc:
			if !ok {
				chFmc = nil
				continue
			}
			count = receivedFmc
		case err := <-chErr:
			cancel()
			return nil, 0, err
		}
	}
}

//func init() {
//	FailMail = failMail{storage: grpc.Server}
//}

func getList(ctx context.Context, chFms chan<- []entity.FailMail, chErr chan<- error, limit, offset int64) {
	defer close(chFms)
	fms, err := grpc.Server.GetList(ctx, limit, offset)
	if err != nil {
		if chErr == nil {
			return
		}
		chErr <- err

	} else {
		chFms <- fms
	}

}

func getCount(ctx context.Context, chFmc chan<- int64, chErr chan<- error) {
	defer close(chFmc)
	count, err := grpc.Server.Count(ctx)
	if err != nil {
		if chErr == nil {
			return
		}
		chErr <- err
	} else {
		chFmc <- count
	}
}
