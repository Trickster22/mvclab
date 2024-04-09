package models

type Rating struct {
	MovieId      int     `json:"movieid"`
	PersonRating float32 `json:"personrating"`
	CriticRating float32 `json:"criticrating"`
}
