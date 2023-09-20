package book

import (
	"context"
	"github.com/bludot/tempmee/inventory/internal/db"
)

type BookRepositoryImpl interface {
	FindAll(ctx context.Context) (*[]Book, error)
	FindById(ctx context.Context, id string) (*Book, error)
	FindByTitle(ctx context.Context, title string) (*Book, error)
	FindByAuthor(ctx context.Context, author string) (*[]Book, error)
	Create(ctx context.Context, book Book) (*Book, error)
	GetBooksByIDs(ctx context.Context, ids []string) (*[]Book, error)
}

type BookRepository struct {
	db *db.DB
}

func NewBookRepository(db *db.DB) BookRepositoryImpl {
	return &BookRepository{db: db}
}

func (r *BookRepository) FindAll(ctx context.Context) (*[]Book, error) {
	var books []Book
	err := r.db.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return &books, err
}

func (r *BookRepository) FindById(ctx context.Context, id string) (*Book, error) {
	var book Book
	err := r.db.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, err
}

func (r *BookRepository) FindByTitle(ctx context.Context, title string) (*Book, error) {
	var book Book
	err := r.db.DB.Where("title = ?", title).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, err
}

func (r *BookRepository) FindByAuthor(ctx context.Context, author string) (*[]Book, error) {
	var book []Book
	err := r.db.DB.Where("author = ?", author).Find(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, err
}

func (r *BookRepository) Create(ctx context.Context, book Book) (*Book, error) {
	// check if book with that name exists
	foundBook, err := r.FindByTitle(ctx, book.Title)
	if err == nil {
		// idempotence, return found book
		return foundBook, err
	}
	err = r.db.DB.Create(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, err
}

func (r *BookRepository) GetBooksByIDs(ctx context.Context, ids []string) (*[]Book, error) {
	var books []Book
	err := r.db.DB.Where("id in (?)", ids).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return &books, err
}
