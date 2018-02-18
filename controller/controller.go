package controller

import (
	"context"

	//db "github.com/Sharykhin/gl-mail-api/database"
	"fmt"

	"github.com/Sharykhin/gl-mail-api/contract"
	"github.com/Sharykhin/gl-mail-api/entity"
)

// Create creates a new failed mail entity
func Create(ctx context.Context, fmr entity.FailMailRequest, db contract.StorageKeeper) (*entity.FailMail, error) {
	// there might be some other stuff ...
	return db.Create(ctx, fmr)
}

// GetList returns limiter number of rows with count value
func GetList(ctx context.Context, db contract.StorageKeeper, limit, offset int) ([]entity.FailMail, int, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	chMessages := make(chan []entity.FailMail)
	chCount := make(chan int)
	chErr := make(chan error)
	defer close(chErr)

	var messages []entity.FailMail
	var count int

	go getList(ctx, db, limit, offset, chMessages, chErr)
	go countMessages(ctx, db, chCount, chErr)

	for {
		if chMessages == nil && chCount == nil {
			break
		}

		select {
		case mm, ok := <-chMessages:
			if !ok {
				chMessages = nil
				continue
			}
			messages = mm
		case c, ok := <-chCount:
			if !ok {
				chCount = nil
				continue
			}
			count = c
		case err := <-chErr:
			cancel()
			return nil, 0, err
		}
	}
	return messages, count, nil

}

func getList(ctx context.Context, db contract.StorageKeeper, limit, offset int, chMessages chan<- []entity.FailMail, chErr chan<- error) {
	messages, err := db.GetList(ctx, limit, offset)
	if err != nil {
		if ctx.Err() != context.Canceled {
			chErr <- fmt.Errorf("could not get list of messages: %v", err)
		}
	}
	chMessages <- messages
	close(chMessages)
}

func countMessages(ctx context.Context, db contract.StorageKeeper, chCount chan<- int, chErr chan<- error) {
	c, err := db.Count(ctx)
	if err != nil {
		if ctx.Err() != context.Canceled {
			chErr <- fmt.Errorf("could not count number of rows: %v", err)
		}
	}
	chCount <- c
	close(chCount)
}

// Old implementation
//func Create(ctx context.Context, mr entity.MessageRequest) (*entity.FailMail, error) {
//	return db.Create(ctx, mr)
//}
