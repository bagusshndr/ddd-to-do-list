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
		// ActivityHandler
		v1.GET("/activity-groups", h.HandlerGetActivites)
		v1.GET("/activity-groups/{id}", h.HandlerGetActivitesByID)
		v1.POST("/activity-groups", h.HandlerCreateActivity)
		v1.PUT("/activity-groups", h.HandlerUpdateActivity)
		v1.DELETE("/activity-groups", h.HandlerDeleteActivity)

		// THandler
		v1.GET("/todo-items", h.HandlerGetTodos)
		v1.GET("/todo-items", h.HandlerGetTodosByID)
		v1.POST("/todo-items", h.HandlerCreateTodo)
		v1.PUT("/todo-items", h.HandlerUpdateTodo)
		v1.PUT("/todo-items", h.HandlerDeleteTodo)
	}

}
