package handler

import (
	"ddd-to-do-list/internal/shared"
	"ddd-to-do-list/internal/usecase"
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

		return shared.NewResponse(false, 400, "failed", nil, nil).JSON(c)
	}
	return shared.NewResponse(true, 200, "success", nil, ActivityResponse{}.Response(activities)).JSON(c)
}

func (h *handler) HandlerGetTodos(c echo.Context) error {
	todos, err := h.usecaseTodo.GetTodo()
	if err != nil {
		log.Println(err)

		return shared.NewResponse(false, 400, "failed", nil, nil).JSON(c)
	}
	return shared.NewResponse(true, 200, "success", nil, TodoResponse{}.Response(todos)).JSON(c)
}

func (h *handler) HandlerGetActivitesByID(c echo.Context) error {
	var body ReqGetID
	c.Bind(&body)
	activities, err := h.usecaseActivity.GetActivityByID(body.ID)
	if err != nil {
		log.Println(err)
		return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse(true, 200, "success", nil, ActivityResponse{}.Response(activities)).JSON(c)
}

func (h *handler) HandlerGetTodosByID(c echo.Context) error {
	var body ReqGetID
	c.Bind(&body)
	todos, err := h.usecaseTodo.GetTodoByID(body.ID)
	if err != nil {
		log.Println(err)
		return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse(true, 200, "success", TodoResponse{}.Response(todos), nil).JSON(c)
}

func (h *handler) HandlerCreateActivity(c echo.Context) error {
	var body ReqCreateActivity
	c.Bind(&body)
	err := h.usecaseActivity.CreateActivity(body.Email, body.Title)
	if err != nil {
		log.Println(err)

		return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse(true, 200, "success", nil, nil).JSON(c)
}

func (h *handler) HandlerCreateTodo(c echo.Context) error {
	var body ReqCreateTodos
	c.Bind(&body)
	for _, reqCreateTodo := range body {
		err := h.usecaseTodo.CreateTodo(reqCreateTodo.ActivityGroupID, reqCreateTodo.Title, reqCreateTodo.Priority)
		if err != nil {
			log.Println(err)

			return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)
		}
	}
	return shared.NewResponse(true, 200, "success", nil, nil).JSON(c)
}

func (h *handler) HandlerUpdateActivity(c echo.Context) error {
	var body ReqUpdateActivity
	c.Bind(&body)
	err := h.usecaseActivity.UpdateActivity(body.ID, body.Email, body.Title)
	if err != nil {
		log.Println(err)

		return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse(true, 200, "success", nil, nil).JSON(c)
}

func (h *handler) HandlerUpdateTodo(c echo.Context) error {
	var body ReqUpdateTodo
	c.Bind(&body)
	err := h.usecaseTodo.UpdateTodo(body.ID, body.ActivityGroupID, body.IsActive, body.Title, body.Priority)
	if err != nil {
		log.Println(err)

		return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse(true, 200, "success", nil, nil).JSON(c)
}

func (h *handler) HandlerDeleteActivity(c echo.Context) error {
	id := c.QueryParam("id")
	uintID, _ := strconv.ParseUint(id, 10, 64)
	err := h.usecaseActivity.DeleteActivity(uintID)
	if err != nil {
		log.Println(err)
		return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse(true, 200, "success", nil, nil).JSON(c)
}

func (h *handler) HandlerDeleteTodo(c echo.Context) error {
	id := c.QueryParam("id")
	uintID, _ := strconv.ParseUint(id, 10, 64)
	err := h.usecaseTodo.DeleteTodo(uintID)
	if err != nil {
		log.Println(err)
		return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)
	}
	return shared.NewResponse(true, 200, "success", nil, nil).JSON(c)
}

func NewHandler(usecaseActivity usecase.ActivityUsecase, usecaseTodo usecase.TodoUsecase) *handler {
	return &handler{
		usecaseActivity: usecaseActivity,
		usecaseTodo:     usecaseTodo,
	}
}
