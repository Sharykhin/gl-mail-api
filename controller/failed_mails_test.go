package controller

import (
	"context"
	"testing"

	"encoding/json"

	"github.com/Sharykhin/gl-mail-api/entity"
	"github.com/stretchr/testify/mock"
)

type mockStorage struct {
	mock.Mock
}

func (m *mockStorage) Create(ctx context.Context, mr entity.FailMailRequest) (*entity.FailMail, error) {
	ret := m.Called(ctx, mr)

	if rf, ok := ret.Get(0).(func(ctx context.Context, m entity.FailMailRequest) (*entity.FailMail, error)); ok {
		m, err := rf(ctx, mr)
		return m, err
	}

	err := ret.Error(1)
	return nil, err
}

func (m *mockStorage) GetList(ctx context.Context, limit, offset int) ([]entity.FailMail, error) {
	return nil, nil
}

func (m *mockStorage) Count(ctx context.Context) (int, error) {
	return 0, nil
}

func TestCreate(t *testing.T) {
	ctx := context.Background()

	mr := entity.FailMailRequest{
		Action: "register",
		Payload: entity.Payload{
			"to": json.RawMessage("test@mail.com"),
		},
		Reason: "test reason",
	}

	m := &entity.FailMail{
		ID:     12,
		Action: "register",
		Payload: entity.Payload{
			"to": json.RawMessage("test@mail.com"),
		},
		Reason: "test reason",
	}

	mockS := new(mockStorage)
	mockS.On("Create", ctx, mr).Return(m, nil).Once()
	_, err := Create(ctx, mr, mockS)
	if err != nil {
		t.Errorf("expected success creation, but got an error: %v", err)
	}
	mockS.AssertExpectations(t)

	//t.Run("create success", func(t *testing.T) {
	//	mr := entity.MessageRequest{
	//		Action: "register",
	//		Payload: map[string]interface{}{
	//			"to": "test@mail.com",
	//		},
	//		Reason: "test reason",
	//	}
	//
	//	ctx := context.Background()
	//
	//	m, err := Create(ctx, mr, mockStorage{})
	//	if err != nil {
	//		t.Errorf("expected successfull creation, but got error: %v", err)
	//	}
	//
	//	if m.Action != "register" {
	//		t.Errorf("expected action 'register' but got '%s'", m.Action)
	//	}
	//
	//	if m.Reason != "test reason" {
	//		t.Errorf("expected reason 'test reason' but got '%s'", m.Reason)
	//	}
	//})
}
