package user

import (
	"log"

	"github.com/c4miloarriagada/hexarq/cmd/internal/domain"
	services "github.com/c4miloarriagada/hexarq/cmd/internal/services/user"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	userCreated, err := services.CreateUserService(user)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	c.JSON(200, gin.H{"user": userCreated})

}
