package repository

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/repository"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type todoRepositoryMysqlTest struct {
	suite.Suite
	mock      sqlmock.Sqlmock
	TodoMYSQL repository.TodoRepository
}

func (t *todoRepositoryMysqlTest) TestGetTodo() {
	todo, _ := aggregate.RebuildTodos(1, 1, "kerja bro", 1, "high")
	query := `SELECT id, activity_group_id, title, is_active, priority FROM todos WHERE is_active = 1`
	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"activity_group_id",
			"title",
			"is_active",
			"priority",
		}).AddRow(
			todo.ID,
			todo.ActivityID,
			todo.Title,
			todo.IsActive,
			todo.Priority,
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		actualTodo, err := t.TodoMYSQL.GetTodo()

		t.NotNil(actualTodo)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New(""))

		actualTodo, err := t.TodoMYSQL.GetTodo()

		log.Println(actualTodo)
		t.Nil(actualTodo)
		t.Error(err)
	})
}

func (t *todoRepositoryMysqlTest) TestGetTodoByID() {
	todo, _ := aggregate.RebuildTodos(1, 1, "kerja bro", 1, "high")
	query := `SELECT id, activity_group_id, title, is_active, priority FROM todos WHERE id = ? AND is_active = 1`
	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"activity_group_id",
			"title",
			"is_active",
			"priority",
		}).AddRow(
			todo.ID,
			todo.ActivityID,
			todo.Title,
			todo.IsActive,
			todo.Priority,
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(todo.ID).WillReturnRows(rows)

		actualTodo, err := t.TodoMYSQL.GetTodoByID(todo.ID)

		t.NotNil(actualTodo)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(todo.ID).WillReturnError(errors.New(""))

		actualTodo, err := t.TodoMYSQL.GetTodoByID(todo.ID)

		log.Println(actualTodo)
		t.Nil(actualTodo)
		t.Error(err)
	})
}

func (t *todoRepositoryMysqlTest) TestCreateTodo() {
	t.Run("success", func() {
		t.mock.ExpectExec("INSERT INTO todos").WillReturnResult(sqlmock.NewResult(1, 1))
		_, err := t.TodoMYSQL.CreateTodo(1, "title", "high")
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.mock.ExpectExec("INSERT INTO todos").WillReturnError(errors.New("failed store todo"))
		_, err := t.TodoMYSQL.CreateTodo(1, "title", "high")
		t.Error(err)
	})
}

func (t *todoRepositoryMysqlTest) TestDeleteTodo() {
	queryDelete := "DELETE FROM todos WHERE id = ?"

	t.Run("success", func() {
		t.mock.ExpectExec(queryDelete).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		t.NoError(t.TodoMYSQL.DeleteTodo(1))
	})

	t.Run("failed", func() {
		t.mock.ExpectExec(queryDelete).WillReturnError(errors.New("failed delete todo"))

		t.Error(t.TodoMYSQL.DeleteTodo(1))
	})
}

func TestTodoRepositoryMySQL(t *testing.T) {
	db, mock, _ := sqlmock.New()

	suite.Run(t, &todoRepositoryMysqlTest{
		mock:      mock,
		TodoMYSQL: NewMysqlTodoRepository(db),
	})
}
