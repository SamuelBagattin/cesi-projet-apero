package restaurantsController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"github.com/SamuelBagattin/cesi-projet-apero/repositories/restaurants"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetDefaultRestaurantController() RestaurantsControllerInterface {
	return &RestaurantController{
		restaurantsRepository.GetDefaultRestaurantRepo(),
	}
}

type RestaurantsControllerInterface interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type RestaurantController struct {
	restaurantRepo restaurantsRepository.RestaurantRepoInterface
}

func (r *RestaurantController) GetAll(c *gin.Context) {
	ListRestau, err := r.restaurantRepo.GetRestaurants()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, ListRestau)
}

func (r *RestaurantController) GetOne(c *gin.Context) {
	IdRestau := c.Param("id")
	OneRestau, err := r.restaurantRepo.GetOneRestaurant(IdRestau)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, OneRestau)
}

func (r *RestaurantController) Create(c *gin.Context) {
	var rest models.Restaurant

	err := c.ShouldBindJSON(&rest)
	if err != nil {
		panic(err)
	}

	err = r.restaurantRepo.Create(rest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (r *RestaurantController) Update(c *gin.Context) {
	var rest models.Restaurant

	err := c.ShouldBindJSON(&rest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = r.restaurantRepo.Update(rest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, nil)

}
