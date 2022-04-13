package services

import "microserviceSample/pkg/repositories"

type Service struct {
}

func ServiceInit(repository *repositories.Repository) *Service {
	return &Service{}
}
