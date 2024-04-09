package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connection() *sqlx.DB {

	db, err := sqlx.Connect("postgres", "user=postgres password = admin dbname=myDb port=5433 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	println("Conncted")
	return db

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
