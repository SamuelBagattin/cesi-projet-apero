package userController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	userRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/User"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	user, err := userRepository.GetAll()
	if err != nil {
		controllers.SendInternalServerError(c, err)
		return
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
		log.Println(err)
		controllers.SendJsonError(c)
		return
	}

	err = userRepository.Create(user)
	if err != nil {
		controllers.SendInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusCreated, "")
	return
}

func GetOne(c *gin.Context) {
	idUser := c.Param("id")
	intId, err := strconv.Atoi(idUser)

	if err != nil {
		log.Println(err)
		controllers.SendIntegerParsingError(c, idUser)
		return
	}

	oneUser, err := userRepository.GetOneUser(intId)
	c.JSON(200, oneUser)
}
