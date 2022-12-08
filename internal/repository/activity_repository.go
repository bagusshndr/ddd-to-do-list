package repository

import "ddd-to-do-list/internal/aggregate"

type ActivityRepository interface {
	GetActivity() (res aggregate.Activities, err error)
	CreateActivity(activity *aggregate.Activity) error
	GetActivityByID(id uint64) (res aggregate.Activities, err error)
}
