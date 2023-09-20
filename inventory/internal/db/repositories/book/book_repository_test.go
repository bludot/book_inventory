package book_test

import (
	"github.com/bludot/tempmee/inventory/config"
	"github.com/bludot/tempmee/inventory/internal/db"
	"github.com/bludot/tempmee/inventory/internal/db/repositories/book"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestBookRepository_FindAll(t *testing.T) {
	t.Run("should return all books", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			bookRepo := book.NewBookRepository(txdb)
			_, _ = bookRepo.Create(nil, book.Book{
				Title:  "title",
				Author: "author",
				Price:  1000,
			})

			books, err := bookRepo.FindAll(nil)
			a.NoError(err)
			a.Len(*books, 1)
			a.Equal("title", (*books)[0].Title)
			a.Equal("author", (*books)[0].Author)
			a.Equal(int64(1000), (*books)[0].Price)
			return nil

		})
	})

	t.Run("should return empty array if no books", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			bookRepo := book.NewBookRepository(txdb)
			books, err := bookRepo.FindAll(nil)
			a.NoError(err)
			a.Len(*books, 0)
			return nil

		})
	})
}

func TestBookRepository_FindById(t *testing.T) {
	t.Run("should return book with id", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			bookRepo := book.NewBookRepository(txdb)
			createdBook, _ := bookRepo.Create(nil, book.Book{
				Title:  "title",
				Author: "author",
				Price:  1000,
			})

			foundBook, err := bookRepo.FindById(nil, createdBook.ID)
			a.NoError(err)
			a.Equal(createdBook.ID, foundBook.ID)
			a.Equal("title", foundBook.Title)
			a.Equal("author", foundBook.Author)
			a.Equal(int64(1000), foundBook.Price)

			return nil

		})
	})

	t.Run("should return error if no book with id", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			bookRepo := book.NewBookRepository(txdb)
			_, err := bookRepo.FindById(nil, "no-id")
			a.Error(err)
			return nil

		})
	})
}

func TestBookRepository_FindByTitle(t *testing.T) {
	t.Run("should return book with title", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			bookRepo := book.NewBookRepository(txdb)
			createdBook, _ := bookRepo.Create(nil, book.Book{
				Title:  "title",
				Author: "author",
				Price:  1000,
			})

			foundBook, err := bookRepo.FindByTitle(nil, createdBook.Title)
			a.NoError(err)
			a.Equal(createdBook.ID, foundBook.ID)
			a.Equal("title", foundBook.Title)
			a.Equal("author", foundBook.Author)
			a.Equal(int64(1000), foundBook.Price)

			return nil

		})
	})

	t.Run("should return error if no book with title", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			bookRepo := book.NewBookRepository(txdb)
			_, err := bookRepo.FindByTitle(nil, "no-title")
			a.Error(err)
			return nil

		})
	})
}

func TestBookRepository_FindByAuthor(t *testing.T) {
	t.Run("should return book with author", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			bookRepo := book.NewBookRepository(txdb)
			createdBook, _ := bookRepo.Create(nil, book.Book{
				Title:  "title",
				Author: "author",
				Price:  1000,
			})

			foundBooks, err := bookRepo.FindByAuthor(nil, createdBook.Author)
			a.NoError(err)
			a.Len(*foundBooks, 1)
			a.Equal(createdBook.ID, (*foundBooks)[0].ID)
			a.Equal("title", (*foundBooks)[0].Title)
			a.Equal("author", (*foundBooks)[0].Author)
			a.Equal(int64(1000), (*foundBooks)[0].Price)

			return nil
		})
	})
}

func TestBookRepository_Create(t *testing.T) {
	t.Run("should create book with title", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			bookRepo := book.NewBookRepository(txdb)
			createdBook, err := bookRepo.Create(nil, book.Book{
				Title:  "title",
				Author: "author",
				Price:  1000,
			})
			a.NoError(err)
			a.Equal("title", createdBook.Title)
			a.Equal("author", createdBook.Author)
			a.Equal(int64(1000), createdBook.Price)

			return nil
		})
	})

	t.Run("should return book if book with title exists", func(t *testing.T) {
		a := assert.New(t)
		cfg := config.LoadConfigOrPanic()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			bookRepo := book.NewBookRepository(txdb)
			createdBook, err := bookRepo.Create(nil, book.Book{
				Title:  "title",
				Author: "author",
				Price:  1000,
			})

			a.NoError(err)
			a.Equal("title", createdBook.Title)
			a.Equal("author", createdBook.Author)
			a.Equal(int64(1000), createdBook.Price)

			createdBook, err = bookRepo.Create(nil, book.Book{
				Title:  "title",
				Author: "author",
				Price:  1000,
			})

			a.NoError(err)
			a.Equal("title", createdBook.Title)
			a.Equal("author", createdBook.Author)
			a.Equal(int64(1000), createdBook.Price)

			allBooks, err := bookRepo.FindAll(nil)
			a.NoError(err)
			a.Len(*allBooks, 1)

			return nil
		})
	})
}
