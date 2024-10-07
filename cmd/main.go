package main

import (
	"github.com/c4miloarriagada/hexarq/cmd/api/handlers/user"
	"github.com/c4miloarriagada/hexarq/cmd/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDb()
	db.CreateTables()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/users", user.CreateUser)
	r.Run()
}
