package placesCategoryRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	_ "github.com/lib/pq"
	"log"
)

func GetPlacesCategories() (*[]*models.PlaceCategory, error) {

	rows, err := config.DatabaseInit().Query("select id, libelle from categorie order by libelle asc")
	if err != nil {
		log.Fatal(err)
	}

	var placeCategories []*models.PlaceCategory

	for rows.Next() {
		category := models.PlaceCategory{}
		if err := rows.Scan(&category.Id, &category.Libelle); err != nil {
			log.Println(err)
			return &placeCategories, err
		}
		placeCategories = append(placeCategories, &category)
	}
	err = rows.Close()
	if err != nil {
		log.Println(err)
		return &placeCategories, err
	}

	return &placeCategories, nil
}

func Create(category models.PlaceCategory) error {

	_, err := config.DatabaseInit().Exec("insert into categorie(libelle) values ($1)",
		category.Libelle)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
