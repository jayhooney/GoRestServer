package server

import (
	"net/http"

	models "../Define/Models"
	database "../Util"
	"github.com/gin-gonic/gin"
)

func aUserReqeust(c *gin.Context) {
	var requestQuery models.AUser
	c.ShouldBindQuery(&requestQuery)
	result := database.AUserProcess(requestQuery.UserID)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func allUserReqeust(c *gin.Context) {
	result := database.AllUserDBProcess()

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func addUserReqeust(c *gin.Context) {

	var requestjson models.AddUser
	c.ShouldBindJSON(&requestjson)

	result := database.AddUserProcess(requestjson.UserName)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func deleteUserReqeust(c *gin.Context) {
	var requestQuery models.DeleteUser
	c.ShouldBindQuery(&requestQuery)

	result := database.DeleteUserProcess(requestQuery.UserID)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func updateUserReqeust(c *gin.Context) {
	var requestBody models.UpdateUser
	c.ShouldBindJSON(&requestBody)

	result := database.UpdateUserProcess(requestBody.UserName, requestBody.UserID)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func Run() {
	router := gin.Default()

	router.GET(`/aUser`, aUserReqeust)
	router.GET(`/allUser`, allUserReqeust)
	router.POST(`/addUser`, addUserReqeust)
	router.DELETE(`/deleteUser`, deleteUserReqeust)
	router.PUT(`/updateUser`, updateUserReqeust)

	router.Run(":8080")
}
