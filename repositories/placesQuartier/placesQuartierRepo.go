package restaurantsQuartierRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	_ "github.com/lib/pq"
	"log"
)

func GetPlacesDistrict() (*[]*models.PlacesDistrict, error) {

	rows, err := config.DatabaseInit().Query("select id, libelle from quartier")
	if err != nil {
		log.Fatal(err)
	}

	var placesDistricts []*models.PlacesDistrict

	for rows.Next() {
		quartier := models.PlacesDistrict{}
		if err := rows.Scan(&quartier.Id, &quartier.Libelle); err != nil {
			log.Println(err)
			return &placesDistricts, err
		}
		placesDistricts = append(placesDistricts, &quartier)
	}
	err = rows.Close()
	if err != nil {
		log.Println(err)
		return &placesDistricts, err
	}

	return &placesDistricts, nil
}

func Create(quartier models.PlacesDistrict) error {

	_, err := config.DatabaseInit().Exec("insert into quartier(libelle) values ($1)",
		quartier.Libelle)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
