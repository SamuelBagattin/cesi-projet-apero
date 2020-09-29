package happyHourRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"log"
)

func Create(happy models.HappyHour) error {

	_, err := config.DatabaseInit().Exec("insert into apero(nom, dateapero, createur_id,  datecreation) values ($1,$2,$3, current_date)", happy.Nom, happy.DateApero, happy.Createur_Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAll() *[]*models.HappyHour {

	rows, err := config.DatabaseInit().Query("select id, nom, dateApero, datecreation, createur_Id from apero")
	if err != nil {
		log.Fatal(err)
	}

	var happyHours []*models.HappyHour

	for rows.Next() {
		happyHour := models.HappyHour{}
		if err := rows.Scan(&happyHour.Id, &happyHour.Nom, &happyHour.DateApero, &happyHour.DateCreation, &happyHour.Createur_Id); err != nil {
			log.Fatal(err)
		}
		happyHours = append(happyHours, &happyHour)
	}
	err = rows.Close()
	if err != nil {
		panic(err)
	}

	return &happyHours
}
