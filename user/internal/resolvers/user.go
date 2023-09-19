package resolvers

import (
	"context"
	"errors"
	"github.com/bludot/tempmee/user/graph/model"
	"github.com/bludot/tempmee/user/internal/services/jwt"
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

func SignIn(ctx context.Context, userService user.UserServiceImpl, jwtService jwt.JWTServiceImpl, email string, password string) (*model.SignInPayload, error) {
	foundUser, err := userService.SignIn(ctx, email, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := jwtService.GenerateToken(foundUser.ID, map[string]interface{}{
		"email": foundUser.Email,
	})

	return &model.SignInPayload{
		User: &model.User{
			ID:    foundUser.ID,
			Name:  foundUser.Name,
			Email: foundUser.Email,
		},
		Token: token,
	}, nil
}
