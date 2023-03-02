package usecase

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/repository"
	"errors"
)

type activityUsecase struct {
	repo repository.ActivityRepository
}

func (u *activityUsecase) GetActivity(page int) (aggregate.Activities, error) {
	activity, err := u.repo.GetActivity(page)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (u *activityUsecase) GetActivityByID(id uint64) (aggregate.Activities, error) {
	activity, _ := u.repo.GetActivityByID(id)
	if len(activity) <= 0 {
		return nil, errors.New("data not found")
	}
	return activity, nil
}

func (u *activityUsecase) CreateActivity(email, titile string) (uint64, error) {
	uid, err := u.repo.CreateActivity(email, titile)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func (u *activityUsecase) UpdateActivity(id uint64, email, title string) error {
	err := u.repo.UpdateActivity(id, email, title)
	if err != nil {
		return err
	}

	return nil
}

func (u *activityUsecase) DeleteActivity(id uint64) error {
	err := u.repo.DeleteActivity(id)
	if err != nil {
		return err
	}

	return nil
}

func NewActivityUsecase(repo repository.ActivityRepository) ActivityUsecase {
	return &activityUsecase{
		repo: repo,
	}
}
