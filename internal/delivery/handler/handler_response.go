package handler

import (
	"ddd-to-do-list/internal/aggregate"
	"time"
)

type ActivityResponses []ActivityResponse

type TodoResponses []TodoResponse

type ActivityResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uint64    `json:"id"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
}

type TodoResponse struct {
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ID              uint64    `json:"id"`
	ActivityGroupID int       `json:"activity_group_id" validate:"required"`
	Title           string    `json:"title" validate:"required"`
	IsActive        int       `json:"is_active" validate:"required"`
	Priority        string    `json:"priority" validate:"required"`
}

func (response ActivityResponse) Responses(activity aggregate.Activities) (result ActivityResponses) {
	for _, src := range activity {
		response.ID = src.ID
		response.Email = src.Email
		response.Title = src.Title
		response.CreatedAt = src.CreatedAt
		response.UpdatedAt = src.UpdatedAt

		result = append(result, ActivityResponse{
			response.CreatedAt,
			response.UpdatedAt,
			response.ID,
			response.Email,
			response.Title,
		},
		)
	}
	return result
}

func (response ActivityResponse) Response(activity aggregate.Activities) ActivityResponse {
	for _, src := range activity {
		response.ID = src.ID
		response.Email = src.Email
		response.Title = src.Title
		response.CreatedAt = src.CreatedAt
		response.UpdatedAt = src.UpdatedAt
	}
	return ActivityResponse{
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
		ID:        response.ID,
		Email:     response.Email,
		Title:     response.Title,
	}
}

func (response ActivityResponse) CreateResponse(id uint64, email, title string) ActivityResponse {
	return ActivityResponse{
		ID:        id,
		Email:     email,
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (response TodoResponse) Responses(todo aggregate.Todos) (result TodoResponses) {
	for _, src := range todo {
		response.CreatedAt = src.CreatedAt
		response.UpdatedAt = src.UpdatedAt
		response.ID = src.ID
		response.ActivityGroupID = src.ActivityID
		response.Title = src.Title
		response.IsActive = src.IsActive
		response.Priority = src.Priority

		result = append(result, TodoResponse{
			response.CreatedAt,
			response.UpdatedAt,
			response.ID,
			response.ActivityGroupID,
			response.Title,
			response.IsActive,
			response.Priority,
		},
		)
	}
	return result
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
		CreatedAt:       response.CreatedAt,
		UpdatedAt:       response.UpdatedAt,
		ID:              response.ID,
		ActivityGroupID: response.ActivityGroupID,
		Title:           response.Title,
		IsActive:        response.IsActive,
		Priority:        response.Priority,
	}
}
