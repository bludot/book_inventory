package book

import (
	"github.com/bludot/tempmee/inventory/internal/db"
)

type BookRepositoryImpl interface {
	FindAll() ([]Book, error)
	FindById(id string) (Book, error)
	FindByTitle(title string) (Book, error)
	FindByAuthor(author string) (Book, error)
	Create(book Book) (Book, error)
}

type BookRepository struct {
	db *db.DB
}

func NewBookRepository(db *db.DB) BookRepositoryImpl {
	return &BookRepository{db: db}
}

func (r *BookRepository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.DB.Find(&books).Error
	return books, err
}

func (r *BookRepository) FindById(id string) (Book, error) {
	var book Book
	err := r.db.DB.Where("id = ?", id).First(&book).Error
	return book, err
}

func (r *BookRepository) FindByTitle(title string) (Book, error) {
	var book Book
	err := r.db.DB.Where("title = ?", title).First(&book).Error
	return book, err
}

func (r *BookRepository) FindByAuthor(author string) (Book, error) {
	var book Book
	err := r.db.DB.Where("author = ?", author).First(&book).Error
	return book, err
}

func (r *BookRepository) Create(book Book) (Book, error) {
	// check if book with that name exists
	foundBook, err := r.FindByTitle(book.Title)
	if err == nil {
		// idempotence, return found book
		return foundBook, err
	}
	err = r.db.DB.Create(&book).Error
	return book, err
}
