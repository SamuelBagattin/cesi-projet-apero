package voteRepository

import (
	"fmt"
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/custom_errors"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
	"strings"
)

func GetAll() (*[]*models.Vote, error) {

	rows, err := config.DatabaseInit().Query("select id, datevote, endroit_id, utilisateur_id, apero_id  from vote ")
	if err != nil {
		return nil, err
	}

	var votes []*models.Vote

	for rows.Next() {
		vote := models.Vote{}
		if err := rows.Scan(&vote.Id, &vote.Date, &vote.PlaceId, &vote.UserId, &vote.HappyhourId); err != nil {
			return nil, err
		}
		votes = append(votes, &vote)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return &votes, nil
}

func Create(vote models.Vote) error {

	_, err := config.DatabaseInit().Exec("insert into vote (utilisateur_id, apero_id, endroit_id, datevote) values ($1,$2,$3,current_date)", vote.UserId, vote.HappyhourId, vote.PlaceId)
	if err != nil {
		return err
	}
	return nil
}
func Update(vote models.Vote) error {

	_, err := config.DatabaseInit().Exec("UPDATE vote SET utilisateur_id = $1, apero_id = $2, endroit_id = $3, datevote = current_date WHERE id = $4", vote.UserId, vote.HappyhourId, vote.PlaceId, vote.Id)
	if err != nil {
		return err
	}
	return nil
}
func CreateAll(votes []models.Vote) error {

	var query strings.Builder
	query.WriteString("insert into vote (utilisateur_id, apero_id, endroit_id, datevote) values")
	for _, vote := range votes {
		query.WriteString(fmt.Sprintf(" (%d, %d, %d, current_date) ", vote.UserId, vote.HappyhourId, vote.PlaceId))
	}

	_, err := config.DatabaseInit().Exec(query.String())
	if err != nil {
		return err
	}
	return nil
}

func UpdateAll(votes []models.Vote) error {

	var query strings.Builder
	query.WriteString("update vote as u set endroit_id = u2.endroit_id, apero_id = u2.apero_id, utilisateur_id = u2.utilisateur_id, datevote = current_date from (values ")
	for i, vote := range votes {
		query.WriteString(fmt.Sprintf(" (%d, %d, %d, %d) ", vote.Id, vote.PlaceId, vote.HappyhourId, vote.UserId))
		if i != len(votes)-1 {
			query.WriteString(",")
		}
	}
	query.WriteString(") as u2(id, endroit_id, apero_id, utilisateur_id) where u2.id = u.id;")

	_, err := config.DatabaseInit().Exec(query.String())
	if err != nil {
		return err
	}
	return nil
}

func GetOne(id int) (*models.Vote, error) {
	row := config.DatabaseInit().QueryRow("select id, utilisateur_id, apero_id, endroit_id  from vote where id = $1", id)

	vote := models.Vote{}

	if err := row.Scan(&vote.Id, &vote.UserId, &vote.HappyhourId, &vote.PlaceId); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, custom_errors.EntityNotFound{
				Id: id,
			}
		}
		return nil, err
	}

	return &vote, nil
}
