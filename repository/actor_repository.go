package repository

import (
	"database/sql"
	"movie_store/models"
)

type ActorRepository struct {
	db *sql.DB
}

func NewActorRepository(db *sql.DB) *ActorRepository {
	return &ActorRepository{db: db}
}

func (a *ActorRepository) GetAllActors() ([]models.Actor, error) {
	query := `SELECT id, name, image FROM actors`
	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []models.Actor

	for rows.Next() {
		var actor models.Actor
		var image sql.NullString
		err := rows.Scan(&actor.ID, &actor.Name, &image)
		if err != nil {
			return nil, err
		}
		if image.Valid {
			value := image.String
			actor.Image = &value
		} else {
			actor.Image = nil
		}
		actors = append(actors, actor)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return actors, nil
}

func (a *ActorRepository) GetActorsByMovieID(movieID int) ([]models.Actor, error) {
	query := `
		SELECT a.id, a.name, a.image
		FROM actors a
		INNER JOIN movie_actors ma ON ma.actor_id = a.id
		WHERE ma.movie_id = ?
		ORDER BY a.id
	`

	rows, err := a.db.Query(query, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []models.Actor

	for rows.Next() {
		var actor models.Actor
		var image sql.NullString
		if err := rows.Scan(&actor.ID, &actor.Name, &image); err != nil {
			return nil, err
		}
		if image.Valid {
			value := image.String
			actor.Image = &value
		} else {
			actor.Image = nil
		}
		actors = append(actors, actor)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return actors, nil
}
