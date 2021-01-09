package happyHourController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers"
	"github.com/SamuelBagattin/cesi-projet-apero/core"
	"github.com/SamuelBagattin/cesi-projet-apero/custom_errors"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	happyHourRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/happyHour"
	placesRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/places"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	var happy models.HappyHour

	err := c.ShouldBindJSON(&happy)

	if err != nil {
		log.Println(err)
		controllers.SendJsonError(c)
	}

	err = happyHourRepository.Create(happy)
	if err != nil {
		controllers.SendInternalServerError(c, err)
		return
	}
	controllers.SendEntityCreatedResponse(c)
	return
}

func GetAll(c *gin.Context) {
	include := c.Query("include")
	knownIncludes := []string{"user"}
	includeFields, err := controllers.GetIncludeFields(include, knownIncludes)
	if err != nil {
		controllers.SendBadRequestError(c, err)
		return
	}
	if includeFields != nil && !core.ArrayContainsString(includeFields, knownIncludes[0]) {
		happyHourList, err := happyHourRepository.GetAll()
		if err != nil {
			log.Println(err)
			controllers.SendInternalServerError(c, err)
			return
		}
		if *happyHourList == nil {
			c.JSON(http.StatusOK, make([]string, 0))
			return
		}
		c.JSON(http.StatusOK, happyHourList)
		return

	} else {
		happyHourList, err := happyHourRepository.GetAllWithCreator()
		if err != nil {
			log.Println(err)
			controllers.SendInternalServerError(c, err)
			return
		}
		if *happyHourList == nil {
			c.JSON(http.StatusOK, make([]string, 0))
			return
		}
		c.JSON(http.StatusOK, happyHourList)
		return
	}

}

func GetOne(c *gin.Context) {
	idPlace := c.Param(controllers.IdQueryparam)
	idPlaceInt, parseError := strconv.Atoi(idPlace)
	if parseError != nil {
		controllers.SendIntegerParsingError(c, idPlace)
		return
	}
	onePlace, err := placesRepository.GetOnePlace(idPlaceInt)
	if err != nil {
		switch v := err.(type) {

		case custom_errors.EntityNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": v.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			log.Println(err)
			return
		}

	}
	c.JSON(http.StatusOK, onePlace)
}
