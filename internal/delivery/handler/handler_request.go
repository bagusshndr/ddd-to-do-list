package handler

type ReqCreateActivity struct {
	Email string `json:"email" validate:"required"`
}
