package handler

import (
	"ddd-to-do-list/internal/shared"
	"ddd-to-do-list/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type handler struct {
	usecase usecase.ActivityUsecase
}

func (h *handler) HandlerGetActivites(res http.ResponseWriter, req *http.Request) {
	activities, err := h.usecase.GetActivity()
	if err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write(shared.NewResponse(false, "", err.Error(), nil, nil).JSON())

		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(shared.NewResponse(true, "success", "", ActivityResponse{}.Response(activities), nil).JSON())
}

func (h *handler) HandlerCreateActivity(res http.ResponseWriter, req *http.Request) {
	var body ReqCreateActivity
	json.NewDecoder(req.Body).Decode(&body)

	err := h.usecase.CreateActivity(body.Email)
	if err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write(shared.NewResponse(false, "", err.Error(), nil, nil).JSON())

		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(shared.NewResponse(true, "success", "", nil, nil).JSON())
}

func NewHandler(usecase usecase.ActivityUsecase) *handler {
	return &handler{
		usecase: usecase,
	}
}
