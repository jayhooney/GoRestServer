package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	models "../models"
	secret "../secret"
)

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func AUserProcess(item models.AUser) string {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	checkErr(err)
	defer db.Close()

	var userName string
	db.QueryRow(`SELECT USER_NM FROM USER_TB WHERE USER_ID = ?;`, item.UserID).Scan(&userName)

	return userName
}

func AllUserDBProcess() []models.UserInfo {
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

	return userSlice

}

func AddUserProcess(item models.AddUser) int64 {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	checkErr(err)
	defer db.Close()

	result, err := db.Exec(`INSERT INTO USER_TB (USER_NM) VALUES (?);`, item.UserName)
	checkErr(err)

	affectedRows, err := result.LastInsertId()
	checkErr(err)

	return affectedRows

}

func DeleteUserProcess(item models.DeleteUser) int64 {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	checkErr(err)
	defer db.Close()

	result, err := db.Exec(`DELETE FROM USER_TB WHERE USER_ID = ?;`, item.UserID)
	checkErr(err)

	affectedRow, _ := result.RowsAffected()

	return affectedRow
}

func UpdateUserProcess(item models.UpdateUser) int64 {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	checkErr(err)
	defer db.Close()

	result, err := db.Exec(`UPDATE USER_TB SET USER_NM = ? WHERE USER_ID = ?;`, item.UserName, item.UserID)
	checkErr(err)

	affectedRows, err := result.RowsAffected()
	checkErr(err)

	return affectedRows

}
