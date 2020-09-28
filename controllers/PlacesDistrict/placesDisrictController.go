package placesDistrictController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"github.com/SamuelBagattin/cesi-projet-apero/repositories/placesQuartier"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	quartier, err := restaurantsQuartierRepository.GetPlacesDistrict()
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, quartier)
	}
}

func Create(c *gin.Context) {
	var quartier models.PlacesDistrict

	err := c.ShouldBindJSON(&quartier)
	if err != nil {
		panic(err)
	}

	err = restaurantsQuartierRepository.Create(quartier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}
