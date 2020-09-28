package placesController

import (
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	placesRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/places"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAll(c *gin.Context) {
	PlacesList := placesRepository.GetPlaces()
	c.JSON(200, PlacesList)
}

func GetOne(c *gin.Context) {
	IdPlace := c.Param("id")
	OnePlace := placesRepository.GetOnePlace(IdPlace)
	c.JSON(200, OnePlace)
}

func Create(c *gin.Context) {
	var place models.Place

	err := c.ShouldBindJSON(&place)
	if err != nil {
		panic(err)
	}

	err = placesRepository.Create(place)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")

	/*
		c.JSON(http.StatusOK, gin.H{
			"note":         place.Note,
			"appreciation": place.Appreciation,
			"prixmoyen":    place.Prixmoyen,
			"adresse":      place.Adresse,
			"ville":        place.Ville,
			//"datecreation":    place.Datecreation,
			"nom":             place.Nom,
			"quartierid":      place.QuartierId,
			"categorieid":     place.CategorieId,
			"notecopiosite":   place.NoteCopiosite,
			"notedeliciosite": place.NoteDeliciosite,
			"notecadre":       place.NoteCadre,
			"noteaccueil":     place.NoteAccueil,
		})*/

}

func Update(c *gin.Context) {
	var rest models.Place

	err := c.ShouldBindJSON(&rest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = placesRepository.Update(rest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")

}
