package repository

import "ddd-to-do-list/internal/aggregate"

type ActivityRepository interface {
	GetActivity() (res aggregate.Activities, err error)
	CreateActivity(activity *aggregate.Activity) error
}
