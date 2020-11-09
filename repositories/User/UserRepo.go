package UserRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
)

func GetAll() (*[]*models.User, error) {

	rows, err := config.DatabaseInit().Query("select id, nom, coalesce(utilisateur.mail,'')as mail, coalesce(utilisateur.numtel,'')as numtel from utilisateur order by nom")
	if err != nil {
		return nil, err
	}

	var users []*models.User

	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Nom, &user.Mail, &user.NumTel); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func Create(user models.User) error {

	_, err := config.DatabaseInit().Exec("insert into utilisateur(nom, mail, numTel) values ($1,$2,$3)", user.Nom, user.Mail, user.NumTel)
	if err != nil {
		return err
	}
	return nil
}

func GetOneUser(id int) (*models.User, error) {
	row := config.DatabaseInit().QueryRow("select id, nom, coalesce(mail,''), coalesce(numtel, '') from utilisateur where id = $1", id)

	user := models.User{}

	if err := row.Scan(&user.Id, &user.Nom, &user.Mail, &user.NumTel); err != nil {
		return nil, err
	}

	return &user, nil
}
