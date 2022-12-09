package handler

import (
	"ddd-to-do-list/internal/aggregate"
)

type ActivityResponse struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Title string `json:"title"`
}

type TodoResponse struct {
	ID              uint64 `json:"id"`
	ActivityGroupID int    `json:"activity_group_id" validate:"required"`
	Title           string `json:"title" validate:"required"`
	IsActive        int    `json:"is_active" validate:"required"`
	Priority        string `json:"priority" validate:"required"`
}

func (response ActivityResponse) Response(activity aggregate.Activities) ActivityResponse {
	for _, src := range activity {
		response.ID = src.ID
		response.Email = src.Email
		response.Title = src.Title
	}

	return ActivityResponse{
		response.ID,
		response.Email,
		response.Title,
	}
}
func (response TodoResponse) Response(todo aggregate.Todos) TodoResponse {
	for _, src := range todo {
		response.ID = src.ID
		response.ActivityGroupID = src.ActivityID
		response.Title = src.Title
		response.IsActive = src.IsActive
		response.Priority = src.Priority
	}
	return TodoResponse{
		ID:              response.ID,
		ActivityGroupID: response.ActivityGroupID,
		Title:           response.Title,
		IsActive:        response.IsActive,
		Priority:        response.Priority,
	}
}
