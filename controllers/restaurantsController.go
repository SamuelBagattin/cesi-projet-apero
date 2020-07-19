package controllers

import (
	"github.com/SamuelBagattin/cesi-projet-apero/repositories"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	ListRestau := repositories.GetRestaurants()
	c.JSON(200, ListRestau)
}
