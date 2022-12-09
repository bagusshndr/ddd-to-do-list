package repository

import (
	"database/sql"
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/infrastructure/database/mysql/model"
	"ddd-to-do-list/internal/repository"
	"errors"

	"github.com/sirupsen/logrus"
)

type todoRepositoryMysql struct {
	db       *sql.DB
	activity activityRepositoryMySQL
}

func NewMysqlTodoRepository(Conn *sql.DB, activity activityRepositoryMySQL) repository.TodoRepository {
	return &todoRepositoryMysql{
		db:       &sql.DB{},
		activity: activity,
	}
}

func (m *todoRepositoryMysql) fetch(query string, args ...interface{}) (aggregate.Todos, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()
	todoDTOs := []model.TodoDTO{}
	for rows.Next() {
		t := model.TodoDTO{}
		err = rows.Scan(
			&t.ID,
			&t.ActivityGroupID,
			&t.Title,
			&t.IsActive,
			&t.Priority,
		)

		if err != nil {
			return nil, err
		}

		todoDTOs = append(todoDTOs, t)
	}

	todos := aggregate.Todos{}
	for _, todoDTO := range todoDTOs {
		aggregateTodo, _ := aggregate.RebuildTodos(
			todoDTO.ID,
			todoDTO.ActivityGroupID,
			todoDTO.Title,
			todoDTO.IsActive,
			todoDTO.Priority,
		)
		todos = append(todos, aggregateTodo)
	}

	return todos, nil
}

func (m *todoRepositoryMysql) GetTodo() (res aggregate.Todos, err error) {
	query := `SELECT id, email, title FROM activities`

	res, err = m.fetch(query)
	if err != nil {
		return nil, errors.New("")
	}

	return
}

func (m *todoRepositoryMysql) GetTodoByID(id uint64) (res aggregate.Todos, err error) {
	query := `SELECT id, email, title FROM activities WHERE id = ? LIMIT 1`

	res, err = m.fetch(query, id)
	if err != nil {
		return nil, errors.New("")
	}

	return
}

func (m *todoRepositoryMysql) CreateTodo(todo *aggregate.Todo) error {
	query := "INSERT INTO activities (id, email, title) VALUES(?, ?, ?)"
	_, err := m.db.Exec(
		query,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *todoRepositoryMysql) DeleteTodo(id uint64) error {
	query := "DELETE activities WHERE id = ?"
	_, err := s.db.Exec(
		query,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
