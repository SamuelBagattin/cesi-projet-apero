package models

import "time"

type Vote struct {
	Id          int64     `json:"id"`
	Date        time.Time `json:"date"`
	PlaceId     int8      `json:"placeId"`
	UserId      int8      `json:"userId"`
	HappyhourId int8      `json:"happyHourId"`
	Place       `json:"place"`
	User        `json:"user"`
	HappyHour   `json:"happyHour"`
}

type UpdateVoteRequest struct {
	Id          int64 `json:"id"`
	PlaceId     int8  `json:"placeId"`
	UserId      int8  `json:"userId"`
	HappyhourId int8  `json:"happyHourId"`
}

type AddVoteRequest struct {
	PlaceId     int8 `json:"placeId"`
	UserId      int8 `json:"userId"`
	HappyhourId int8 `json:"happyHourId"`
}
