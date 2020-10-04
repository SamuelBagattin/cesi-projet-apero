package models

import "encoding/binary"

type User struct {
	Id     int64            `json:"id"`
	Nom    string           `json:"nom"`
	NumTel string           `json:"NumTel"`
	Photo  binary.ByteOrder `json:"Photo"`
	Mail   string           `json:"mail"`
}
