package history

import (
	"backend/internal/domain"
	"context"
)

type Service interface {
	// GetAll obtain all histories.
	GetAll(ctx context.Context) ([]domain.History, error)

	// Save a new History.
	Save(ctx context.Context, history domain.History) (domain.History, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) GetAll(ctx context.Context) ([]domain.History, error) {
	return s.repository.GetAll(ctx)
}

func (s service) Save(ctx context.Context, history domain.History) (domain.History, error) {
	id, err := s.repository.Save(ctx, history)
	if err != nil {
		return domain.History{}, err
	}

	history.ID = id

	return history, nil
}
