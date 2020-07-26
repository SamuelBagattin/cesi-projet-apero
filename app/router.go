package main

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func InitalizeRouter() {
	r := gin.Default()
	r.Use(cors.Default())
	restaurantsGroup := r.Group("/restaurants")
	{
		restaurantsGroup.GET("", controllers.GetAll)
		restaurantsGroup.GET("/:id", controllers.GetOne)
		restaurantsGroup.POST("", controllers.Create)
	}
	var err = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}
