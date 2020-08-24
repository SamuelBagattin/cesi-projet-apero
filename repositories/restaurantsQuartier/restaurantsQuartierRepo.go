package restaurantsQuartierRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	_ "github.com/lib/pq"
	"log"
)

func GetRestaurantsQuartier() (*[]*models.RestaurantsQuartier, error) {

	rows, err := config.DatabaseInit().Query("select id, libelle from quartier")
	if err != nil {
		log.Fatal(err)
	}

	var restaurantsQuartier []*models.RestaurantsQuartier

	for rows.Next() {
		quartier := models.RestaurantsQuartier{}
		if err := rows.Scan(&quartier.Id, &quartier.Libelle); err != nil {
			log.Println(err)
			return &restaurantsQuartier, err
		}
		restaurantsQuartier = append(restaurantsQuartier, &quartier)
	}
	err = rows.Close()
	if err != nil {
		log.Println(err)
		return &restaurantsQuartier, err
	}

	return &restaurantsQuartier, nil
}

func Create(quartier models.RestaurantsQuartier) error {

	_, err := config.DatabaseInit().Exec("insert into quartier(libelle) values ($1)",
		quartier.Libelle)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
