package main

import (
	"lab1/controllers"
	"log"
	"net/http"
)

func main() {
	handelRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handelRoutes() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/movies", controllers.Movies)
	http.HandleFunc("/directors", controllers.Directors)
	http.HandleFunc("/actors", controllers.Actors)
	http.HandleFunc("/movie", controllers.Movie)
}
