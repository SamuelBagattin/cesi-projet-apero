package models

import "time"

type HappyHour struct {
	Id           int64     `json:"id"`
	Nom          string    `json:"nom"`
	DateApero    time.Time `json:"dateApero"`
	DateCreation time.Time `json:"dateCreation"`
	CreateurId   int32     `json:"createur_Id"`
}
