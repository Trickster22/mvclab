package models

type Studio struct {
	Id      int    `db:"studioid"`
	Name    string `db:"studioname"`
	Year    int    `db:"studioyear"`
	Country string `db:"studiocountry"`
}
