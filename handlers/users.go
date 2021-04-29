package handlers

import (
	"database/sql"
	"net/http"
	"api/models"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var user models.User

		c.Bind(&user)

		UserID, err := models.PutTask(db, user.GlobEmail, user.GlobAddress, user.GlobPassword)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": UserID,
			})
		} else {
			return err
		}

	}
}

func EditTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var user models.User
		c.Bind(&user)

		_, err := models.EditTask(db, user.GlobUserID, user.GlobEmail, user.GlobAddress, user.GlobPassword)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": user,
			})
		} else {
			return err
		}
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID, _ := strconv.Atoi(c.Param("UserID"))

		_, err := models.DeleteTask(db, UserID)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": UserID,
			})
		} else {
			return err
		}

	}
}
