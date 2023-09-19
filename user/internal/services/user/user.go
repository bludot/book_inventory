package user

import (
	"context"
	"errors"
	"github.com/bludot/tempmee/user/internal/db/repositories/user"
)

type UserServiceImpl interface {
	Register(ctx context.Context, name string, email string, password string) (*user.User, error)
	SignIn(ctx context.Context, email string, password string) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	GetAllUsers(ctx context.Context, limit int) ([]*user.User, error)
}

type UserService struct {
	Repository user.UserRepositoryImpl
}

func NewUserService(userRepository user.UserRepositoryImpl) UserServiceImpl {
	return &UserService{
		Repository: userRepository,
	}
}

func (a *UserService) Register(ctx context.Context, name string, email string, password string) (*user.User, error) {
	return a.Repository.CreateUser(name, email, password)
}

func (a *UserService) SignIn(ctx context.Context, email string, password string) (*user.User, error) {
	foundUser, err := a.Repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if foundUser == nil {
		return nil, errors.New("user not found")
	}
	if foundUser.Password != password {
		return nil, errors.New("password is incorrect")
	}
	return foundUser, nil

}

func (a *UserService) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	return a.Repository.FindByEmail(email)
}

func (a *UserService) GetAllUsers(ctx context.Context, limit int) ([]*user.User, error) {
	return a.Repository.FindAll()
}
