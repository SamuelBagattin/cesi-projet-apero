package placesRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/custom_errors"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	_ "github.com/lib/pq"
)

func GetPlaces() (*[]*models.Place, error) {

	rows, err := config.DatabaseInit().Query("select id, (noteaccueil + notecadre + notedeliciosite + notecopiosite)::decimal/4 as note, appreciation, prixmoyen, adresse, ville, datecreation, nom, notecopiosite, notedeliciosite, notecadre, noteaccueil, quartier_id, categorie_id from endroit order by datecreation desc")
	if err != nil {
		return nil, err
	}

	var places []*models.Place

	for rows.Next() {
		place := models.Place{}
		if err := rows.Scan(&place.Id, &place.Note, &place.Appreciation, &place.Prixmoyen, &place.Adresse, &place.Ville,
			&place.Datecreation, &place.Nom, &place.NoteCopiosite, &place.NoteDeliciosite, &place.NoteCadre, &place.NoteAccueil, &place.QuartierId, &place.CategorieId); err != nil {
			return nil, err
		}
		places = append(places, &place)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return &places, nil
}

func GetOnePlace(id int) (*models.Place, error) {

	row := config.DatabaseInit().QueryRow("select id, appreciation, prixmoyen, adresse, ville, datecreation, nom, quartier_id, categorie_id, notecopiosite, notedeliciosite, notecadre, noteaccueil from endroit where id = $1", id)

	place := models.Place{}

	if err := row.Scan(&place.Id, &place.Appreciation, &place.Prixmoyen, &place.Adresse, &place.Ville,
		&place.Datecreation, &place.Nom, &place.QuartierId, &place.CategorieId, &place.NoteCopiosite, &place.NoteDeliciosite, &place.NoteCadre, &place.NoteAccueil); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, custom_errors.EntityNotFound{
				Id: id,
			}
		}
		return nil, err
	}

	return &place, nil
}

func Create(place models.Place) error {

	_, err := config.DatabaseInit().Exec("insert into endroit(nom, appreciation, quartier_id, categorie_id, prixmoyen, adresse, ville, notecopiosite, notedeliciosite, notecadre, noteaccueil, datecreation) values ($1,$2,$3,$4, $5, $6, $7, $8, $9, $10, $11, current_date)",
		place.Nom, place.Appreciation, place.QuartierId, place.CategorieId, place.Prixmoyen, place.Adresse, place.Ville, place.NoteCopiosite, place.NoteDeliciosite, place.NoteCadre, place.NoteAccueil)

	if err != nil {
		return err
	}
	return nil

}

func Update(place models.Place) error {

	_, err := config.DatabaseInit().Exec("UPDATE endroit set nom = $1,  appreciation = $2,  quartier_id = $3,  categorie_id = $4, prixmoyen= $5,  adresse = $6, ville=$7, notecopiosite=$8, notedeliciosite = $9, notecadre =$10, noteaccueil = $11  WHERE id = $12",
		place.Nom, place.Appreciation, place.QuartierId, place.CategorieId, place.Prixmoyen, place.Adresse, place.Ville, place.NoteCopiosite, place.NoteDeliciosite, place.NoteCadre, place.NoteAccueil, place.Id)

	if err != nil {
		return err
	}
	return nil

}
