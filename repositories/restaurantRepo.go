package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)

func GetRestaurants() []restaurant {
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

	var restaurants []restaurant

	for rows.Next() {
		restau := restaurant{}
		if err := rows.Scan(&restau.Id, &restau.Note, &restau.Appreciation, &restau.Prixmoyen, &restau.Adresse, &restau.Ville,
			&restau.Datecreation, &restau.Nom, &restau.NoteId, &restau.QuartierId, &restau.CategorieId); err != nil {
			log.Fatal(err)
		}
		result, _ := json.Marshal(restau)
		log.Printf(string(result))
		restaurants = append(restaurants, restau)
	}
	err = rows.Close()
	if err != nil {
		panic(err)
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
	return restaurants
}

type restaurant struct {
	Id           int64 `json:"id"`
	Note         int8  `json:"note"`
	Appreciation string
	Prixmoyen    float32
	Adresse      string
	Ville        string
	Nom          string
	NoteId       int
	QuartierId   int
	CategorieId  int
	Datecreation time.Time
}
