package service

import "arsip-sejarah-al/internal/repository"

type UserService interface {
	// Define methods for user service here
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Implement methods for user service
