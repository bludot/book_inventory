package resolvers

import (
	"context"
	"errors"
	"github.com/bludot/tempmee/user/graph/model"
	"github.com/bludot/tempmee/user/internal/services/user"
)

func Register(ctx context.Context, userService user.UserServiceImpl, name string, email string, password string) (*model.User, error) {
	registeredUser, err := userService.Register(ctx, name, email, password)
	if err != nil {
		return nil, err

	}

	return &model.User{
		ID:    registeredUser.ID,
		Name:  registeredUser.Name,
		Email: registeredUser.Email,
	}, nil
}

func SignIn(ctx context.Context, userService user.UserServiceImpl, email string, password string) (*model.User, error) {
	foundUser, err := userService.SignIn(ctx, email, password)
	if err != nil {
		return &model.User{}, errors.New("invalid credentials")
	}

	return &model.User{
		ID:    foundUser.ID,
		Name:  foundUser.Name,
		Email: foundUser.Email,
	}, nil
}
