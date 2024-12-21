package service

import (
	"context"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/dto"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/repository"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service/serviceModels"
	"log"
)

type User interface {
	Signup(ctx context.Context, input dto.UserSignUp) (string, error)
	findUserByEmail(ctx context.Context, email string) (*serviceModels.User, error)
	Login(ctx context.Context, input any) (string, error)
	GetVerificationCode(ctx context.Context, user serviceModels.User) (int, error)
	VerifyCode(ctx context.Context, id uint, code int) error
	CreateProfile(ctx context.Context, id uint, input any) error
	GetProfile(ctx context.Context, id uint) (*serviceModels.User, error)
	UpdateProfile(ctx context.Context, id uint, input any) error
	BecomeSeller(ctx context.Context, id uint, input any) (string, error)
	FindCart(ctx context.Context, id uint) ([]any, error)
	CreateCart(ctx context.Context, input any, user serviceModels.User) ([]any, error)
	CreateOrder(ctx context.Context, user serviceModels.User) (int, error)
	GetOrders(ctx context.Context, user serviceModels.User) ([]any, error)
	GetOrderById(ctx context.Context, id uint, uId uint) (any, error)
}

type UserService struct {
	userRepository repository.User
}

func (u *UserService) Signup(ctx context.Context, input dto.UserSignUp) (string, error) {
	log.Println(input)

	return "this is my token", nil
}
func (u *UserService) findUserByEmail(ctx context.Context, email string) (*serviceModels.User, error) {
	return nil, nil
}
func (u *UserService) Login(ctx context.Context, input any) (string, error) {
	return "", nil
}
func (u *UserService) GetVerificationCode(ctx context.Context, user serviceModels.User) (int, error) {
	return 0, nil
}
func (u *UserService) VerifyCode(ctx context.Context, id uint, code int) error {
	return nil
}
func (u *UserService) CreateProfile(ctx context.Context, id uint, input any) error {
	return nil
}
func (u *UserService) GetProfile(ctx context.Context, id uint) (*serviceModels.User, error) {
	return nil, nil
}
func (u *UserService) UpdateProfile(ctx context.Context, id uint, input any) error {
	return nil
}
func (u *UserService) BecomeSeller(ctx context.Context, id uint, input any) (string, error) {
	return "", nil
}
func (u *UserService) FindCart(ctx context.Context, id uint) ([]any, error) {
	return nil, nil
}
func (u *UserService) CreateCart(ctx context.Context, input any, user serviceModels.User) ([]any, error) {
	return nil, nil
}
func (u *UserService) CreateOrder(ctx context.Context, user serviceModels.User) (int, error) {
	return 0, nil
}
func (u *UserService) GetOrders(ctx context.Context, user serviceModels.User) ([]any, error) {
	return nil, nil
}
func (u *UserService) GetOrderById(ctx context.Context, id uint, uId uint) (any, error) {
	return nil, nil
}

func NewUserService(userRepository repository.User) User {
	return &UserService{
		userRepository: userRepository,
	}
}
