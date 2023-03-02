package handler

import (
	"ddd-to-do-list/internal/usecase"
	"testing"

	"github.com/stretchr/testify/suite"
)

type handlerTest struct {
	suite.Suite
	activityUsecase *usecase.ActivityUsecaseMock
	todoUsecase     *usecase.TodoUsecaseMock
	handler         *handler
}

func (t *handlerTest) SetupSuite() {
	t.activityUsecase = new(usecase.ActivityUsecaseMock)
	t.handler = NewHandler(
		t.activityUsecase,
		t.todoUsecase,
	)
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(handlerTest))
}
