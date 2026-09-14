package services

import (
	"movie_store/models"
	"movie_store/repository"
)

type MovieService struct {
	repository *repository.MovieRepository
}

func NewMovieService(repository *repository.MovieRepository) *MovieService {
	return &MovieService{repository: repository}
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	return s.repository.GetAllMovies()
}
