package repository

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service/serviceModels"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User interface {
	CreateUser(ctx context.Context, user *serviceModels.User) (*serviceModels.User, error)
	FindUser(ctx context.Context, email string) (*serviceModels.User, error)
	FindUserById(ctx context.Context, id uint) (*serviceModels.User, error)
	UpdateUser(ctx context.Context, id uint, user *serviceModels.User) (*serviceModels.User, error)

	CreateBankAccount(ctx context.Context, bankAccount *serviceModels.BankAccount) error
}

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) CreateUser(ctx context.Context, user *serviceModels.User) (*serviceModels.User, error) {
	if err := u.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, errors.New("error while creating user: " + err.Error())
	}
	return user, nil
}

func (u *UserRepository) FindUser(ctx context.Context, email string) (*serviceModels.User, error) {
	var user serviceModels.User

	err := u.db.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("error while finding user: " + err.Error())
	}

	return &user, nil
}

func (u *UserRepository) FindUserById(ctx context.Context, id uint) (*serviceModels.User, error) {
	var user serviceModels.User

	err := u.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("error while finding user: " + err.Error())
	}

	return &user, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, id uint, user *serviceModels.User) (*serviceModels.User, error) {
	var us serviceModels.User
	if err := u.db.WithContext(ctx).Model(&us).Clauses(clause.Returning{}).Where("id=?", id).Updates(user).Error; err != nil {
		return nil, errors.New("error while updating user: " + err.Error())
	}
	return &us, nil
}

func (u *UserRepository) CreateBankAccount(ctx context.Context, bankAccount *serviceModels.BankAccount) error {
	return u.db.WithContext(ctx).Create(&bankAccount).Error
}

func NewUserRepository(db *gorm.DB) User {
	return &UserRepository{
		db: db,
	}
}
