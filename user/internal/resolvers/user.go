package resolvers

import (
	"context"
	"github.com/bludot/tempmee/user/graph/model"
	"github.com/bludot/tempmee/user/internal/services/user"
)

func Register(ctx context.Context, userService user.UserServiceImpl, name string, email string, password string) (*model.User, error) {
	registeredUser, err := userService.Register(ctx, name, email, password)
	if err != nil {
		return &model.User{}, err

	}

	return &model.User{
		ID:    registeredUser.ID,
		Name:  registeredUser.Name,
		Email: registeredUser.Email,
	}, nil
}
