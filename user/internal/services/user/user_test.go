package user_test

import (
	"context"
	user2 "github.com/bludot/tempmee/user/internal/db/repositories/user"
	"github.com/bludot/tempmee/user/internal/services/user"
	"github.com/bludot/tempmee/user/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserService_GetAllUsers(t *testing.T) {

	t.Run("should return all users", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockUserRepositoryImpl(ctrl)
		r.EXPECT().FindAll().Return([]*user2.User{
			&user2.User{
				Name:     "test",
				Email:    "test",
				Password: "test",
			},
		}, nil)

		userService := user.NewUserService(r)
		ctx := context.Background()

		users, err := userService.GetAllUsers(ctx, 10)
		a.NoError(err)
		a.NotNil(users)
		a.Equal(1, len(users))
		a.Equal("test", users[0].Name)

	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockUserRepositoryImpl(ctrl)
		r.EXPECT().FindAll().Return(nil, assert.AnError)

		userService := user.NewUserService(r)
		ctx := context.Background()

		users, err := userService.GetAllUsers(ctx, 10)
		a.Error(err)
		a.Nil(users)

	})
}

func TestUserService_GetUserByEmail(t *testing.T) {

	t.Run("should return user by email", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockUserRepositoryImpl(ctrl)
		r.EXPECT().FindByEmail("test").Return(&user2.User{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}, nil)

		userService := user.NewUserService(r)
		ctx := context.Background()

		user, err := userService.GetUserByEmail(ctx, "test")
		a.NoError(err)
		a.NotNil(user)
		a.Equal("test", user.Name)

	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockUserRepositoryImpl(ctrl)
		r.EXPECT().FindByEmail("test").Return(nil, assert.AnError)

		userService := user.NewUserService(r)
		ctx := context.Background()

		user, err := userService.GetUserByEmail(ctx, "test")
		a.Error(err)
		a.Nil(user)

	})
}

func TestUserService_Register(t *testing.T) {
	t.Run("should register user", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockUserRepositoryImpl(ctrl)
		r.EXPECT().CreateUser("test", "test", "test").Return(&user2.User{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}, nil)

		userService := user.NewUserService(r)
		ctx := context.Background()

		user, err := userService.Register(ctx, "test", "test", "test")
		a.NoError(err)
		a.NotNil(user)
		a.Equal("test", user.Name)
	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockUserRepositoryImpl(ctrl)
		r.EXPECT().CreateUser("test", "test", "test").Return(nil, assert.AnError)

		userService := user.NewUserService(r)
		ctx := context.Background()

		user, err := userService.Register(ctx, "test", "test", "test")
		a.Error(err)
		a.Nil(user)
	})

}

func TestUserService_SignIn(t *testing.T) {
	t.Run("should sign in user", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockUserRepositoryImpl(ctrl)
		r.EXPECT().FindByEmail("test").Return(&user2.User{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}, nil)

		userService := user.NewUserService(r)
		ctx := context.Background()

		user, err := userService.SignIn(ctx, "test", "test")
		a.NoError(err)
		a.NotNil(user)
		a.Equal("test", user.Name)
	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockUserRepositoryImpl(ctrl)
		r.EXPECT().FindByEmail("test").Return(nil, assert.AnError)

		userService := user.NewUserService(r)
		ctx := context.Background()

		user, err := userService.SignIn(ctx, "test", "test")
		a.Error(err)
		a.Nil(user)
	})

	t.Run("should return nil user", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockUserRepositoryImpl(ctrl)
		r.EXPECT().FindByEmail("test").Return(nil, nil)

		userService := user.NewUserService(r)
		ctx := context.Background()

		user, err := userService.SignIn(ctx, "test", "test")
		a.Error(err)
		a.Nil(user)
	})

	t.Run("should return error for wrong password", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		r := mocks.NewMockUserRepositoryImpl(ctrl)
		r.EXPECT().FindByEmail("test").Return(&user2.User{
			Name:     "test",
			Email:    "test",
			Password: "wrongPass",
		}, nil)

		userService := user.NewUserService(r)
		ctx := context.Background()

		user, err := userService.SignIn(ctx, "test", "test1")
		a.Error(err)
		a.Nil(user)
	})
}
