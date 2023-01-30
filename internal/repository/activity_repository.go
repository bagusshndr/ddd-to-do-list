package repository

import "ddd-to-do-list/internal/aggregate"

type ActivityRepository interface {
	GetActivity(page int) (res aggregate.Activities, err error)
	CreateActivity(email, titile string) (uint64, error)
	GetActivityByID(id uint64) (res aggregate.Activities, err error)
	UpdateActivity(id uint64, email, title string) error
	DeleteActivity(id uint64) error
}
