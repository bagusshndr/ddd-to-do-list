package aggregate

import "time"

type Todos []*Todo

type Todo struct {
	ID         uint64
	ActivityID int
	Activity   MapActivities
	Title      string
	IsActive   int
	Priority   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func NewTodo(activitActivityID int, title string, isActive int, priority string) (*Todo, error) {
	return &Todo{
		ActivityID: activitActivityID,
		Title:      title,
		IsActive:   isActive,
		Priority:   priority,
	}, nil
}

func NewTodos(activity MapActivities, title string, isActive int, priority string) (*Todo, error) {
	return &Todo{
		Activity: activity,
		Title:    title,
		IsActive: isActive,
		Priority: priority,
	}, nil
}

func RebuildTodos(id uint64, activity int, title string, isActive int, priority string) (*Todo, error) {
	return &Todo{
		ID:         id,
		ActivityID: activity,
		Title:      title,
		IsActive:   isActive,
		Priority:   priority,
	}, nil
}

func RebuildTodo(id uint64, activity MapActivities, title string, isActive int, priority string) (*Todo, error) {
	return &Todo{
		ID:       id,
		Activity: activity,
		Title:    title,
		IsActive: isActive,
		Priority: priority,
	}, nil
}
