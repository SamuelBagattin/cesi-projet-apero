package controllers

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"github.com/SamuelBagattin/cesi-projet-apero/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAll(c *gin.Context) {
	ListRestau := repositories.GetRestaurants()
	c.JSON(200, ListRestau)
}

func GetOne(c *gin.Context) {
	IdRestau := c.Param("id")
	OneRestau := repositories.GetOneRestaurant(IdRestau)
	c.JSON(200, OneRestau)
}

func Create(c *gin.Context) {
	var rest models.Restaurant

	err := c.ShouldBindJSON(&rest)
	if err != nil {
		panic(err)
	}

	err = repositories.Create(rest)
	if err != nil {
		log.Print("osuergzorfgjzofgjizoefijzoeifjzoeifjzoifjozeifjoziejfozeifjozeijfozeijfzoeifj")
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
