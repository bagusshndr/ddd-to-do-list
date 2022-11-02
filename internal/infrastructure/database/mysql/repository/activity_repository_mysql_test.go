package repository

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/repository"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type activityRepositoryMysqlTest struct {
	suite.Suite
	mock          sqlmock.Sqlmock
	activityMYSQL repository.ActivityRepository
}

func (t *activityRepositoryMysqlTest) TestFetch() {
	activity := aggregate.RebuildActivity(1, "bagus@bagus.com", "kerja bro")
	query := `SELECT id, email, title FROM activities`
	t.Run("success", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"email",
			"title",
		}).AddRow(
			activity.ID,
			activity.Email,
			activity.Title,
		)

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		actualActivity, err := t.activityMYSQL.GetActivity()

		t.NotNil(actualActivity)
		t.NoError(err)
	})
	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New(""))

		actualActivity, err := t.activityMYSQL.GetActivity()

		t.Nil(actualActivity)
		t.Error(err)
	})
}

func (t *activityRepositoryMysqlTest) TestCreate() {
	activity := aggregate.RebuildActivity(1, "bagus@bagus.com", "kerja bro")

	t.Run("success", func() {
		t.mock.ExpectExec("INSERT INTO activities").WithArgs(
			activity.ID,
			activity.Email,
			activity.Title,
		).WillReturnResult(sqlmock.NewResult(1, 1))
		t.NoError(t.activityMYSQL.CreateActivity(activity))
	})

	t.Run("failed created", func() {
		t.mock.ExpectExec("INSERT INTO activities").WithArgs(
			activity.ID,
			activity.Email,
			activity.Title,
		).WillReturnError(errors.New(""))
		t.Error(t.activityMYSQL.CreateActivity(activity))
	})

}

func TestActivityRepositoryMySQL(t *testing.T) {
	db, mock, _ := sqlmock.New()

	suite.Run(t, &activityRepositoryMysqlTest{
		mock:          mock,
		activityMYSQL: NewMysqlActivityRepository(db),
	})
}
