package placesCategoryController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"github.com/SamuelBagattin/cesi-projet-apero/repositories/placesCategory"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAll(c *gin.Context) {
	categories, err := placesCategoryRepository.GetPlacesCategories()
	if err != nil {
		log.Println(err)
		controllers.SendInternalServerError(c, err)
		return
	}
	if *categories == nil {
		c.JSON(http.StatusOK, make([]string, 0))
		return
	}
	c.JSON(http.StatusOK, categories)
	return
}

func Create(c *gin.Context) {
	var category models.PlaceCategory

	err := c.ShouldBindJSON(&category)
	if err != nil {
		log.Println(err)
		controllers.SendJsonError(c)
		return
	}

	err = placesCategoryRepository.Create(category)
	if err != nil {
		controllers.SendInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusCreated, "")
	return
}
