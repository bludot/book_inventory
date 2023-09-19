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
		u.EXPECT().SignIn(gomock.Any(), "test", "test").Return(&user.User{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}, nil)

		user, err := resolvers.SignIn(nil, u, "test", "test")
		a.NoError(err)
		a.NotNil(user)
		a.Equal("test", user.Name)

	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		u := mocks.NewMockUserServiceImpl(ctrl)
		u.EXPECT().SignIn(gomock.Any(), "test", "test").Return(nil, assert.AnError)

		user, err := resolvers.SignIn(nil, u, "test", "test")
		a.Error(err)
		a.Nil(user)
	})
}
