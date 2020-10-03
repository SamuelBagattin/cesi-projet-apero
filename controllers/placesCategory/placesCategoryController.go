package placesCategoryController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"github.com/SamuelBagattin/cesi-projet-apero/repositories/placesCategory"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	categories, err := placesCategoryRepository.GetPlacesCategories()
	if err != nil {
		c.JSON(500, err)
	} else {
		if *categories == nil {
			c.JSON(http.StatusOK, make([]string, 0))
			return
		}
		c.JSON(200, categories)
	}
}

func Create(c *gin.Context) {
	var category models.PlaceCategory

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = placesCategoryRepository.Create(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}
