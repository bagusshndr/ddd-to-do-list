package handler

import (
	"ddd-to-do-list/internal/shared"
	"ddd-to-do-list/internal/usecase"
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecaseActivity usecase.ActivityUsecase
	usecaseTodo     usecase.TodoUsecase
}

func (h *handler) HandlerGetActivites(c echo.Context) error {
	activities, err := h.usecaseActivity.GetActivity()
	if err != nil {
		log.Println(err)

		return shared.NewResponse("Failed", 400, "Failed", nil, nil).JSON(c)
	}
	return shared.NewResponse("Success", 200, "Success", nil, ActivityResponse{}.Responses(activities)).JSON(c)
}

func (h *handler) HandlerGetTodos(c echo.Context) error {
	todos, err := h.usecaseTodo.GetTodo()
	if err != nil {
		log.Println(err)

		return shared.NewResponse("Failed", 400, "Failed", nil, nil).JSON(c)
	}
	return shared.NewResponse("Success", 200, "Success", nil, TodoResponse{}.Response(todos)).JSON(c)
}

func (h *handler) HandlerGetActivitesByID(c echo.Context) error {
	var body ReqGetID
	c.Bind(&body)
	activities, err := h.usecaseActivity.GetActivityByID(body.ID)
	if err != nil {
		log.Println(err)
		id := fmt.Sprintf("Activity with ID %d Not Found", body.ID)
		return shared.NewResponse("Not Found", 400, id, err.Error(), nil).JSON(c)
	}
	return shared.NewResponse("Success", 200, "success", nil, ActivityResponse{}.Response(activities)).JSON(c)
}

func (h *handler) HandlerGetTodosByID(c echo.Context) error {
	var body ReqGetID
	c.Bind(&body)
	todos, err := h.usecaseTodo.GetTodoByID(body.ID)
	if err != nil {
		log.Println(err)
		return shared.NewResponse("Failed", 400, "Failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse("Success", 200, "success", TodoResponse{}.Response(todos), nil).JSON(c)
}

func (h *handler) HandlerCreateActivity(c echo.Context) error {
	var body ReqCreateActivity
	c.Bind(&body)
	err := h.usecaseActivity.CreateActivity(body.Email, body.Title)
	if err != nil {
		log.Println(err)

		return shared.NewResponse("Failed", 400, "Failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse("Success", 200, "Success", nil, ActivityResponse{}.CreateResponse(body.Email, body.Title)).JSON(c)
}

func (h *handler) HandlerCreateTodo(c echo.Context) error {
	var body ReqCreateTodo
	c.Bind(&body)

	err := h.usecaseTodo.CreateTodo(body.ActivityGroupID, body.Title, body.Priority)
	if err != nil {
		log.Println(err)

		return shared.NewResponse("Failed", 400, "Failed", err.Error(), nil).JSON(c)
	}

	return shared.NewResponse("Success", 200, "Success", nil, nil).JSON(c)
}

func (h *handler) HandlerUpdateActivity(c echo.Context) error {
	var body ReqUpdateActivity
	c.Bind(&body)
	err := h.usecaseActivity.UpdateActivity(body.ID, body.Email, body.Title)
	if err != nil {
		log.Println(err)

		return shared.NewResponse("Failed", 400, "Failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse("Success", 200, "success", nil, nil).JSON(c)
}

func (h *handler) HandlerUpdateTodo(c echo.Context) error {
	var body ReqUpdateTodo
	c.Bind(&body)
	err := h.usecaseTodo.UpdateTodo(body.ID, body.ActivityGroupID, body.IsActive, body.Title, body.Priority)
	if err != nil {
		log.Println(err)

		return shared.NewResponse("Failed", 400, "Failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse("Success", 200, "success", nil, nil).JSON(c)
}

func (h *handler) HandlerDeleteActivity(c echo.Context) error {
	id := c.QueryParam("id")
	uintID, _ := strconv.ParseUint(id, 10, 64)
	err := h.usecaseActivity.DeleteActivity(uintID)
	if err != nil {
		log.Println(err)
		return shared.NewResponse("Failed", 400, "Failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse("Success", 200, "success", nil, nil).JSON(c)
}

func (h *handler) HandlerDeleteTodo(c echo.Context) error {
	id := c.QueryParam("id")
	uintID, _ := strconv.ParseUint(id, 10, 64)
	err := h.usecaseTodo.DeleteTodo(uintID)
	if err != nil {
		log.Println(err)
		return shared.NewResponse("Failed", 400, "Failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse("Success", 200, "success", nil, nil).JSON(c)
}

func NewHandler(usecaseActivity usecase.ActivityUsecase, usecaseTodo usecase.TodoUsecase) *handler {
	return &handler{
		usecaseActivity: usecaseActivity,
		usecaseTodo:     usecaseTodo,
	}
}
