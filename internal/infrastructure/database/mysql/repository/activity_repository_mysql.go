package repository

import (
	"database/sql"
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/infrastructure/database/mysql/model"
	"ddd-to-do-list/internal/repository"
	"errors"

	"github.com/sirupsen/logrus"
)

type activityRepositoryMySQL struct {
	db *sql.DB
}

func NewMysqlActivityRepository(Conn *sql.DB) repository.ActivityRepository {
	return &activityRepositoryMySQL{Conn}
}

func (m *activityRepositoryMySQL) fetch(query string, args ...interface{}) (aggregate.Activities, error) {
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
	activityDTOs := []model.ActivityDTO{}
	for rows.Next() {
		t := model.ActivityDTO{}
		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.Title,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		activityDTOs = append(activityDTOs, t)
	}

	activities := aggregate.Activities{}
	for _, activityDTO := range activityDTOs {
		activities = append(activities, aggregate.RebuildActivity(
			activityDTO.ID,
			activityDTO.Email,
			activityDTO.Title,
		))
	}

	return activities, nil
}

func (m *activityRepositoryMySQL) GetActivity() (res aggregate.Activities, err error) {
	query := `SELECT id, email, title FROM activities`

	res, err = m.fetch(query)
	if err != nil {
		return nil, errors.New("")
	}

	return
}

func (m *activityRepositoryMySQL) CreateActivity(activity *aggregate.Activity) error {
	query := "INSERT INTO activities (id, email, title) VALUES(?, ?, ?)"
	_, err := m.db.Exec(
		query,
		activity.ID,
		activity.Email,
		activity.Title,
	)
	if err != nil {
		return err
	}
	return nil
}
