package user

import (
	"errors"
	"github.com/bludot/tempmee/user/internal/db"
)

type UserRepositoryImpl interface {
	FindAll() ([]*User, error)
	FindById(id string) (*User, error)
	FindByName(name string) (*User, error)
	FindByEmail(email string) (*User, error)
	CreateUser(name string, email string, password string) (*User, error)
}

type UserRepository struct {
	db *db.DB
}

func NewUserRepository(db *db.DB) UserRepositoryImpl {
	return &UserRepository{db: db}
}

func (a *UserRepository) CreateUser(name string, email string, password string) (*User, error) {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	// find user before adding
	_, err := a.FindByEmail(email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	err = a.db.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *UserRepository) FindAll() ([]*User, error) {
	var user []*User
	err := a.db.DB.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *UserRepository) FindById(id string) (*User, error) {
	var user User
	err := a.db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserRepository) FindByName(name string) (*User, error) {
	var user User
	err := a.db.DB.Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := a.db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
