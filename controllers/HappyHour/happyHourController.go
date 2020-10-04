package happyHourController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	happyHourRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/happyHour"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var happy models.HappyHour

	err := c.ShouldBindJSON(&happy)
	if err != nil {
		panic(err)
	}

	err = happyHourRepository.Create(happy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}

func GetAll(c *gin.Context) {

	happyHourList := happyHourRepository.GetAll()
	if *happyHourList == nil {
		c.JSON(http.StatusOK, make([]string, 0))
		return
	}
	c.JSON(200, happyHourList)
}

func GetAllWithCreator(c *gin.Context) {
	happyHourList := happyHourRepository.GetAllWithCreator()
	if *happyHourList == nil {
		c.JSON(http.StatusOK, make([]string, 0))
		return
	}
	c.JSON(200, happyHourList)
}
