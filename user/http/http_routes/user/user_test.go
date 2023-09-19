package user_test

import (
	"bytes"
	"github.com/bludot/tempmee/user/config"
	userHanlders "github.com/bludot/tempmee/user/http/http_routes/user"
	"github.com/bludot/tempmee/user/http/http_utils"
	userRepo "github.com/bludot/tempmee/user/internal/db/repositories/user"
	"github.com/bludot/tempmee/user/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	t.Run("should return 200 OK", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		httpRecorder := httptest.NewRecorder()

		cfg := config.LoadConfigOrPanic()
		userService := mocks.NewMockUserServiceImpl(ctrl)
		userService.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&userRepo.User{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}, nil)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			UserService: userService,
		}

		handler := userHanlders.RegisterUserHandler(httpResolver)

		bodyByte := []byte(`{"name": "test", "email": "test", "password": "test"}`)
		bodyReader := bytes.NewReader(bodyByte)
		// pass body to handler
		handler.ServeHTTP(httpRecorder, httptest.NewRequest("POST", "/api/user/register", bodyReader))

		a.Equal(200, httpRecorder.Code)
	})

	t.Run("should return 500 Internal Server Error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		httpRecorder := httptest.NewRecorder()

		cfg := config.LoadConfigOrPanic()
		userService := mocks.NewMockUserServiceImpl(ctrl)
		userService.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			UserService: userService,
		}

		handler := userHanlders.RegisterUserHandler(httpResolver)

		bodyByte := []byte(`{"name": "test", "email": "test", "password": "test"}`)
		bodyReader := bytes.NewReader(bodyByte)
		// pass body to handler
		handler.ServeHTTP(httpRecorder, httptest.NewRequest("POST", "/api/user/register", bodyReader))

		a.Equal(500, httpRecorder.Code)
	})

	t.Run("should return 400 Bad Request", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		httpRecorder := httptest.NewRecorder()

		cfg := config.LoadConfigOrPanic()
		userService := mocks.NewMockUserServiceImpl(ctrl)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			UserService: userService,
		}

		handler := userHanlders.RegisterUserHandler(httpResolver)

		// malformed JSON
		bodyByte := []byte(`{"name": "test", "email": "test", }`)
		bodyReader := bytes.NewReader(bodyByte)
		// pass body to handler
		handler.ServeHTTP(httpRecorder, httptest.NewRequest("POST", "/api/user/register", bodyReader))

		a.Equal(400, httpRecorder.Code)
	})
}

func TestSignInUserHandler(t *testing.T) {
	t.Run("should return 200 OK", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		httpRecorder := httptest.NewRecorder()

		cfg := config.LoadConfigOrPanic()
		userService := mocks.NewMockUserServiceImpl(ctrl)
		userService.EXPECT().SignIn(gomock.Any(), gomock.Any(), gomock.Any()).Return(&userRepo.User{
			Name:     "test",
			Email:    "test",
			Password: "test",
		}, nil)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			UserService: userService,
		}

		handler := userHanlders.SignInUserHandler(httpResolver)

		bodyByte := []byte(`{"email": "test", "password": "test"}`)
		bodyReader := bytes.NewReader(bodyByte)
		// pass body to handler
		handler.ServeHTTP(httpRecorder, httptest.NewRequest("POST", "/api/user/signin", bodyReader))

		a.Equal(200, httpRecorder.Code)
	})

	t.Run("should return 401 Internal Server Error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		httpRecorder := httptest.NewRecorder()

		cfg := config.LoadConfigOrPanic()
		userService := mocks.NewMockUserServiceImpl(ctrl)
		userService.EXPECT().SignIn(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			UserService: userService,
		}

		handler := userHanlders.SignInUserHandler(httpResolver)

		bodyByte := []byte(`{"email": "test", "password": "test"}`)
		bodyReader := bytes.NewReader(bodyByte)
		// pass body to handler
		handler.ServeHTTP(httpRecorder, httptest.NewRequest("POST", "/api/user/signin", bodyReader))

		a.Equal(401, httpRecorder.Code)
	})

	t.Run("should return 400 Bad Request", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		httpRecorder := httptest.NewRecorder()

		cfg := config.LoadConfigOrPanic()
		userService := mocks.NewMockUserServiceImpl(ctrl)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			UserService: userService,
		}

		handler := userHanlders.SignInUserHandler(httpResolver)

		// malformed JSON
		bodyByte := []byte(`{"email": "test", "password": "test", }`)
		bodyReader := bytes.NewReader(bodyByte)
		// pass body to handler
		handler.ServeHTTP(httpRecorder, httptest.NewRequest("POST", "/api/user/signin", bodyReader))

		a.Equal(400, httpRecorder.Code)
	})
}
