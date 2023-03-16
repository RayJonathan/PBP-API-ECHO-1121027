package models

import (
	"database/sql"
	"fmt"
	"net/http"

	"example.com/m/db"
	"example.com/m/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {

	var obj User
	var pwd string

	fmt.Println("password", password)

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ? "

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username Not Found ")
		return false, err
	}
	if err != nil {
		fmt.Println("Query Errors")
		return false, err
	}

	fmt.Println("pwd : ", pwd)

	match, err := helpers.CheckPasswordHash(password, pwd)
	fmt.Println("match : ", match)

	if !match {

		fmt.Println("HASH And password doesn't match.")
		return false, err

	}
	return true, nil

}

func StoreUsers(username string, password string) (Response, error) {

	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT users (username, password) VALUES ( ?, ?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	hashpwd, _ := helpers.HashPassword(password)
	fmt.Print(hashpwd)

	result, err := stmt.Exec(username, hashpwd)

	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}
