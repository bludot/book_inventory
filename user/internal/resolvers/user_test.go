package resolvers_test

import (
	"github.com/bludot/tempmee/user/internal/db/repositories/user"
	"github.com/bludot/tempmee/user/internal/resolvers"
	"github.com/bludot/tempmee/user/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRegister(t *testing.T) {
	t.Run("should register user", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		u := mocks.NewMockUserServiceImpl(ctrl)
		u.EXPECT().Register(gomock.Any(), "test", "test", "test").Return(&user.User{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}, nil)

		user, err := resolvers.Register(nil, u, "test", "test", "test")
		a.NoError(err)
		a.NotNil(user)
		a.Equal("test", user.Name)

	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		u := mocks.NewMockUserServiceImpl(ctrl)
		u.EXPECT().Register(gomock.Any(), "test", "test", "test").Return(nil, assert.AnError)

		user, err := resolvers.Register(nil, u, "test", "test", "test")
		a.Error(err)
		a.Nil(user)

	})
}

func TestSignIn(t *testing.T) {
	t.Run("should sign in user", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		u := mocks.NewMockUserServiceImpl(ctrl)
		j := mocks.NewMockJWTServiceImpl(ctrl)
		u.EXPECT().SignIn(gomock.Any(), "test", "test").Return(&user.User{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}, nil)
		j.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).Return("test", nil)

		payload, err := resolvers.SignIn(nil, u, j, "test", "test")
		a.NoError(err)
		a.NotNil(payload)
		a.Equal("test", payload.User.Name)
		a.Equal("test", payload.Token)

	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		u := mocks.NewMockUserServiceImpl(ctrl)
		j := mocks.NewMockJWTServiceImpl(ctrl)
		u.EXPECT().SignIn(gomock.Any(), "test", "test").Return(nil, assert.AnError)

		payload, err := resolvers.SignIn(nil, u, j, "test", "test")
		a.Error(err)
		a.Nil(payload)
	})
}
