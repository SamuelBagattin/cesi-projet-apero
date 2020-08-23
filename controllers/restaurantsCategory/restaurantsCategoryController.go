package restaurantsCategoryController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"github.com/SamuelBagattin/cesi-projet-apero/repositories/restaurantsCategory"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	categories, err := restaurantsCategoryRepository.GetRestaurantsCategories()
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, categories)
	}
}

func Create(c *gin.Context) {
	var category models.RestaurantCategory

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = restaurantsCategoryRepository.Create(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}
