package usecase

import "ddd-to-do-list/internal/aggregate"

type ActivityUsecase interface {
	GetActivity(page int) (aggregate.Activities, error)
	GetActivityByID(id uint64) (aggregate.Activities, error)
	CreateActivity(email, titile string) (uint64, error)
	UpdateActivity(id uint64, email, title string) error
	DeleteActivity(id uint64) error
}
