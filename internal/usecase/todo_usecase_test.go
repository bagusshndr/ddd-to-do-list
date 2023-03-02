package usecase

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/repository"
	"ddd-to-do-list/test"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type todoUsecaseTest struct {
	suite.Suite
	todo    *aggregate.Todo
	repo    *repository.TodoMock
	usecase TodoUsecase
}

func (t *todoUsecaseTest) SetupSuite() {
	t.todo, _ = aggregate.RebuildTodos(1, 1, "kerja bro", 1, "high")
	t.repo = new(repository.TodoMock)
	t.usecase = NewTodoUsecase(t.repo)
}

func (t *todoUsecaseTest) TestGetTodo() {
	todos := test.Todos()
	t.Run("success", func() {

		t.repo.On("GetTodo").Return(todos, nil).Once()
		todo, err := t.usecase.GetTodo()

		t.NotNil(todo)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.repo.On("GetTodo").Return((aggregate.Todos)(nil), errors.New("error")).Once()
		todo, err := t.usecase.GetTodo()

		t.Nil(todo)
		t.Error(err)
	})
}

func (t *todoUsecaseTest) TestGetTodoByID() {
	todos := test.Todos()
	t.Run("success", func() {

		t.repo.On("GetTodoByID", uint64(1)).Return(todos, nil).Once()
		todo, err := t.usecase.GetTodoByID(uint64(1))

		t.NotNil(todo)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.repo.On("GetTodoByID", uint64(1)).Return((aggregate.Todos)(nil), errors.New("error")).Once()
		todo, err := t.usecase.GetTodoByID(uint64(1))

		t.Nil(todo)
		t.Error(err)
	})
}

func (t *todoUsecaseTest) TestCreateTodo() {
	todo := test.Todo
	t.Run("success", func() {

		t.repo.On("CreateTodo", todo.ActivityID, todo.Title, todo.Priority).Return(uint64(1), nil).Once()
		_, err := t.usecase.CreateTodo(todo.ActivityID, todo.Title, todo.Priority)

		t.NoError(err)
	})

	t.Run("failed", func() {
		t.repo.On("CreateTodo", todo.ActivityID, todo.Title, todo.Priority).Return(uint64(0), errors.New("error")).Once()
		_, err := t.usecase.CreateTodo(todo.ActivityID, todo.Title, todo.Priority)

		t.Error(err)
	})
}

func (t *todoUsecaseTest) TestUpdateTodo() {
	todo := test.Todo
	t.Run("success", func() {

		t.repo.On("UpdateTodo", todo.ID, todo.ActivityID, todo.IsActive, todo.Title, todo.Priority).Return(nil).Once()
		err := t.usecase.UpdateTodo(todo.ID, todo.ActivityID, todo.IsActive, todo.Title, todo.Priority)

		t.NoError(err)
	})

	t.Run("failed", func() {
		t.repo.On("UpdateTodo", todo.ID, todo.ActivityID, todo.IsActive, todo.Title, todo.Priority).Return(errors.New("error")).Once()
		err := t.usecase.UpdateTodo(todo.ID, todo.ActivityID, todo.IsActive, todo.Title, todo.Priority)

		t.Error(err)
	})
}

func (t *todoUsecaseTest) TestDeleteTodo() {
	t.Run("success", func() {

		t.repo.On("DeleteTodo", uint64(1)).Return(nil).Once()
		err := t.usecase.DeleteTodo(uint64(1))

		t.NoError(err)
	})

	t.Run("failed", func() {
		t.repo.On("DeleteTodo", uint64(1)).Return(errors.New("error")).Once()
		err := t.usecase.DeleteTodo(uint64(1))

		t.Error(err)
	})
}

func TestTodoUsecase(t *testing.T) {
	suite.Run(t, new(todoUsecaseTest))
}
