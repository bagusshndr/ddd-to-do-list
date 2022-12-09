package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	_activityHttpDelivery "ddd-to-do-list/internal/delivery/handler"
	routers "ddd-to-do-list/internal/delivery/router"
	_activityRepo "ddd-to-do-list/internal/infrastructure/database/mysql/repository"
	_activityUcase "ddd-to-do-list/internal/usecase"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
		if err = godotenv.Load(".env"); err != nil {
			return
		}
	}
	dbHost := os.Getenv(`MYSQL_HOST`)
	dbPort := os.Getenv(`MYSQL_PORT`)
	dbUser := os.Getenv(`MYSQL_USER`)
	dbPass := os.Getenv(`MYSQL_PASSWORD`)
	dbName := os.Getenv(`MYSQL_DBNAME`)

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _activityHttpDelivery.InitMiddleware()
	e.Use(middL.CORS)
	ar := _activityRepo.NewMysqlActivityRepository(dbConn)
	au := _activityUcase.NewActivityUsecase(ar)
	tr := _activityRepo.NewMysqlTodoRepository(dbConn)
	tu := _activityUcase.NewTodoUsecase(tr)
	_activityHttpDelivery.NewHandler(au, tu)

	routers.Router(e, au, tu)

	log.Fatal(e.Start(":" + os.Getenv("HTTP_PORT")))
}
