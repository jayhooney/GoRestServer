package database

import (
	"database/sql"
	"log"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	models "models"
	secret "secret"
)

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func getConn() *sql.DB {
	db, err := sql.Open(secret.GetEngine(), secret.GetDBInfo())
	checkErr(err)

	// 단순히 MaxOpenConns 값만 올려준다고 해도
	// MaxIdleConns 값이 디폴트 2 이므로
	// 생성은 10개를 하더라도, 결국 재활용을 2개 밖에 할 수없다.
	// 이를 해결해서 실질적으로 성능 향상을 기대하려면
	// MaxIdleConns를 MaxOpenConns와 같거나 비슷하게 설정해야한다.
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	// MySQL 서버는 길게 유지하고 있는 커넥션을 강제로 끊는다.
	// 이를 해결하기 위해서 커넥션 풀의 ConnMaxLifetime 값을
	// MySQL서버의 wait_time보다 작게 설정하여 해결가능하다.
	db.SetConnMaxLifetime(time.Hour)

	// defer db.Close()

	return db
}

// AUserProcess returns string
func AUserProcess(item models.AUser) string {
	db := getConn()
	defer db.Close()

	var userName string
	db.QueryRow(`SELECT USER_NM FROM USER_TB WHERE USER_ID = ?;`, item.UserID).Scan(&userName)

	return userName
}

// AllUserDBProcess returns  []models.UserInfo
func AllUserDBProcess() []models.UserInfo {
	db := getConn()
	defer db.Close()

	var userSlice []models.UserInfo
	userRows, err := db.Query(`SELECT USER_ID, USER_NM FROM USER_TB;`)
	checkErr(err)

	for userRows.Next() {
		var userInfo models.UserInfo
		err := userRows.Scan(&userInfo.ID, &userInfo.Name)
		checkErr(err)

		userSlice = append(userSlice, userInfo)
	}

	return userSlice

}

// AddUserProcess returns int64
func AddUserProcess(item models.AddUser) int64 {
	db := getConn()
	defer db.Close()

	result, err := db.Exec(`INSERT INTO USER_TB (USER_NM) VALUES (?);`, item.UserName)
	checkErr(err)

	affectedRows, err := result.LastInsertId()
	checkErr(err)

	return affectedRows

}

// DeleteUserProcess returns int64
func DeleteUserProcess(item models.DeleteUser) int64 {
	db := getConn()
	defer db.Close()

	result, err := db.Exec(`DELETE FROM USER_TB WHERE USER_ID = ?;`, item.UserID)
	checkErr(err)

	affectedRow, _ := result.RowsAffected()

	return affectedRow
}

// UpdateUserProcess returns int64
func UpdateUserProcess(item models.UpdateUser) int64 {
	db := getConn()
	defer db.Close()

	result, err := db.Exec(`UPDATE USER_TB SET USER_NM = ? WHERE USER_ID = ?;`, item.UserName, item.UserID)
	checkErr(err)

	affectedRows, err := result.RowsAffected()
	checkErr(err)

	return affectedRows
}
