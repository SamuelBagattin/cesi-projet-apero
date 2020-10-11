package happyHourController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers"
	"github.com/SamuelBagattin/cesi-projet-apero/core"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	happyHourRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/happyHour"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Create(c *gin.Context) {
	var happy models.HappyHour

	err := c.ShouldBindJSON(&happy)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
	include := c.Query("include")
	knownIncludes := []string{"user"}
	includeFields, err := controllers.GetIncludeFields(include, knownIncludes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if includeFields != nil && !core.ArrayContainsString(includeFields, knownIncludes[0]) {
		happyHourList := happyHourRepository.GetAll()
		if *happyHourList == nil {
			c.JSON(http.StatusOK, make([]string, 0))
			return
		}
		c.JSON(200, happyHourList)
		return

	} else {
		happyHourList := happyHourRepository.GetAllWithCreator()
		if *happyHourList == nil {
			c.JSON(http.StatusOK, make([]string, 0))
			return
		}
		c.JSON(200, happyHourList)
		return
	}

}
