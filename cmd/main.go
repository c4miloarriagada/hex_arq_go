package main

import (
	"fmt"
	"log"

	"github.com/c4miloarriagada/hexarq/cmd/db"
	"github.com/gin-gonic/gin"
)

func main() {
	connectDb()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func connectDb() {
	db, err := db.GetDB()

	if err != nil {
		log.Fatalf("Error trying to connect: %v", err)
	}

	pg, err := db.DB()
	if err != nil {
		log.Fatalf("Error trying to get db connection: %v", err)
	}

	err = pg.Ping()

	if err != nil {
		log.Fatalf("Cant establish ping to db: %v", err)
	} else {
		fmt.Println("Successfully conecction to db")
	}
}
