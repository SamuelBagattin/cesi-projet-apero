package placesController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers"
	"github.com/SamuelBagattin/cesi-projet-apero/custom_errors"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	placesRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/places"
	"strconv"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAll(c *gin.Context) {
	placesList, err := placesRepository.GetPlaces()
	if err != nil {
		log.Println(err)
		controllers.SendInternalServerError(c, err)
		return
	}
	if *placesList == nil {
		c.JSON(http.StatusOK, make([]string, 0))
		return
	}
	c.JSON(http.StatusOK, placesList)
	return
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

func Create(c *gin.Context) {
	var place models.Place

	err := c.ShouldBindJSON(&place)
	if err != nil {
		log.Println(err)
		controllers.SendJsonError(c)
		return
	}

	err = placesRepository.Create(place)
	if err != nil {
		controllers.SendInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusCreated, "")
	return
}

func Update(c *gin.Context) {
	var place models.Place

	err := c.ShouldBindJSON(&place)
	if err != nil {
		log.Println(err)
		controllers.SendJsonError(c)
		return
	}

	err = placesRepository.Update(place)
	if err != nil {
		controllers.SendInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusCreated, "")
	return
}
