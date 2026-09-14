package main

import (
	"fmt"
	"net/http"

	"movie_store/handlers"
	"movie_store/repository"
	"movie_store/services"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func main() {
	db := repository.GetDB()
	movieRepository := repository.NewMovieRepository(db)
	movieService := services.NewMovieService(movieRepository)
	movieHandler := handlers.NewMovieHandler(movieService)
	actorRepository := repository.NewActorRepository(db)
	actorService := services.NewActorService(actorRepository)
	actorHandler := handlers.NewActorHandler(actorService)
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/movies", movieHandler.GetAllMovies).Methods("GET")
	router.HandleFunc("/actors", actorHandler.GetAllActors).Methods("GET")
	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
