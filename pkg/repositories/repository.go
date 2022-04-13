package repositories

import "gorm.io/gorm"

type Repository struct {
}

func RepositoryInit(db *gorm.DB) *Repository {
	return &Repository{}
}
