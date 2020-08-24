package restaurantsController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"github.com/SamuelBagattin/cesi-projet-apero/repositories/restaurants"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	ListRestau := restaurantsRepository.GetRestaurants()
	c.JSON(200, ListRestau)
}

func GetOne(c *gin.Context) {
	IdRestau := c.Param("id")
	OneRestau := restaurantsRepository.GetOneRestaurant(IdRestau)
	c.JSON(200, OneRestau)
}

func Create(c *gin.Context) {
	var rest models.Restaurant

	err := c.ShouldBindJSON(&rest)
	if err != nil {
		panic(err)
	}

	err = restaurantsRepository.Create(rest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")

	/*
		c.JSON(http.StatusOK, gin.H{
			"note":         rest.Note,
			"appreciation": rest.Appreciation,
			"prixmoyen":    rest.Prixmoyen,
			"adresse":      rest.Adresse,
			"ville":        rest.Ville,
			//"datecreation":    rest.Datecreation,
			"nom":             rest.Nom,
			"quartierid":      rest.QuartierId,
			"categorieid":     rest.CategorieId,
			"notecopiosite":   rest.NoteCopiosite,
			"notedeliciosite": rest.NoteDeliciosite,
			"notecadre":       rest.NoteCadre,
			"noteaccueil":     rest.NoteAccueil,
		})*/

}

func Update(c *gin.Context) {
	var rest models.Restaurant

	err := c.ShouldBindJSON(&rest)
	if err != nil {
		panic(err)
	}

	err = restaurantsRepository.Update(rest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")

}
