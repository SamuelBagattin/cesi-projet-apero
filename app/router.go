package main

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers/restaurants"
	"github.com/SamuelBagattin/cesi-projet-apero/controllers/restaurantsCategory"
	"github.com/SamuelBagattin/cesi-projet-apero/controllers/restaurantsQuartier"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func InitalizeRouter() {
	r := gin.Default()
	r.Use(cors.Default())
	restaurantsGroup := r.Group("/restaurants")
	{
		restaurantsC := restaurantsController.GetDefaultRestaurantController()
		restaurantsGroup.GET("", restaurantsC.GetAll)
		restaurantsGroup.GET("/:id", restaurantsC.GetOne)
		restaurantsGroup.POST("", restaurantsC.Create)
		restaurantsGroup.PUT("", restaurantsC.Update)
	}
	categoryGroup := r.Group("/restaurantCategories")
	{
		categoryGroup.GET("", restaurantsCategoryController.GetAll)
		categoryGroup.POST("", restaurantsCategoryController.Create)
	}

	quartierGroup := r.Group("/restaurantQuartier")
	{
		quartierGroup.GET("", restaurantsQuartierController.GetAll)
		quartierGroup.POST("", restaurantsQuartierController.Create)
	}
	var err = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}
