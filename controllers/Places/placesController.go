package placesController

import (
	"fmt"
	"github.com/SamuelBagattin/cesi-projet-apero/custom_errors"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	placesRepository "github.com/SamuelBagattin/cesi-projet-apero/repositories/places"
	"strconv"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAll(c *gin.Context) {
	placesList := placesRepository.GetPlaces()
	if *placesList == nil {
		c.JSON(http.StatusOK, make([]string, 0))
		return
	}
	c.JSON(http.StatusOK, placesList)
}

func GetOne(c *gin.Context) {
	idPlace := c.Param("id")
	idPlaceInt, parseError := strconv.Atoi(idPlace)
	if parseError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Place id is not a valid integer : %s", idPlace),
		})
		return
	}
	onePlace, err := placesRepository.GetOnePlace(idPlaceInt)
	if err != nil {
		switch v := err.(type) {

		case custom_errors.EntityNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": v.Error(),
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			log.Println(err)
			return
		}

	}
	c.JSON(http.StatusOK, onePlace)
}

func Create(c *gin.Context) {
	var place models.Place

	err := c.ShouldBindJSON(&place)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
	var place models.Place

	err := c.ShouldBindJSON(&place)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = placesRepository.Update(place)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, "")

}
