package models

type Genre struct {
	Id   int    `db:"genreid"`
	Name string `db:"genrename"`
}
