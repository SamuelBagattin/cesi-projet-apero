package userController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	userRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/User"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	user, err := userRepository.GetAll()
	if err != nil {
		c.JSON(500, err)
	} else {
		if *user == nil {
			c.JSON(http.StatusOK, make([]string, 0))
			return
		}
		c.JSON(200, user)
	}
}

func Create(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	err = userRepository.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}

func GetOneUser(c *gin.Context) {
	idUser := c.Param("id")
	oneUser := userRepository.GetOneUser(idUser)
	c.JSON(200, oneUser)
}
