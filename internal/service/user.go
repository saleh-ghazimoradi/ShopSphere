package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/dto"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/helper"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/repository"
	"github.com/saleh-ghazimoradi/ShopSphere/internal/service/serviceModels"
	"github.com/saleh-ghazimoradi/ShopSphere/pkg/notification"
	"time"
)

type User interface {
	Signup(ctx context.Context, input dto.UserSignUp) (string, error)
	findUserByEmail(ctx context.Context, email string) (*serviceModels.User, error)
	Login(ctx context.Context, email, password string) (string, error)
	GetVerificationCode(ctx context.Context, user *serviceModels.User) error
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
	isVerifiedUser(ctx context.Context, id uint) bool
}

type UserService struct {
	userRepository repository.User
	authService    helper.Auth
	notifyClient   notification.NotifyClient
}

func (u *UserService) Signup(ctx context.Context, input dto.UserSignUp) (string, error) {

	hPassword, err := u.authService.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := u.userRepository.CreateUser(ctx, &serviceModels.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
	})
	if err != nil {
		return "", err
	}

	return u.authService.GenerateToken(user.ID, user.Email, user.UserType)
}

func (u *UserService) findUserByEmail(ctx context.Context, email string) (*serviceModels.User, error) {
	return u.userRepository.FindUser(ctx, email)
}

func (u *UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.findUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	err = u.authService.VerifyPassword(password, user.Password)
	if err != nil {
		return "", errors.New("wrong password")
	}

	return u.authService.GenerateToken(user.ID, user.Email, user.UserType)
}

func (u *UserService) isVerifiedUser(ctx context.Context, id uint) bool {
	currentUser, err := u.userRepository.FindUserById(ctx, id)
	return err == nil && currentUser.Verified
}

func (u *UserService) GetVerificationCode(ctx context.Context, user *serviceModels.User) error {

	if u.isVerifiedUser(ctx, user.ID) {
		return errors.New("user already verified")
	}

	code, err := u.authService.GenerateCode()
	if err != nil {
		return errors.New("could not generate code")
	}

	us := serviceModels.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = u.userRepository.UpdateUser(ctx, user.ID, &us)
	if err != nil {
		return errors.New("unable to update verification code")
	}

	user, err = u.userRepository.FindUserById(ctx, user.ID)
	if err != nil {
		return errors.New("user not found")
	}

	msg := fmt.Sprintf("Your verification code is %v", code)
	err = u.notifyClient.SendSMS(context.Background(), user.Phone, msg)
	if err != nil {
		return errors.New("error on sending sms")
	}

	return nil
}

func (u *UserService) VerifyCode(ctx context.Context, id uint, code int) error {

	if u.isVerifiedUser(ctx, id) {
		return errors.New("user already verified")
	}

	user, err := u.userRepository.FindUserById(ctx, id)
	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("invalid verification code")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	updateUser := &serviceModels.User{
		Verified: true,
	}

	_, err = u.userRepository.UpdateUser(ctx, id, updateUser)
	if err != nil {
		return errors.New("unable to update verification code")
	}

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

func NewUserService(userRepository repository.User, authService helper.Auth, notifyClient notification.NotifyClient) User {
	return &UserService{
		userRepository: userRepository,
		authService:    authService,
		notifyClient:   notifyClient,
	}
}
