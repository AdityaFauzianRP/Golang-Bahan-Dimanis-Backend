package main

import (
	"arsip-sejarah-al/config"
	"arsip-sejarah-al/internal/handler"
	"arsip-sejarah-al/internal/middleware"
	"arsip-sejarah-al/internal/model"
	"arsip-sejarah-al/internal/repository"
	"arsip-sejarah-al/internal/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config-production.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbpool, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer dbpool.Close()

	apiRouteRepo := repository.NewAPIRouteRepository(dbpool)

	apiRouteService := service.NewAPIRouteService(apiRouteRepo)

	userHandler := handler.NewUserHandler()

	r := gin.Default()

	routes, err := apiRouteService.GetAllRoutes(context.Background())
	if err != nil {
		log.Fatalf("Failed to load routes from database: %v", err)
	}

	validationRepo := repository.NewValidationSpecRepository(dbpool)

	for _, route := range routes {
		if err := validateRoute(route); err != nil {
			log.Printf("Skipping invalid route: %v", err)
			continue
		}

		if route.Middleware {
			specs, err := validationRepo.GetValidationSpecs(context.Background(), route.Path)
			if err != nil {
				log.Fatalf("Failed to load validation specs: %v", err)
			}

			secure := r.Group("").Use(middleware.JWTAuthMiddleware(), middleware.ValidateUser(dbpool))
			secure.Use(middleware.ValidateDataMiddleware(specs))
			path := adjustPathForSecureGroup(route.Path)
			secure.Handle(route.Method, path, userHandler.GetHandler(route.FunctionName))
		} else {
			specs, err := validationRepo.GetValidationSpecs(context.Background(), route.Path)
			if err != nil {
				log.Fatalf("Failed to load validation specs: %v", err)
			}
			standar := r.Group("").Use(middleware.ValidateDataMiddleware(specs))
			standar.Handle(route.Method, route.Path, userHandler.GetHandler(route.FunctionName))
		}
	}

	log.Fatal(r.Run(":8082"))
}

func adjustPathForSecureGroup(path string) string {
	return path
}

func validateRoute(route model.APIRoute) error {
	if route.Path == "" || route.Method == "" || route.FunctionName == "" {
		return fmt.Errorf("invalid route data: %v", route)
	}

	validMethods := map[string]bool{"GET": true, "POST": true, "PUT": true, "DELETE": true}
	if !validMethods[route.Method] {
		return fmt.Errorf("invalid HTTP method: %s", route.Method)
	}

	return nil
}
