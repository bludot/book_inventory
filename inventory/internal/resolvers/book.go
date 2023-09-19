package resolvers

import (
	"context"
	"fmt"
	bookRepository "github.com/bludot/tempmee/inventory/internal/db/repositories/book"
	"github.com/bludot/tempmee/inventory/internal/services/book"
)

func Books(ctx context.Context, bookService book.BookServiceImpl) ([]*bookRepository.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func BookByID(ctx context.Context, bookService book.BookServiceImpl, id string) (*bookRepository.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func BookByTitle(ctx context.Context, bookService book.BookServiceImpl, title string) (*bookRepository.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func BooksByAuthor(ctx context.Context, bookService book.BookServiceImpl, author string) ([]*bookRepository.Book, error) {
	panic(fmt.Errorf("not implemented"))
}
