package main

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers/Places"
	"github.com/SamuelBagattin/cesi-projet-apero/controllers/PlacesDistrict"
	userController "github.com/SamuelBagattin/cesi-projet-apero/controllers/User"
	"github.com/SamuelBagattin/cesi-projet-apero/controllers/placesCategory"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func InitalizeRouter() {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowMethods = []string{"*"}
	corsConfig.MaxAge = 24 * time.Hour
	r.Use(cors.New(corsConfig))
	restaurantsGroup := r.Group("/places")
	{
		restaurantsGroup.GET("", placesController.GetAll)
		restaurantsGroup.GET("/:id", placesController.GetOne)
		restaurantsGroup.POST("", placesController.Create)
		restaurantsGroup.PUT("", placesController.Update)
	}
	categoryGroup := r.Group("/placeCategories")
	{
		categoryGroup.GET("", placesCategoryController.GetAll)
		categoryGroup.POST("", placesCategoryController.Create)
	}

	quartierGroup := r.Group("/placeDistricts")
	{
		quartierGroup.GET("", placesDistrictController.GetAll)
		quartierGroup.POST("", placesDistrictController.Create)
	}

	happyHourGroup := r.Group("/happyHours")
	{
		happyHourGroup.GET("", happyHourController.GetAll)
		happyHourGroup.POST("", happyHourController.Create)
	}

	userGroup := r.Group("/users")
	{
		userGroup.GET("", userController.GetAll)
		userGroup.GET("/:id", userController.GetOneUser)
		userGroup.POST("", userController.Create)
	}

	var err = r.Run()
	if err != nil {
		log.Fatal(err)
	}

}
