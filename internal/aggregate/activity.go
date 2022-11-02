package aggregate

import "errors"

type Activities []*Activity

type Activity struct {
	ID    uint64
	Email string
	Title string
}

func NewActivity(email, title string) (*Activity, error) {
	if email == "" {
		return &Activity{}, errors.New("email cannot be empty")
	}
	if title == "" {
		return &Activity{}, errors.New("title cannot be empty")
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
