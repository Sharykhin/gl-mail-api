package controller

import (
	"context"
	"testing"

	"encoding/json"

	"github.com/Sharykhin/gl-mail-api/entity"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

type mockStorage struct {
	mock.Mock
}

func (m mockStorage) GetList(ctx context.Context, limit, offset int) ([]entity.FailMail, error) {
	return nil, nil
}

func (m mockStorage) Count(ctx context.Context) (int, error) {
	return 0, nil
}

type successMockStorage struct {
	mockStorage
}

func (m successMockStorage) Create(ctx context.Context, mr entity.FailMailRequest) (*entity.FailMail, error) {
	ret := m.Called(ctx, mr)
	fm, err := ret.Get(0).(*entity.FailMail), ret.Get(1)
	if err != nil {
		return nil, err.(error)
	}

	return fm, nil
}

type errorMockStorage struct {
	mockStorage
}

func (m *errorMockStorage) Create(ctx context.Context, mr entity.FailMailRequest) (*entity.FailMail, error) {
	m.Called(ctx, mr)
	var err error = errors.New("something went wrong")
	return nil, err
}

func TestCreate(t *testing.T) {
	t.Run("success creation", func(t *testing.T) {
		ctx := context.Background()
		var payload = []byte(`{"to": "test@mail.com"}`)

		fmr := entity.FailMailRequest{
			Action:  "register",
			Payload: json.RawMessage(payload),
			Reason:  "test reason",
		}

		m := &entity.FailMail{
			ID:      12,
			Action:  "register",
			Payload: entity.Payload(payload),
			Reason:  "test reason",
		}

		mockS := new(successMockStorage)
		mockS.On("Create", ctx, fmr).Return(m, nil).Once()
		_, err := Create(ctx, fmr, mockS)
		if err != nil {
			t.Errorf("expected success creation, but got an error: %v", err)
		}
		mockS.AssertExpectations(t)
	})

	t.Run("bad creation", func(t *testing.T) {
		ctx := context.Background()
		var payload = []byte(`{"to": "test@mail.com"}`)

		fmr := entity.FailMailRequest{
			Action:  "register",
			Payload: json.RawMessage(payload),
			Reason:  "test reason",
		}

		var expectedErr error = errors.New("something went wrong")
		mockS := new(errorMockStorage)
		mockS.On("Create", ctx, fmr).Return(nil, expectedErr).Once()
		fm, err := Create(ctx, fmr, mockS)
		if err == nil {
			t.Errorf("expected error %s, but got a result: %v", err, fm)
		}
		mockS.AssertExpectations(t)
		if expectedErr.Error() != err.Error() {
			t.Errorf("expected error %s, but got another error %s", expectedErr, err)
		}
	})
}
