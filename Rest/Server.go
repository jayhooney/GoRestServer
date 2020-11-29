package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.GET(`/get`, func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"msg": "GET",
		// })
		fmt.Printf("GET")
	})

	router.POST(`/post`, func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"msg": "POST",
		// })
		fmt.Printf("POST")
	})

	router.DELETE(`/delete`, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "DELETE",
		})
	})

	router.PUT(`/put`, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "PUT",
		})
	})

	router.Run(":8080")
}
