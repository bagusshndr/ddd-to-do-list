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

type activityRepositoryMysqlTest struct {
	suite.Suite
	mock          sqlmock.Sqlmock
	activityMYSQL repository.ActivityRepository
}

func (t *activityRepositoryMysqlTest) TestGetActivity() {
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

		actualActivity, err := t.activityMYSQL.GetActivity(10)

		t.NotNil(actualActivity)
		t.NoError(err)
	})
	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New(""))

		actualActivity, err := t.activityMYSQL.GetActivity(10)

		t.Nil(actualActivity)
		t.Error(err)
	})
}

func (t *activityRepositoryMysqlTest) TestGetActivityByID() {
	activity := aggregate.RebuildActivity(1, "bagus@bagus.com", "kerja bro")
	query := `SELECT id, email, title FROM activities WHERE id = ?`
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

		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(activity.ID).WillReturnRows(rows)

		actualActivity, err := t.activityMYSQL.GetActivityByID(activity.ID)

		log.Println(actualActivity)
		t.NotNil(actualActivity)
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(activity.ID).WillReturnError(errors.New(""))

		actualActivity, err := t.activityMYSQL.GetActivityByID(activity.ID)

		log.Println(actualActivity)
		t.Nil(actualActivity)
		t.Error(err)
	})

}

func (t *activityRepositoryMysqlTest) TestCreate() {
	t.Run("success", func() {
		t.mock.ExpectExec("INSERT INTO activities").WillReturnResult(sqlmock.NewResult(1, 1))
		_, err := t.activityMYSQL.CreateActivity("email@gmail.com", "title")
		t.NoError(err)
	})

	t.Run("failed", func() {
		t.mock.ExpectExec("INSERT INTO activities").WillReturnError(errors.New("failed store activity"))
		_, err := t.activityMYSQL.CreateActivity("email@gmail.com", "title")
		t.Error(err)
	})
}

// func (t *activityRepositoryMysqlTest) TestUpdate() {
// 	queryUpdate := "UPDATE activities SET email = ?, title = ? WHERE id = ?"

// 	t.Run("success", func() {
// 		t.mock.ExpectExec(queryUpdate).WithArgs("email@gmail.com", "title", 1).WillReturnResult(sqlmock.NewResult(1, 1))

// 		t.NoError(t.activityMYSQL.UpdateActivity(1, "email@gmail.com", "title"))
// 	})

// 	// t.Run("failed", func() {
// 	// 	t.mock.ExpectExec(queryUpdateAwb).WithArgs("ABC123").WillReturnError(errors.New("failed update awb"))

// 	// 	t.Error(t.awbRepositoryMySQL.UpdateAwb("ABC123"))
// 	// })
// }

func (t *activityRepositoryMysqlTest) TestDelete() {
	queryDelete := "DELETE FROM activities WHERE id = ?"

	t.Run("success", func() {
		t.mock.ExpectExec(queryDelete).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		t.NoError(t.activityMYSQL.DeleteActivity(1))
	})

	t.Run("failed", func() {
		t.mock.ExpectExec(queryDelete).WillReturnError(errors.New("failed delete activity"))

		t.Error(t.activityMYSQL.DeleteActivity(1))
	})
}

func TestActivityRepositoryMySQL(t *testing.T) {
	db, mock, _ := sqlmock.New()

	suite.Run(t, &activityRepositoryMysqlTest{
		mock:          mock,
		activityMYSQL: NewMysqlActivityRepository(db),
	})
}
