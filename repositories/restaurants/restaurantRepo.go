package restaurantsRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

func GetRestaurants() *[]*models.Restaurant {

	rows, err := config.DatabaseInit().Query("select * from restaurant")
	if err != nil {
		log.Fatal(err)
	}

	var restaurants []*models.Restaurant

	for rows.Next() {
		restau := models.Restaurant{}
		if err := rows.Scan(&restau.Id, &restau.Note, &restau.Appreciation, &restau.Prixmoyen, &restau.Adresse, &restau.Ville,
			&restau.Datecreation, &restau.Nom, &restau.QuartierId, &restau.CategorieId, &restau.NoteCopiosite, &restau.NoteDeliciosite, &restau.NoteCadre, &restau.NoteAccueil); err != nil {
			log.Fatal(err)
		}
		restaurants = append(restaurants, &restau)
	}
	err = rows.Close()
	if err != nil {
		panic(err)
	}

	return &restaurants
}

func GetOneRestaurant(id string) models.Restaurant {

	intId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
	}

	row := config.DatabaseInit().QueryRow("select * from restaurant where id = $1", intId)

	restau := models.Restaurant{}

	if err := row.Scan(&restau.Id, &restau.Note, &restau.Appreciation, &restau.Prixmoyen, &restau.Adresse, &restau.Ville,
		&restau.Datecreation, &restau.Nom, &restau.QuartierId, &restau.CategorieId, &restau.NoteCopiosite, &restau.NoteDeliciosite, &restau.NoteCadre, &restau.NoteAccueil); err != nil {
		log.Fatal(err)
	}

	return restau
}

func Create(rest models.Restaurant) error {

	_, err := config.DatabaseInit().Exec("insert into restaurant(nom, appreciation, quartier_id, categorie_id, note, prixmoyen, adresse, ville, notecopiosite, notedeliciosite, notecadre, noteaccueil, datecreation) values ($1,$2,$3,$4, $5, $6, $7, $8, $9, $10, $11, $12, current_date)",
		rest.Nom, rest.Appreciation, rest.QuartierId, rest.CategorieId, rest.Note, rest.Prixmoyen, rest.Adresse, rest.Ville, rest.NoteCopiosite, rest.NoteDeliciosite, rest.NoteCadre, rest.NoteAccueil)

	if err != nil {
		log.Println(err)
		return err
	}

	/*var idResult, idError = res.LastInsertId()

	if idError != nil {
		log.Println(idError)
		return 0, idError
	}*/

	return nil

}
