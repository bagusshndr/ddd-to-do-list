package aggregate

import (
	"errors"
	"time"
)

type Activities []*Activity
type MapActivities map[uint64]Activity

type Activity struct {
	ID        uint64
	Email     string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewActivity(email, title string) (*Activity, error) {
	if email == "" {
		return &Activity{}, errors.New("email cannot be null")
	}
	if title == "" {
		return &Activity{}, errors.New("title cannot be null")
	}
	return &Activity{
		Email: email,
		Title: title,
	}, nil
}

func RebuildActivity(id uint64, email, title string) *Activity {
	return &Activity{
		ID:    id,
		Email: email,
		Title: title,
	}
}

func RebuildActivities(id uint64, email, title string) *Activities {
	return &Activities{
		&Activity{
			ID:    id,
			Email: email,
			Title: title,
		},
	}
}
