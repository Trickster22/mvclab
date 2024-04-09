package controllers

import (
	"html/template"
	"lab1/db"
	"lab1/models"
	"net/http"
)

type AllMovieInfo struct {
	Movie    models.Movie
	Director string
	Actors   []string
}

var connection = db.Connection()

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/home.html")
	if err != nil {
		print(err)
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		print(err)
	}
}

func Movies(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/movies.html")
	if err != nil {
		print(err)
	}
	movies := db.GetAllMovies(connection)
	err = tmpl.Execute(w, movies)

	if err != nil {
		print(err)
	}
}

func Movie(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/movie.html")
	query := r.URL.Query()
	movieId := query.Get("id")
	if err != nil {
		print(err)
	}
	var allMovieInfo AllMovieInfo
	allMovieInfo.Movie = db.GetMovie(connection, movieId)
	allMovieInfo.Director = db.GetDirectorName(connection, movieId)
	allMovieInfo.Actors = db.GetActors(connection, movieId)
	err = tmpl.Execute(w, allMovieInfo)

	if err != nil {
		print(err)
	}
}

func Directors(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/directors.html")
	if err != nil {
		print(err)
	}
	directors := db.GetAllDirectors(connection)
	err = tmpl.Execute(w, directors)

	if err != nil {
		print(err)
	}
}

func Actors(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/actors.html")
	if err != nil {
		print(err)
	}
	actors := db.GetAllActors(connection)
	err = tmpl.Execute(w, actors)

	if err != nil {
		print(err)
	}
}
