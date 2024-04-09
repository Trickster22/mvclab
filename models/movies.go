package models

type Movie struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  Genre  `json:"genre"`
	Studio Studio `json:"studio"`
}

type DbMovie struct {
	Id       int    `db:"movieid"`
	Name     string `db:"moviename"`
	Year     int    `db:"movieyear"`
	GenreId  string `db:"moviegenre"`
	StudioId string `db:"studioid"`
}
