package repository

import "ddd-to-do-list/internal/aggregate"

type ActivityRepository interface {
	GetActivity() (res aggregate.Activities, err error)
	CreateActivity(email string) error
	GetActivityByID(id uint64) (res aggregate.Activities, err error)
}
