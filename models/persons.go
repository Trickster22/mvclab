package models

import "time"

type Post int

const (
	Actor Post = iota
	Director
)

type Person struct {
	Id        int       `db:"personid"`
	Name      string    `db:"personname"`
	Datebirth time.Time `db:"personbirth"`
	Country   string    `db:"personcountry"`
	Post      Post      `db:"personpost"`
}
