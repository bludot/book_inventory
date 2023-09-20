package book

import (
	"context"
	"github.com/bludot/tempmee/inventory/internal/db/repositories/book"
)

type BookServiceImpl interface {
	FindAllBooks(ctx context.Context) (*[]book.Book, error)
	FindBookById(ctx context.Context, id string) (*book.Book, error)
	FindBookByTitle(ctx context.Context, title string) (*book.Book, error)
	FindBooksByAuthor(ctx context.Context, author string) (*[]book.Book, error)
	CreateBook(ctx context.Context, book book.Book) (*book.Book, error)
	FindBooksByIDs(ctx context.Context, ids []string) (*[]book.Book, error)
}

type BookService struct {
	Repo book.BookRepositoryImpl
}

func NewBookService(repo book.BookRepositoryImpl) BookServiceImpl {
	return &BookService{
		Repo: repo,
	}
}

func (i *BookService) FindAllBooks(ctx context.Context) (*[]book.Book, error) {
	return i.Repo.FindAll(ctx)
}

func (i *BookService) FindBookById(ctx context.Context, id string) (*book.Book, error) {
	return i.Repo.FindById(ctx, id)
}

func (i *BookService) FindBookByTitle(ctx context.Context, title string) (*book.Book, error) {
	return i.Repo.FindByTitle(ctx, title)
}

func (i *BookService) FindBooksByAuthor(ctx context.Context, author string) (*[]book.Book, error) {
	return i.Repo.FindByAuthor(ctx, author)
}

func (i *BookService) CreateBook(ctx context.Context, book book.Book) (*book.Book, error) {
	return i.Repo.Create(ctx, book)
}

func (i *BookService) FindBooksByIDs(ctx context.Context, ids []string) (*[]book.Book, error) {
	return i.Repo.GetBooksByIDs(ctx, ids)
}
