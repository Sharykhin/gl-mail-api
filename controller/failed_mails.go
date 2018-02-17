package controller

import (
	"context"

	//db "github.com/Sharykhin/gl-mail-api/database"
	"fmt"

	"github.com/Sharykhin/gl-mail-api/entity"
)

// TODO: experimental way to make code testable

// Something Interface is something for something to implement
type StorageKeeper interface {
	Create(ctx context.Context, fmr entity.FailMailRequest) (*entity.Message, error)
	GetList(ctx context.Context, limit, offset int) ([]entity.Message, error)
	Count(ctx context.Context) (int, error)
}

// Create creates a new failed mail entity
func Create(ctx context.Context, fmr entity.FailMailRequest, db StorageKeeper) (*entity.Message, error) {
	// there might be some other stuff ...
	return db.Create(ctx, fmr)
}

// GetList returns limiter number of rows with count value
func GetList(ctx context.Context, db StorageKeeper, limit, offset int) ([]entity.Message, int, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	chMessages := make(chan []entity.Message)
	//defer close(chMessages)
	chCount := make(chan int)
	//defer close(chCount)
	chErr := make(chan error)
	defer close(chErr)

	var messages []entity.Message
	var count int

	go getList(ctx, db, limit, offset, chMessages, chErr)
	go countMessages(ctx, db, chCount, chErr)

	for {
		if chMessages == nil && chCount == nil {
			break
		}

		select {
		case mm, ok := <-chMessages:
			fmt.Println("get messages", mm, ok)
			if !ok {
				chMessages = nil
				continue
			}
			messages = mm
		case c, ok := <-chCount:
			fmt.Println("get count", c, ok)
			if !ok {
				chCount = nil
				continue
			}
			count = c
		case err, ok := <-chErr:
			fmt.Println("HA HAH A", ok)
			cancel()
			return nil, 0, err
		}
	}
	return messages, count, nil

}

func getList(ctx context.Context, db StorageKeeper, limit, offset int, chMessages chan<- []entity.Message, chErr chan<- error) {
	messages, err := db.GetList(ctx, limit, offset)
	if err != nil {
		if ctx.Err() != context.Canceled {
			chErr <- fmt.Errorf("could not get list of messages: %v", err)
		}
	}
	chMessages <- messages
	close(chMessages)
}

func countMessages(ctx context.Context, db StorageKeeper, chCount chan<- int, chErr chan<- error) {
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
//func Create(ctx context.Context, mr entity.MessageRequest) (*entity.Message, error) {
//	return db.Create(ctx, mr)
//}
