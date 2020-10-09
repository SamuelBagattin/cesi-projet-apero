package happyHourController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	happyHourRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/happyHour"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Create(c *gin.Context) {
	user := c.Param("user")
	if user == "noUser" {
		var happy models.HappyHour

		err := c.ShouldBindJSON(&happy)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		err = happyHourRepository.Create(happy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, "")
	}
}

func GetAll(c *gin.Context) {
	include := c.Query("include")
	additionalFields := strings.Split(include, ",")
	usr := "user"
	includeUser := false
	unknownIncl := make([]string, 1)
	for _, element := range additionalFields {
		if element != usr {
			unknownIncl = append(unknownIncl, element)
		} else {
			includeUser = true
		}
	}
	if len(unknownIncl) > 0 {
		errorMessage := "Unknown fields to include: "
		for index, el := range unknownIncl {
			if index == 0 {
				errorMessage = errorMessage + el
			} else {
				errorMessage = ", " + errorMessage + el
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
	}
	if !includeUser {
		happyHourList := happyHourRepository.GetAll()
		if *happyHourList == nil {
			c.JSON(http.StatusOK, make([]string, 0))
			return
		}
		c.JSON(200, happyHourList)

	} else {
		happyHourList := happyHourRepository.GetAllWithCreator()
		if *happyHourList == nil {
			c.JSON(http.StatusOK, make([]string, 0))
			return
		}
		c.JSON(200, happyHourList)
	}

}
