package main

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers/restaurants"
	"github.com/SamuelBagattin/cesi-projet-apero/controllers/restaurantsCategory"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func InitalizeRouter() {
	r := gin.Default()
	r.Use(cors.Default())
	restaurantsGroup := r.Group("/restaurants")
	{
		restaurantsGroup.GET("", restaurantsController.GetAll)
		restaurantsGroup.GET("/:id", restaurantsController.GetOne)
		restaurantsGroup.POST("", restaurantsController.Create)
		restaurantsGroup.PUT("", restaurantsController.Update)
	}
	categoryGroup := r.Group("/restaurantCategories")
	{
		categoryGroup.GET("", restaurantsCategoryController.GetAll)
		categoryGroup.POST("", restaurantsCategoryController.Create)
	}
	var err = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}
