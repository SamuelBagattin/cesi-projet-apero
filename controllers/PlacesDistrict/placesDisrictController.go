package placesDistrictController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"github.com/SamuelBagattin/cesi-projet-apero/repositories/placesQuartier"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAll(c *gin.Context) {
	district, err := restaurantsQuartierRepository.GetPlacesDistrict()
	if err != nil {
		log.Println(err)
		controllers.SendInternalServerError(c, err)
		return
	}
	if *district == nil {
		c.JSON(http.StatusOK, make([]string, 0))
		return
	}
	c.JSON(200, district)
	return
}

func Create(c *gin.Context) {
	var district models.PlacesDistrict

	err := c.ShouldBindJSON(&district)
	if err != nil {
		log.Println(err)
		controllers.SendJsonError(c)
		return
	}

	err = restaurantsQuartierRepository.Create(district)
	if err != nil {
		controllers.SendInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusCreated, "")
	return
}
