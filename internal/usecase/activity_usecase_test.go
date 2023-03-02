package usecase

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/repository"
	"ddd-to-do-list/test"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type activityUsecaseTest struct {
	suite.Suite
	activity *aggregate.Activity
	repo     *repository.ActivityMock
	usecase  ActivityUsecase
}

func (t *activityUsecaseTest) SetupSuite() {
	t.activity = aggregate.RebuildActivity(1, "bagus@bagus.com", "kerja bro")
	t.repo = new(repository.ActivityMock)
	t.usecase = NewActivityUsecase(t.repo)
}

func (t *activityUsecaseTest) TestGetActivity() {
	t.Run("success", func() {
		activities := test.Activities()

		t.repo.On("GetActivity", 10).Return(activities, nil).Once()
		activity, err := t.usecase.GetActivity(10)

		t.NotNil(activity)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.repo.On("GetActivity", 10).Return((aggregate.Activities)(nil), errors.New("error")).Once()
		activity, err := t.usecase.GetActivity(10)

		t.Nil(activity)
		t.Error(err)
	})
}

func (t *activityUsecaseTest) TestGetActivityByID() {
	id := 10
	uid := uint64(id)
	t.Run("success", func() {
		activities := test.Activities()

		t.repo.On("GetActivityByID", uid).Return(activities, nil).Once()
		activity, err := t.usecase.GetActivityByID(uid)

		t.NotNil(activity)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.repo.On("GetActivityByID", uid).Return((aggregate.Activities)(nil), errors.New("error")).Once()
		activity, err := t.usecase.GetActivityByID(uid)

		t.Nil(activity)
		t.Error(err)
	})
}

func (t *activityUsecaseTest) TestCreateActivity() {
	data := test.Activity
	t.Run("success", func() {
		t.repo.On("CreateActivity", data.Email, data.Title).Return(data.ID, nil).Once()
		activity, err := t.usecase.CreateActivity(data.Email, data.Title)

		t.NotNil(activity)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.repo.On("CreateActivity", data.Email, data.Title).Return(uint64(0), errors.New("error")).Once()
		_, err := t.usecase.CreateActivity(data.Email, data.Title)

		t.Error(err)
	})
}

func (t *activityUsecaseTest) TestUpdateActivity() {
	data := test.Activity
	t.Run("success", func() {
		t.repo.On("UpdateActivity", data.Email, data.Title).Return(nil).Once()
		err := t.usecase.UpdateActivity(data.ID, data.Email, data.Title)

		t.NoError(err)
	})

	t.Run("failed", func() {
		t.repo.On("UpdateActivity", data.Email, data.Title).Return(errors.New("error")).Once()
		err := t.usecase.UpdateActivity(data.ID, data.Email, data.Title)

		t.Error(err)
	})
}

func (t *activityUsecaseTest) TestDeleteActivity() {
	data := test.Activity
	t.Run("success", func() {
		t.repo.On("DeleteActivity", data.ID).Return(nil).Once()
		err := t.usecase.DeleteActivity(data.ID)

		t.NoError(err)
	})

	t.Run("failed", func() {
		t.repo.On("DeleteActivity", data.ID).Return(errors.New("error")).Once()
		err := t.usecase.DeleteActivity(data.ID)

		t.Error(err)
	})
}

func TestActivityUsecase(t *testing.T) {
	suite.Run(t, new(activityUsecaseTest))
}
