package router

import (
	"ddd-to-do-list/internal/delivery/handler"
	"ddd-to-do-list/internal/usecase"

	"github.com/labstack/echo/v4"
)

func Router(route *echo.Echo, usecaseActivity usecase.ActivityUsecase, usecaseTodo usecase.TodoUsecase) {
	h := handler.NewHandler(usecaseActivity, usecaseTodo)

	v1 := route.Group("v1")
	{
		v1.GET("/getAllActivity", h.HandlerGetActivites)
		v1.GET("/getAllTodo", h.HandlerGetTodos)
		v1.GET("/getActivityByID", h.HandlerGetActivitesByID)
		v1.GET("/getTodoByID", h.HandlerGetTodosByID)
		v1.POST("/createActivity", h.HandlerCreateActivity)
		v1.POST("/createTodo", h.HandlerCreateTodo)
		v1.PUT("/updateActivity", h.HandlerUpdateActivity)
		v1.PUT("/updateTodo", h.HandlerUpdateTodo)
		v1.DELETE("/deleteActivity", h.HandlerDeleteActivity)
		v1.PUT("/deleteTodo", h.HandlerDeleteTodo)
	}

}
