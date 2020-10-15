package voteRepository

import (
	"github.com/SamuelBagattin/cesi-projet-apero/config"
	"github.com/SamuelBagattin/cesi-projet-apero/custom_errors"
	"github.com/SamuelBagattin/cesi-projet-apero/models"
)

func GetAll() (*[]*models.Vote, error) {

	rows, err := config.DatabaseInit().Query("select id, nbVotes, datevote, endroit_id, utilisateur_id, apero_id  from vote ")
	if err != nil {
		return nil, err
	}

	var votes []*models.Vote

	for rows.Next() {
		vote := models.Vote{}
		if err := rows.Scan(&vote.Id, &vote.NbVotes, &vote.Date, &vote.PlaceId, &vote.UserId, &vote.HappyhourId); err != nil {
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

	_, err := config.DatabaseInit().Exec("insert into vote (utilisateur_id, apero_id, endroit_id) values ($1,$2,$3, vote.UserId, vote.HappyHourId, vote.PlaceId")
	if err != nil {
		return err
	}
	return nil
}

func GetOne(id int) (*models.Vote, error) {
	row := config.DatabaseInit().QueryRow("select id, nbvotes, utilisateur_id, apero_id, endroit_id  from vote where id = $1", id)

	vote := models.Vote{}

	if err := row.Scan(&vote.Id, &vote.NbVotes, &vote.UserId, &vote.HappyhourId, &vote.PlaceId); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, custom_errors.EntityNotFound{
				Id: id,
			}
		}
		return nil, err
	}

	return &vote, nil
}
