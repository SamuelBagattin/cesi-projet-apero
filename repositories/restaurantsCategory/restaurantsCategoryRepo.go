package restaurantsCategoryRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	_ "github.com/lib/pq"
	"log"
)

func GetRestaurantsCategories() (*[]*models.RestaurantCategory, error) {

	rows, err := config.DatabaseInit().Query("select id, libelle from categorie")
	if err != nil {
		log.Fatal(err)
	}

	var restaurantCategories []*models.RestaurantCategory

	for rows.Next() {
		category := models.RestaurantCategory{}
		if err := rows.Scan(&category.Id, &category.Libelle); err != nil {
			log.Println(err)
			return &restaurantCategories, err
		}
		restaurantCategories = append(restaurantCategories, &category)
	}
	err = rows.Close()
	if err != nil {
		log.Println(err)
		return &restaurantCategories, err
	}

	return &restaurantCategories, nil
}

func Create(rest models.RestaurantCategory) error {

	_, err := config.DatabaseInit().Exec("insert into categorie(libelle) values ($1)",
		rest.Libelle)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
