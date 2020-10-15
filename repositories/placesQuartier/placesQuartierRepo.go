package restaurantsQuartierRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	_ "github.com/lib/pq"
	"log"
)

func GetPlacesDistrict() (*[]*models.PlacesDistrict, error) {

	rows, err := config.DatabaseInit().Query("select id, libelle from quartier order by libelle asc")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var placesDistricts []*models.PlacesDistrict

	for rows.Next() {
		quartier := models.PlacesDistrict{}
		if err := rows.Scan(&quartier.Id, &quartier.Libelle); err != nil {
			return nil, err
		}
		placesDistricts = append(placesDistricts, &quartier)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return &placesDistricts, nil
}

func Create(quartier models.PlacesDistrict) error {

	_, err := config.DatabaseInit().Exec("insert into quartier(libelle) values ($1)",
		quartier.Libelle)

	if err != nil {
		return err
	}

	return nil

}
