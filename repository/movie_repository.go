package repository

import (
	"database/sql"
	"movie_store/models"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetAllMovies() ([]models.Movie, error) {
	query := `SELECT id, name, release_date FROM movies`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // used to release the db connection after function execution executed just before the function returns, regardless of whether it returns normally or due to an error.

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		var releaseDate sql.NullString
		err := rows.Scan(&movie.ID, &movie.Name, &releaseDate)
		if err != nil {
			return nil, err
		}
		if releaseDate.Valid {
			value := releaseDate.String
			movie.ReleaseDate = &value
		} else {
			movie.ReleaseDate = nil
		}
		movies = append(movies, movie)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepository) GetMovieByID(id int) (*models.Movie, error) {
	query := `SELECT id, name, release_date FROM movies WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var movie models.Movie
	var releaseDate sql.NullString
	err := row.Scan(&movie.ID, &movie.Name, &releaseDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No movie found with the given ID
		}
		return nil, err
	}
	if releaseDate.Valid {
		value := releaseDate.String
		movie.ReleaseDate = &value
	} else {
		movie.ReleaseDate = nil
	}
	return &movie, nil
}

func (r *MovieRepository) GetMovieDetails(id int) (models.Movie, []models.Actor, error) {
	// This can be better implemented using Joins but using go routines and channels to better understand thhem
	movieCh := make(chan models.Movie, 1)
	actorsCh := make(chan []models.Actor, 1)
	errCh := make(chan error, 2)

	go func() {
		movie, err := r.GetMovieByID(id)
		if err != nil {
			errCh <- err
			return
		}
		if movie == nil {
			errCh <- sql.ErrNoRows
			return
		}
		movieCh <- *movie
	}()

	go func() {
		actors, err := NewActorRepository(r.db).GetActorsByMovieID(id)
		if err != nil {
			errCh <- err
			return
		}
		actorsCh <- actors
	}()

	var movie models.Movie
	var actors []models.Actor

	for i := 0; i < 2; i++ {
		select {
		case value := <-movieCh:
			movie = value
		case value := <-actorsCh:
			actors = value
		case err := <-errCh:
			return models.Movie{}, nil, err
		}
	}

	return movie, actors, nil
}
