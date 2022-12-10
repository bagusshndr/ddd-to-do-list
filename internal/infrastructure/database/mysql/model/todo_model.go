package model

import (
	"time"
)

type TodoDTO struct {
	ID              uint64    `gorm:"id"`
	ActivityGroupID int       `gorm:"activity_group_id"`
	Title           string    `gorm:"title"`
	IsActive        int       `gorm:"is_active"`
	Priority        string    `gorm:"priority"`
	CreatedAt       time.Time `gorm:"created_at"`
	UpdatedAt       time.Time `gorm:"updated_at"`
}

func (t *TodoDTO) TableName() string {
	return "todos"
}
