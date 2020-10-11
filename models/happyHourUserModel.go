package models

import "time"

type HappyHourUser struct {
	Id           int64     `json:"id"`
	Nom          string    `json:"nom"`
	DateApero    time.Time `json:"dateApero"`
	DateCreation time.Time `json:"dateCreation"`
	User         User      `json:"user"`
}
