package voteController

import (
	"encoding/json"
	"github.com/SamuelBagattin/cesi-projet-apero/controllers"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	voteRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/votes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	idApero := c.Param(controllers.IdQueryparam)
	idAperoInt, parseError := strconv.Atoi(idApero)
	if parseError != nil {
		controllers.SendIntegerParsingError(c, idApero)
		return
	}
	vote, err := voteRepository.GetAll(idAperoInt)
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

	var vote models.AddVoteRequest
	var voteList []models.AddVoteRequest

	var bodyError error
	var listBodyError error

	raw, _ := c.GetRawData()
	bodyError = json.Unmarshal(raw, &vote)
	if bodyError != nil {
		listBodyError = json.Unmarshal(raw, &voteList)
	}

	//Add single vote
	if bodyError == nil {
		err := voteRepository.Create(vote)
		if err != nil {
			controllers.SendInternalServerError(c, err)
			return
		}
		c.JSON(http.StatusCreated, "")
		return
	} else if listBodyError == nil {
		// add multiple votes
		err := voteRepository.CreateAll(voteList)
		if err != nil {
			controllers.SendInternalServerError(c, err)
			return
		}
		c.JSON(http.StatusCreated, "")
		return

	} else {
		controllers.SendJsonError(c)
	}

}
func Update(c *gin.Context) {

	var vote models.UpdateVoteRequest
	var voteList []models.UpdateVoteRequest

	var bodyError error
	var listBodyError error

	raw, _ := c.GetRawData()
	bodyError = json.Unmarshal(raw, &vote)
	if bodyError != nil {
		listBodyError = json.Unmarshal(raw, &voteList)
	}

	//Add single vote
	if bodyError == nil {
		err := voteRepository.Update(vote)
		if err != nil {
			controllers.SendInternalServerError(c, err)
			return
		}
		c.JSON(http.StatusCreated, "")
		return
	} else if listBodyError == nil {
		// add multiple votes
		err := voteRepository.UpdateAll(voteList)
		if err != nil {
			controllers.SendInternalServerError(c, err)
			return
		}
		c.JSON(http.StatusCreated, "")
		return

	} else {
		controllers.SendJsonError(c)
	}

}
