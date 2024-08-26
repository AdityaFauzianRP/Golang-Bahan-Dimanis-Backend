package repository

import (
	"arsip-sejarah-al/internal/model"
	"context"
	"database/sql"
)

type APIRouteRepository interface {
	GetAllRoutes(ctx context.Context) ([]model.APIRoute, error)
}

type apiRouteRepository struct {
	db *sql.DB
}

func NewAPIRouteRepository(db *sql.DB) APIRouteRepository {
	return &apiRouteRepository{db: db}
}

func (r *apiRouteRepository) GetAllRoutes(ctx context.Context) ([]model.APIRoute, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, path, method, function_name, middleware  FROM api_routes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []model.APIRoute
	for rows.Next() {
		var route model.APIRoute
		if err := rows.Scan(&route.ID, &route.Path, &route.Method, &route.FunctionName, &route.Middleware); err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	return routes, nil
}
