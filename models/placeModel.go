package models

import (
	"time"
)

type Place struct {
	Id              int64     `json:"id"`
	Note            float32   `json:"note"`
	Appreciation    string    `json:"appreciation"`
	Prixmoyen       float32   `json:"prixmoyen"`
	Adresse         string    `json:"adresse"`
	Ville           string    `json:"ville"`
	Datecreation    time.Time `json:"datecreation"`
	Nom             string    `json:"nom"`
	QuartierId      int       `json:"quartierid"`
	CategorieId     int       `json:"categorieid"`
	NoteCopiosite   int       `json:"notecopiosite"`
	NoteDeliciosite int       `json:"notedeliciosite"`
	NoteCadre       int       `json:"notecadre"`
	NoteAccueil     int       `json:"noteaccueil"`
}
type PlaceRequest struct {
	Appreciation    string  `json:"appreciation"`
	Prixmoyen       float32 `json:"prixmoyen"`
	Adresse         string  `json:"adresse"`
	Ville           string  `json:"ville"`
	Nom             string  `json:"nom"`
	QuartierId      int     `json:"quartierid"`
	CategorieId     int     `json:"categorieid"`
	NoteCopiosite   int     `json:"notecopiosite"`
	NoteDeliciosite int     `json:"notedeliciosite"`
	NoteCadre       int     `json:"notecadre"`
	NoteAccueil     int     `json:"noteaccueil"`
}
