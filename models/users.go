package models

import (
	"database/sql"
)

type User struct {
	GlobUserID 	int `json:"userid"`
	GlobEmail	string `json:"email"`
	GlobAddress	string `json:"address"`
	GlobPassword	string `json:"password"`
}

type TaskCollection struct {
	Users []User `json:"all_users"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM users"
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		user := User{}
		err2 := rows.Scan(&user.GlobUserID, &user.GlobEmail, &user.GlobAddress, &user.GlobPassword)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Users = append(result.Users, user)
	}
	return result
}

func PutTask(db *sql.DB, Email string, Address string, Password string) (int64, error) {
	sql := "INSERT INTO users(Email, Address,Password) VALUES(?,?,?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'Email'
	result, err2 := stmt.Exec(Email,Address,Password)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func EditTask(db *sql.DB, userId int, Email string, Address string, Password string) (int64, error) {
	sql := "UPDATE users set Email = ?, Address = ?, Password = ? WHERE UserID = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(Email, Address,Password, userId)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func DeleteTask(db *sql.DB, UserID int) (int64, error) {
	sql := "DELETE FROM users WHERE UserID = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'UserID'
	result, err2 := stmt.Exec(UserID)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
