package model

import (
	"time"
)

type TodoDTO struct {
	ID              uint64
	ActivityGroupID int
	Title           string
	IsActive        int
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
