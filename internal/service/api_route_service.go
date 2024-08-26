package service

import (
	"arsip-sejarah-al/internal/model"
	"arsip-sejarah-al/internal/repository"
	"context"
)

type APIRouteService interface {
	GetAllRoutes(ctx context.Context) ([]model.APIRoute, error)
}

type apiRouteService struct {
	repo repository.APIRouteRepository
}

func NewAPIRouteService(repo repository.APIRouteRepository) APIRouteService {
	return &apiRouteService{repo: repo}
}

func (s *apiRouteService) GetAllRoutes(ctx context.Context) ([]model.APIRoute, error) {
	return s.repo.GetAllRoutes(ctx)
}
