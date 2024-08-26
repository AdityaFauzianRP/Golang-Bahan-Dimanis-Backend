package repository

import (
	"database/sql"
)

type UserRepository interface {
	// Define methods for user repository here
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// Implement methods for user repository
