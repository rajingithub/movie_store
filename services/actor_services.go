package services

import (
	"movie_store/models"
	"movie_store/repository"
)

type ActorService struct {
	repository *repository.ActorRepository
}

func NewActorService(repository *repository.ActorRepository) *ActorService {
	return &ActorService{repository: repository}
}

func (s *ActorService) GetAllActors() ([]models.Actor, error) {
	return s.repository.GetAllActors()
}
