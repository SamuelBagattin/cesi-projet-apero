package happyHourRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"log"
)

func Create(happy models.HappyHour) error {

	_, err := config.DatabaseInit().Exec("insert into apero(nom, dateapero, createur_id,  datecreation) values ($1, $2, $3, current_date)", happy.Nom, happy.DateApero, happy.Createur_Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAllWithCreator() *[]*models.HappyHourUser {

	rows, err := config.DatabaseInit().Query("select a.id, a.nom, a.dateApero, a.datecreation, a.createur_Id, COALESCE(u.nom,'') as nom, COALESCE(u.mail,'') as mail, COALESCE(u.numtel,'') as numTel from apero a left join utilisateur u on u.id = a.createur_id order by dateapero desc")
	if err != nil {
		log.Fatal(err)
	}

	var happyHourUsers []*models.HappyHourUser

	for rows.Next() {
		happyHourUser := models.HappyHourUser{}
		if err := rows.Scan(&happyHourUser.Id, &happyHourUser.Nom, &happyHourUser.DateApero, &happyHourUser.DateCreation, &happyHourUser.User.Id, &happyHourUser.User.Nom, &happyHourUser.User.Mail, &happyHourUser.User.NumTel); err != nil {
			log.Fatal(err)
		}
		happyHourUsers = append(happyHourUsers, &happyHourUser)
	}
	err = rows.Close()
	if err != nil {
		panic(err)
	}

	return &happyHourUsers
}

func GetAll() *[]*models.HappyHour {

	rows, err := config.DatabaseInit().Query("select a.id, a.nom, a.dateApero, a.datecreation, a.createur_Id from apero a order by dateapero desc")
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
