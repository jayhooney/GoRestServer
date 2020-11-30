package database

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"

	models "../Define/Models"
	secret "../Define/Secret"
)

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func AUserProcess(userID int) string {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	checkErr(err)
	defer db.Close()

	var userName string
	db.QueryRow(`SELECT USER_NM FROM USER_TB WHERE USER_ID = ?;`, userID).Scan(&userName)

	return userName
}

func AllUserDBProcess() string {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	checkErr(err)
	defer db.Close()

	var userSlice []models.UserInfo
	userRows, err := db.Query(`SELECT USER_ID, USER_NM FROM USER_TB;`)

	for userRows.Next() {
		var userInfo models.UserInfo
		err := userRows.Scan(&userInfo.ID, &userInfo.Name)
		checkErr(err)

		userSlice = append(userSlice, userInfo)
	}

	result, _ := json.Marshal(userSlice)

	return string(result)

}

func AddUserProcess(userName string) int64 {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	checkErr(err)
	defer db.Close()

	result, err := db.Exec(`INSERT INTO USER_TB (USER_NM) VALUES (?);`, userName)
	checkErr(err)

	affectedRows, err := result.LastInsertId()
	checkErr(err)

	return affectedRows

}

func DeleteUserProcess(userID int) int64 {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	checkErr(err)
	defer db.Close()

	result, err := db.Exec(`DELETE FROM USER_TB WHERE USER_ID = ?;`, userID)
	checkErr(err)

	affectedRow, _ := result.RowsAffected()

	return affectedRow
}
