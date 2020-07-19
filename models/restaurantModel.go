package models

import "time"

type Restaurant struct {
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
