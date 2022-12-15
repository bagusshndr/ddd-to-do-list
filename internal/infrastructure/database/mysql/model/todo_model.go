package model

import (
	"time"
)

type TodoDTO struct {
	ID              uint64    `gorm:"id"`
	Title           string    `gorm:"title"`
	ActivityGroupID int       `gorm:"activity_group_id"`
	IsActive        int       `gorm:"is_active"`
	Priority        string    `gorm:"priority"`
	CreatedAt       time.Time `gorm:"created_at"`
	UpdatedAt       time.Time `gorm:"updated_at"`
}

func (t *TodoDTO) TableName() string {
	return "todos"
}
