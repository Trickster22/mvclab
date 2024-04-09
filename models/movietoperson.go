package models

type MovieToPerson struct {
	Id       int `json:"id"`
	MovieId  int `json:"movieid"`
	PersonId int `json:"personid"`
}
