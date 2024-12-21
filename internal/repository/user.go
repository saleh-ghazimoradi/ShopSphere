package repository

import "gorm.io/gorm"

type User interface {
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &UserRepository{
		db: db,
	}
}
