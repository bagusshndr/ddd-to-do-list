package usecase

import (
	"ddd-to-do-list/internal/aggregate"

	"github.com/stretchr/testify/mock"
)

type ActivityUsecaseMock struct {
	mock.Mock
}

func (m *ActivityUsecaseMock) GetActivity(page int) (res aggregate.Activities, err error) {
	args := m.Called(page)

	return args.Get(0).(aggregate.Activities), args.Error(1)
}

func (m *ActivityUsecaseMock) CreateActivity(email, title string) (uint64, error) {
	args := m.Called(email, title)

	return args.Get(0).(uint64), args.Error(1)
}

func (m *ActivityUsecaseMock) GetActivityByID(id uint64) (res aggregate.Activities, err error) {
	args := m.Called(id)

	return args.Get(0).(aggregate.Activities), args.Error(1)
}

func (m *ActivityUsecaseMock) UpdateActivity(id uint64, email, title string) error {
	args := m.Called(email, title)

	return args.Error(0)
}

func (m *ActivityUsecaseMock) DeleteActivity(id uint64) error {
	args := m.Called(id)

	return args.Error(0)
}
