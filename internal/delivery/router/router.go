package router

import (
	"ddd-to-do-list/internal/delivery/handler"
	"ddd-to-do-list/internal/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func Router(usecase usecase.ActivityUsecase) *mux.Router {
	h := handler.NewHandler(usecase)

	router := mux.NewRouter()

	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("", h.HandlerGetActivites).Methods(http.MethodGet, http.MethodOptions)

	return router
}
