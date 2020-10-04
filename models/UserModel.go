package models

type User struct {
	Id     int64  `json:"id"`
	Nom    string `json:"nom"`
	NumTel string `json:"NumTel"`
	//Photo    `json:"Photo"`
	Mail string `json:"mail"`
}
