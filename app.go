package main

import (
	"database/sql"
	"api/handlers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	db := initDB("storage.db")
	migrate(db)
	
	e.GET("/users", handlers.GetTasks(db))
	e.POST("/users", handlers.PutTask(db))
	e.PUT("/users", handlers.EditTask(db))
	e.DELETE("/users/:UserID", handlers.DeleteTask(db))
	
	g := e.Group("")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "EJFA" && password == "123456" {
			return true, nil
		}
		return false, nil
	}))
	
	g.GET("/login", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Authenticated Succes")
	})


	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS users(
        UserID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		Email VARCHAR NOT NULL,
		Address VARCHAR NOT NULL,
		Password VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
