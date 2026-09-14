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
