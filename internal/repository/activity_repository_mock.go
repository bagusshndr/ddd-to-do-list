package repository

import (
	"ddd-to-do-list/internal/aggregate"

	"github.com/stretchr/testify/mock"
)

type ActivityMock struct {
	mock.Mock
}

func (m *ActivityMock) GetActivity() (res aggregate.Activities, err error) {
	args := m.Called()

	return args.Get(0).(aggregate.Activities), args.Error(1)
}

func (m *ActivityMock) CreateActivity(activity *aggregate.Activity) error {
	args := m.Called(activity)

	return args.Error(0)
}
