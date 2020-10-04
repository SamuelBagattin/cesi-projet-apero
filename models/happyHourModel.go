package models

import "time"

type HappyHour struct {
	Id           int64     `json:"id"`
	Nom          string    `json:"nom"`
	DateApero    string    `json:"dateApero"`
	DateCreation time.Time `json:"dateCreation"`
	Createur_Id  int32     `json:"createur_Id"`
	User         User      `json:"user"`
}
