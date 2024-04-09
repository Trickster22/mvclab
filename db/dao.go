package db

import (
	"lab1/models"

	"github.com/jmoiron/sqlx"
)

func GetStudio(db *sqlx.DB, studioId string) models.Studio {
	studio := models.Studio{}
	err := db.Get(&studio, "SELECT * FROM studios WHERE studioid = $1", studioId)
	CheckError(err)
	return studio
}

func GetAllMovies(db *sqlx.DB) []models.Movie {
	movies := []models.DbMovie{}
	err := db.Select(&movies, "SELECT * FROM movie")
	CheckError(err)
	result := make([]models.Movie, 0)
	for _, m := range movies {
		var movie models.Movie
		movie.Id = m.Id
		movie.Name = m.Name
		movie.Year = m.Year
		movie.Genre = getGenre(db, m.GenreId)
		movie.Studio = GetStudio(db, m.StudioId)
		result = append(result, movie)
	}
	return result
}

func getGenre(db *sqlx.DB, genreId string) models.Genre {
	genre := models.Genre{}
	err := db.Get(&genre, "SELECT * FROM genre WHERE genreid = $1", genreId)
	CheckError(err)
	return genre
}

func GetAllDirectors(db *sqlx.DB) []models.Person {
	return GetAllPersons(db, 1)
}

func GetAllActors(db *sqlx.DB) []models.Person {
	return GetAllPersons(db, 0)
}

func GetAllPersons(db *sqlx.DB, post int) []models.Person {
	person := []models.Person{}
	err := db.Select(&person, "SELECT * FROM person WHERE personpost = $1", post)
	CheckError(err)
	return person
}

func GetMovie(db *sqlx.DB, movieId string) models.Movie {
	dbMovie := models.DbMovie{}
	err := db.Get(&dbMovie, "SELECT * FROM movie WHERE movieid = $1", movieId)
	CheckError(err)
	var movie models.Movie
	movie.Id = dbMovie.Id
	movie.Name = dbMovie.Name
	movie.Year = dbMovie.Year
	movie.Genre = getGenre(db, dbMovie.GenreId)
	movie.Studio = GetStudio(db, dbMovie.StudioId)
	return movie
}

func GetDirectorName(db *sqlx.DB, movieId string) string {
	var directorName string
	err := db.Get(&directorName, "SELECT personname FROM person AS P INNER JOIN movietoperson AS MP ON MP.personid = P.personid WHERE MP.movieid = $1 AND P.personpost = 1", movieId)
	if err != nil {
		return "Неизвестно"
	}
	return directorName
}

func GetActors(db *sqlx.DB, movieId string) []string {
	actors := []string{}
	err := db.Select(&actors, "SELECT personname FROM person AS P INNER JOIN movietoperson AS MP ON MP.personid = P.personid WHERE MP.movieid = $1 AND P.personpost = 0", movieId)
	CheckError(err)
	if err != nil {
		return make([]string, 0)
	}
	return actors

}
