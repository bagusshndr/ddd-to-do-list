package usecase

import "ddd-to-do-list/internal/aggregate"

type ActivityUsecase interface {
	GetActivity() (aggregate.Activities, error)
	CreateActivity(email string) error
}
