package controller

import (
	"context"
	"testing"

	"encoding/json"

	"github.com/Sharykhin/gl-mail-api/entity"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
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

func (m mockStorage) Create(ctx context.Context, mr entity.FailMailRequest) (*entity.FailMail, error) {
	return nil, nil
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

type successMockList struct {
	mockStorage
}

func (m successMockList) GetList(ctx context.Context, limit, offset int) ([]entity.FailMail, error) {
	ret := m.Called(ctx, limit, offset)
	return ret.Get(0).([]entity.FailMail), nil
}

func (m successMockList) Count(ctx context.Context) (int, error) {
	return 10, nil
}

type errorMockList struct {
	mockStorage
}

func (m errorMockList) GetList(ctx context.Context, limit, offset int) ([]entity.FailMail, error) {
	m.Called(ctx, limit, offset)
	return nil, errors.New("something went wrong")
}

func (m errorMockList) Count(ctx context.Context) (int, error) {
	return 0, nil
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

func TestGetList(t *testing.T) {
	t.Run("success getting", func(t *testing.T) {
		ctx := context.Background()
		cc, _ := context.WithCancel(ctx)

		var fm = []entity.FailMail{
			{
				ID:      19,
				Action:  "register",
				Payload: []byte(`{}`),
				Reason:  "test reason",
			},
			{
				ID:      20,
				Action:  "register",
				Payload: []byte(`{}`),
				Reason:  "test reason",
			},
		}

		mockS := new(successMockList)
		mockS.On("GetList", cc, 10, 0).Return(fm, nil).Once()
		mockS.On("Count", cc).Return(10, nil).Once()

		fm, c, err := GetList(ctx, mockS, 10, 0)
		if err != nil {
			t.Errorf("unexpected an errror: %v", err)
		}

		assert.Equal(t, 10, c, "count should be equal 10")
		assert.NotNil(t, fm, "an array of failed mails can not be nil")
		assert.Equal(t, 2, len(fm), "number of mails should be equal 2")
	})

	t.Run("bad getting", func(t *testing.T) {
		ctx := context.Background()
		cc, _ := context.WithCancel(ctx)

		var fm = []entity.FailMail{
			{
				ID:      19,
				Action:  "register",
				Payload: []byte(`{}`),
				Reason:  "test reason",
			},
			{
				ID:      20,
				Action:  "register",
				Payload: []byte(`{}`),
				Reason:  "test reason",
			},
		}

		mockS := new(errorMockList)
		mockS.On("GetList", cc, 10, 0).Return(fm, nil).Once()
		mockS.On("Count", cc).Return(10, nil).Once()

		fm, c, err := GetList(ctx, mockS, 10, 0)
		if err == nil {
			t.Errorf("expected error but got retusl: %v", fm)
		}

		assert.Nil(t, fm, "result  should be nil")
		assert.Equal(t, 0, c, "count should return zero")
	})
}
