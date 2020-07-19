package main

import (
	"github.com/SamuelBagattin/cesi-projet-apero/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	restaurantsGroup := r.Group("/restaurants")
	{
		restaurantsGroup.GET("", controllers.GetAll)
	}

	//liste routes :
	/*
		get : entrant : rien : retourne la liste de tous les objets 200
		post restaurants "/restaurants" entrée : 1 modèle restaurant en entier. retourne l'entité avec le bon id .code 201 à vérif
		Pas de delete pour pas supprimer le sondage. Voir pour un archivage.
		put "/restaurants/:id"   :id = variable de path. Entrée : modèle entier. Retourne le nouvel objet 200








	*/

	var err = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}
