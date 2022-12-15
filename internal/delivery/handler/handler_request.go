package handler

type ReqCreateTodos []ReqCreateTodo

type ReqGetID struct {
	ID uint64 `json:"id" validate:"required"`
}

type ReqCreateActivity struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type ReqUpdateActivity struct {
	Email string `json:"email" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type ReqCreateTodo struct {
	ActivityGroupID int    `json:"activity_group_id" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Priority        string `json:"priority" validate:"required"`
}

type ReqUpdateTodo struct {
	ID              uint64 `json:"id" validate:"required"`
	ActivityGroupID int    `json:"activity_group_id" validate:"required"`
	Title           string `json:"title" validate:"required"`
	IsActive        int    `json:"is_active" validate:"required"`
	Priority        string `json:"priority" validate:"required"`
}
