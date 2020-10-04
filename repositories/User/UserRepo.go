package UserRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"log"
	"strconv"
)

func GetAll() (*[]*models.User, error) {

	rows, err := config.DatabaseInit().Query("select id, nom, coalesce(utilisateur.mail,'')as mail, coalesce(utilisateur.numtel,'')as numtel from utilisateur order by nom asc")
	if err != nil {
		log.Fatal(err)
	}

	var users []*models.User

	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Nom, &user.Mail, &user.NumTel); err != nil {
			log.Println(err)
			return &users, err
		}
		users = append(users, &user)
	}
	err = rows.Close()
	if err != nil {
		log.Println(err)
		return &users, err
	}

	return &users, nil
}

func Create(user models.User) error {

	_, err := config.DatabaseInit().Exec("insert into utilisateur(nom, mail, numTel) values ($1,$2,$3)", user.Nom, user.Mail, user.NumTel)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetOneUser(id string) models.User {

	intId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
	}

	row := config.DatabaseInit().QueryRow("select * from utilisateur where id = $1", intId)

	user := models.User{}

	if err := row.Scan(&user.Id, &user.Nom, &user.Mail, &user.NumTel); err != nil {
		log.Fatal(err)
	}

	return user
}
