package usecase

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/repository"
)

type activityUsecase struct {
	repo repository.ActivityRepository
}

func (u *activityUsecase) GetActivity() (aggregate.Activities, error) {
	activity, err := u.repo.GetActivity()
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func NewActivityUsecase(repo repository.ActivityRepository) ActivityUsecase {
	return &activityUsecase{
		repo: repo,
	}
}
