package restaurantsRepository

import (
	"errors"
	"fmt"
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

func GetDefaultRestaurantRepo() RestaurantRepoInterface {
	return &RestaurantRepo{}
}

type RestaurantRepoInterface interface {
	GetRestaurants() (*[]*models.Restaurant, error)
	GetOneRestaurant(id string) (*models.Restaurant, error)
	Create(rest models.Restaurant) error
	Update(rest models.Restaurant) error
}

type RestaurantRepo struct {
}

func (r *RestaurantRepo) GetRestaurants() (*[]*models.Restaurant, error) {

	rows, err := config.DatabaseInit().Query("select id, note, appreciation, prixmoyen, adresse, ville, datecreation, nom, notecopiosite, notedeliciosite, notecadre, noteaccueil, quartier_id, categorie_id from restaurant")
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal server error")
	}

	var restaurants []*models.Restaurant

	for rows.Next() {
		restau := models.Restaurant{}
		if err := rows.Scan(&restau.Id, &restau.Note, &restau.Appreciation, &restau.Prixmoyen, &restau.Adresse, &restau.Ville,
			&restau.Datecreation, &restau.Nom, &restau.NoteCopiosite, &restau.NoteDeliciosite, &restau.NoteCadre, &restau.NoteAccueil, &restau.QuartierId, &restau.CategorieId); err != nil {
			log.Println(err)
			return nil, errors.New("internal server error")
		}
		restaurants = append(restaurants, &restau)
	}
	err = rows.Close()
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return &restaurants, nil
}

func (r *RestaurantRepo) GetOneRestaurant(id string) (*models.Restaurant, error) {

	intId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		return nil, errors.New("provided id is not an integer")
	}

	row := config.DatabaseInit().QueryRow("select * from restaurant where id = $1", intId)

	restau := models.Restaurant{}

	if err := row.Scan(&restau.Id, &restau.Note, &restau.Appreciation, &restau.Prixmoyen, &restau.Adresse, &restau.Ville,
		&restau.Datecreation, &restau.Nom, &restau.QuartierId, &restau.CategorieId, &restau.NoteCopiosite, &restau.NoteDeliciosite, &restau.NoteCadre, &restau.NoteAccueil); err != nil {
		log.Println(err)
		if err.Error() == "no rows in result set" {
			formattedError := fmt.Sprintf("restaurant id %d doesn't exist", intId)
			return nil, errors.New(formattedError)
		}

		return nil, errors.New("internal server error")
	}

	return &restau, nil
}

func (r *RestaurantRepo) Create(rest models.Restaurant) error {

	_, err := config.DatabaseInit().Exec("insert into restaurant(nom, appreciation, quartier_id, categorie_id, note, prixmoyen, adresse, ville, notecopiosite, notedeliciosite, notecadre, noteaccueil, datecreation) values ($1,$2,$3,$4, $5, $6, $7, $8, $9, $10, $11, $12, current_date)",
		rest.Nom, rest.Appreciation, rest.QuartierId, rest.CategorieId, rest.Note, rest.Prixmoyen, rest.Adresse, rest.Ville, rest.NoteCopiosite, rest.NoteDeliciosite, rest.NoteCadre, rest.NoteAccueil)

	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	return nil

}

func (r *RestaurantRepo) Update(rest models.Restaurant) error {

	_, err := config.DatabaseInit().Exec("UPDATE restaurant set nom = $1,  appreciation = $2,  quartier_id = $3,  categorie_id = $4, prixmoyen= $5,  adresse = $6, ville=$7, notecopiosite=$8, notedeliciosite = $9, notecadre =$10, noteaccueil = $11  WHERE id = $12",
		rest.Nom, rest.Appreciation, rest.QuartierId, rest.CategorieId, rest.Prixmoyen, rest.Adresse, rest.Ville, rest.NoteCopiosite, rest.NoteDeliciosite, rest.NoteCadre, rest.NoteAccueil, rest.Id)

	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	return nil

}
