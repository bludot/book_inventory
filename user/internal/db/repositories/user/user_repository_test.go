package user_test

import (
	"github.com/bludot/tempmee/user/config"
	"github.com/bludot/tempmee/user/internal/db"
	"github.com/bludot/tempmee/user/internal/db/repositories/user"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	t.Run("should create user", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			userRepo := user.NewUserRepository(txdb)
			user, err := userRepo.CreateUser("test1", "test1", "test1")
			a.NoError(err)
			a.NotNil(user)
			return nil
		})

	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			userRepo := user.NewUserRepository(txdb)
			_, err := userRepo.CreateUser("test1", "test1", "test1")
			a.NoError(err)

			user, err := userRepo.CreateUser("test1", "test1", "test1")
			a.Error(err)
			a.Nil(user)

			return nil
		})
	})
}

func TestUserRepository_FindByEmail(t *testing.T) {
	t.Run("should find user by email", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			userRepo := user.NewUserRepository(txdb)
			_, err := userRepo.CreateUser("test1", "test1", "test1")
			a.NoError(err)

			user, err := userRepo.FindByEmail("test1")
			a.NoError(err)
			a.NotNil(user)
			a.Equal("test1", user.Name)
			return nil
		})
	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			userRepo := user.NewUserRepository(txdb)

			user, err := userRepo.FindByEmail("test2")
			a.Error(err)
			a.Nil(user)
			return nil
		})
	})
}

func TestUserRepository_FindAll(t *testing.T) {
	t.Run("should find all users", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			userRepo := user.NewUserRepository(txdb)
			_, err := userRepo.CreateUser("test1", "test1", "test1")
			a.NoError(err)

			users, err := userRepo.FindAll()
			a.NoError(err)
			a.NotNil(users)
			a.Equal("test1", users[0].Name)
			return nil
		})
	})

	t.Run("should get no users", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			userRepo := user.NewUserRepository(txdb)

			users, err := userRepo.FindAll()
			a.NoError(err)
			a.Equal(0, len(users))
			return nil
		})
	})

	t.Run("should get multiple users", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			userRepo := user.NewUserRepository(txdb)
			_, err := userRepo.CreateUser("test1", "test1", "test1")
			a.NoError(err)

			_, err = userRepo.CreateUser("test2", "test2", "test2")
			a.NoError(err)

			users, err := userRepo.FindAll()
			a.NoError(err)
			a.NotNil(users)
			a.Equal("test1", users[0].Name)
			a.Equal("test2", users[1].Name)
			return nil
		})
	})
}

func TestUserRepository_FindById(t *testing.T) {
	t.Run("should find user by id", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			userRepo := user.NewUserRepository(txdb)
			createdUser, err := userRepo.CreateUser("test1", "test1", "test1")
			a.NoError(err)

			user, err := userRepo.FindById(createdUser.ID)
			a.NoError(err)
			a.NotNil(user)
			a.Equal("test1", user.Name)
			return nil
		})
	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			userRepo := user.NewUserRepository(txdb)

			user, err := userRepo.FindById("2")
			a.Error(err)
			a.Nil(user)
			return nil
		})
	})
}

func TestUserRepository_FindByName(t *testing.T) {
	t.Run("should find user by name", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()

			userRepo := user.NewUserRepository(txdb)
			_, err := userRepo.CreateUser("test1", "test1", "test1")
			a.NoError(err)

			user, err := userRepo.FindByName("test1")
			a.NoError(err)
			a.NotNil(user)
			a.Equal("test1", user.Name)
			return nil
		})
	})

	t.Run("should return error", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()
		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()

			userRepo := user.NewUserRepository(txdb)
			user, err := userRepo.FindByName("test1")
			a.Error(err)
			a.Nil(user)
			return nil
		})
	})
}
