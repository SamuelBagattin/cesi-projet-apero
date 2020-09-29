package placesRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

func GetPlaces() *[]*models.Place {

	rows, err := config.DatabaseInit().Query("select id, note, appreciation, prixmoyen, adresse, ville, datecreation, nom, notecopiosite, notedeliciosite, notecadre, noteaccueil, quartier_id, categorie_id from endroit")
	if err != nil {
		log.Fatal(err)
	}

	var places []*models.Place

	for rows.Next() {
		place := models.Place{}
		if err := rows.Scan(&place.Id, &place.Note, &place.Appreciation, &place.Prixmoyen, &place.Adresse, &place.Ville,
			&place.Datecreation, &place.Nom, &place.NoteCopiosite, &place.NoteDeliciosite, &place.NoteCadre, &place.NoteAccueil, &place.QuartierId, &place.CategorieId); err != nil {
			log.Fatal(err)
		}
		places = append(places, &place)
	}
	err = rows.Close()
	if err != nil {
		panic(err)
	}

	return &places
}

func GetOnePlace(id string) models.Place {

	intId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
	}

	row := config.DatabaseInit().QueryRow("select * from endroit where id = $1", intId)

	place := models.Place{}

	if err := row.Scan(&place.Id, &place.Note, &place.Appreciation, &place.Prixmoyen, &place.Adresse, &place.Ville,
		&place.Datecreation, &place.Nom, &place.QuartierId, &place.CategorieId, &place.NoteCopiosite, &place.NoteDeliciosite, &place.NoteCadre, &place.NoteAccueil); err != nil {
		log.Fatal(err)
	}

	return place
}

func Create(place models.Place) error {

	_, err := config.DatabaseInit().Exec("insert into endroit(nom, appreciation, quartier_id, categorie_id, note, prixmoyen, adresse, ville, notecopiosite, notedeliciosite, notecadre, noteaccueil, datecreation) values ($1,$2,$3,$4, $5, $6, $7, $8, $9, $10, $11, $12, current_date)",
		place.Nom, place.Appreciation, place.QuartierId, place.CategorieId, place.Note, place.Prixmoyen, place.Adresse, place.Ville, place.NoteCopiosite, place.NoteDeliciosite, place.NoteCadre, place.NoteAccueil)

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

func Update(place models.Place) error {

	_, err := config.DatabaseInit().Exec("UPDATE endroit set nom = $1,  appreciation = $2,  quartier_id = $3,  categorie_id = $4, prixmoyen= $5,  adresse = $6, ville=$7, notecopiosite=$8, notedeliciosite = $9, notecadre =$10, noteaccueil = $11  WHERE id = $12",
		place.Nom, place.Appreciation, place.QuartierId, place.CategorieId, place.Prixmoyen, place.Adresse, place.Ville, place.NoteCopiosite, place.NoteDeliciosite, place.NoteCadre, place.NoteAccueil, place.Id)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
