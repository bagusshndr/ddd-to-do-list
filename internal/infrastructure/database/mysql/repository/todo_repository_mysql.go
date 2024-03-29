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
	db *sql.DB
}

func NewMysqlTodoRepository(Conn *sql.DB) repository.TodoRepository {
	return &todoRepositoryMysql{
		db: Conn,
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
	query := `SELECT id, activity_group_id, title, is_active, priority FROM todos WHERE is_active = 1`

	res, err = m.fetch(query)
	if err != nil {
		return nil, errors.New("")
	}

	return
}

func (m *todoRepositoryMysql) GetTodoByID(id uint64) (res aggregate.Todos, err error) {
	query := `SELECT id, activity_group_id, title, is_active, priority FROM todos WHERE id = ? AND is_active = 1`

	res, err = m.fetch(query, id)
	if err != nil {
		return nil, errors.New("")
	}

	return
}

func (m *todoRepositoryMysql) CreateTodo(activitGroupID int, title, priority string) (uint64, error) {
	query := "INSERT INTO todos (activity_group_id, title, is_active, priority) VALUES(?, ?, 1, ?)"
	res, err := m.db.Exec(
		query,
		activitGroupID,
		title,
		priority,
	)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()

	uId := uint64(id)
	return uId, nil
}

func (m *todoRepositoryMysql) UpdateTodo(id uint64, activitGroupID, IsActive int, title, priority string) error {
	query := "UPDATE todos SET activity_group_id = ?, title = ?, is_active = ?, priority = ? WHERE id = ?"
	_, err := m.db.Exec(
		query,
		activitGroupID,
		title,
		IsActive,
		priority,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *todoRepositoryMysql) DeleteTodo(id uint64) error {
	query := "DELETE FROM todos WHERE id = ?"
	res, err := s.db.Exec(
		query,
		id,
	)
	if res == nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
