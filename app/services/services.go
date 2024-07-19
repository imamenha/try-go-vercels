package services

import "try-go-vercels/app/repositories"

type Service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{repo: repo}
}

// Tambahkan metode service
