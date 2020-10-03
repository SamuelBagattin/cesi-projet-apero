package placesDistrictController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"github.com/SamuelBagattin/cesi-projet-apero/repositories/placesQuartier"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	district, err := restaurantsQuartierRepository.GetPlacesDistrict()
	if err != nil {
		c.JSON(500, err)
	} else {
		if *district == nil {
			c.JSON(http.StatusOK, make([]string, 0))
			return
		}
		c.JSON(200, district)
	}
}

func Create(c *gin.Context) {
	var district models.PlacesDistrict

	err := c.ShouldBindJSON(&district)
	if err != nil {
		panic(err)
	}

	err = restaurantsQuartierRepository.Create(district)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}
