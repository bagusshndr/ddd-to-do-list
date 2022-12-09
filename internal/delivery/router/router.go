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
		v1.GET("/list-todo", h.HandlerGetActivites)
		v1.GET("", h.HandlerGetActivitesByID)
		v1.POST("/createActivity", h.HandlerCreateActivity)
		v1.PUT("/updateActivity", h.HandlerUpdateActivity)
	}

}
