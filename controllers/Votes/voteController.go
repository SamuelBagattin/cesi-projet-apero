package voteController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	voteRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/votes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAll(c *gin.Context) {
	vote, err := voteRepository.GetAll()
	if err != nil {
		c.JSON(500, err)
	} else {
		if *vote == nil {
			c.JSON(http.StatusOK, make([]string, 0))
			return
		}
		c.JSON(200, vote)
	}
}

func Create(c *gin.Context) {

	var vote models.Vote

	err := c.ShouldBindJSON(&vote)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = voteRepository.Create(vote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")
}
