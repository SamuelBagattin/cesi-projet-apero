package repositories

import (
	"database/sql"
	"fmt"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)

func GetRestaurants() *[]*models.Restaurant {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select * from restaurant")
	if err != nil {
		log.Fatal(err)
	}

	var restaurants []*models.Restaurant

	for rows.Next() {
		restau := models.Restaurant{}
		if err := rows.Scan(&restau.Id, &restau.Note, &restau.Appreciation, &restau.Prixmoyen, &restau.Adresse, &restau.Ville,
			&restau.Datecreation, &restau.Nom, &restau.NoteId, &restau.QuartierId, &restau.CategorieId); err != nil {
			log.Fatal(err)
		}
		restaurants = append(restaurants, &restau)
	}
	err = rows.Close()
	if err != nil {
		panic(err)
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
	return &restaurants
}
