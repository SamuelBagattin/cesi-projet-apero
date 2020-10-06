package models

import "time"

type Vote struct {
	Id          int64     `json:"id"`
	NbVotes     int8      `json:"nbVotes"`
	Date        time.Time `json:"NumTel"`
	PlaceId     int8      `json:"Place_Id"`
	UserId      int8      `json:"User_Id"`
	HappyhourId int8      `json:"HappyHour_Id"`
}
