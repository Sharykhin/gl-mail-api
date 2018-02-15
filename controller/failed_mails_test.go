package controller

import (
	"context"
	"testing"

	"github.com/Sharykhin/gl-mail-api/entity"
)

type mockStorage struct {
}

func (ms mockStorage) Create(ctx context.Context, m entity.MessageRequest) (*entity.Message, error) {
	return &entity.Message{
		ID:     12,
		Action: "register",
		Payload: map[string]interface{}{
			"to": "test@mail.com",
		},
		Reason: "test reason",
	}, nil
}

func TestCreate(t *testing.T) {
	t.Run("create success", func(t *testing.T) {
		mr := entity.MessageRequest{
			Action: "register",
			Payload: map[string]interface{}{
				"to": "test@mail.com",
			},
			Reason: "test reason",
		}

		ctx := context.Background()

		m, err := Create(ctx, mr, mockStorage{})
		if err != nil {
			t.Errorf("expected successfull creation, but got error: %v", err)
		}

		if m.Action != "register" {
			t.Errorf("expected action 'register' but got '%s'", m.Action)
		}

		if m.Reason != "test reason" {
			t.Errorf("expected reason 'test reason' but got '%s'", m.Reason)
		}
	})
}
