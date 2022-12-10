package model

import (
	"ddd-to-do-list/internal/aggregate"
	"time"
)

type ActivityDTO struct {
	ID        uint64    `gorm:"id"`
	Email     string    `gorm:"email"`
	Title     string    `gorm:"title"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

func (a ActivityDTO) Aggregate() *aggregate.Activity {
	activity, _ := aggregate.NewActivity(a.Email, a.Title)
	return activity
}

func (t *ActivityDTO) TableName() string {
	return "activities"
}
