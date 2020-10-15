package voteController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	voteRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/votes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAll(c *gin.Context) {
	vote, err := voteRepository.GetAll()
	if err != nil {
		controllers.SendInternalServerError(c, err)
		return
	}
	if *vote == nil {
		c.JSON(http.StatusOK, make([]string, 0))
		return
	}
	c.JSON(200, vote)
}

func Create(c *gin.Context) {

	var vote models.Vote

	err := c.ShouldBindJSON(&vote)
	if err != nil {
		log.Println(err)
		controllers.SendJsonError(c)
	}

	err = voteRepository.Create(vote)
	if err != nil {
		controllers.SendInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusCreated, "")
}
