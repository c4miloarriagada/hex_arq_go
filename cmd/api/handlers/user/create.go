package user

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" binding:"required"`
	// BornDate time.
}

func CreateUser(c *gin.Context) {}
