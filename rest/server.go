package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "models"
	database "utils"
)

func response(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func aUserReqeust(c *gin.Context) {
	var requestQuery models.AUser
	c.ShouldBindQuery(&requestQuery)
	result := database.AUserProcess(requestQuery)

	response(c, result)
}

func allUserReqeust(c *gin.Context) {
	result := database.AllUserDBProcess()

	response(c, result)
}

func addUserReqeust(c *gin.Context) {

	var requestBody models.AddUser
	c.ShouldBindJSON(&requestBody)

	result := database.AddUserProcess(requestBody)

	response(c, result)
}

func deleteUserReqeust(c *gin.Context) {
	var requestQuery models.DeleteUser
	c.ShouldBindQuery(&requestQuery)

	result := database.DeleteUserProcess(requestQuery)

	response(c, result)
}

func updateUserReqeust(c *gin.Context) {
	var requestBody models.UpdateUser
	c.ShouldBindJSON(&requestBody)

	result := database.UpdateUserProcess(requestBody)

	response(c, result)
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
