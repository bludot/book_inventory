package resolvers

import (
	"context"
	"github.com/bludot/tempmee/inventory/graph/model"
	bookRepository "github.com/bludot/tempmee/inventory/internal/db/repositories/book"
	"github.com/bludot/tempmee/inventory/internal/services/book"
)

func Books(ctx context.Context, bookService book.BookServiceImpl) ([]*model.Book, error) {
	books, err := bookService.FindAllBooks(ctx)
	if err != nil {
		return nil, err
	}

	mappedBooks := Map[bookRepository.Book, *model.Book](*books, func(b bookRepository.Book) *model.Book {
		return &model.Book{
			ID:     b.ID,
			Title:  b.Title,
			Author: b.Author,
			Price:  int(b.Price),
		}
	})

	return mappedBooks, nil
}

func BookByID(ctx context.Context, bookService book.BookServiceImpl, id string) (*model.Book, error) {
	book, err := bookService.FindBookById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Book{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Price:  int(book.Price),
	}, nil
}

func BookByTitle(ctx context.Context, bookService book.BookServiceImpl, title string) (*model.Book, error) {
	book, err := bookService.FindBookByTitle(ctx, title)
	if err != nil {
		return nil, err
	}

	return &model.Book{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Price:  int(book.Price),
	}, nil
}

func BooksByAuthor(ctx context.Context, bookService book.BookServiceImpl, author string) ([]*model.Book, error) {
	books, err := bookService.FindBooksByAuthor(ctx, author)
	if err != nil {
		return nil, err
	}

	mappedBooks := Map[bookRepository.Book, *model.Book](*books, func(b bookRepository.Book) *model.Book {
		return &model.Book{
			ID:     b.ID,
			Title:  b.Title,
			Author: b.Author,
			Price:  int(b.Price),
		}
	})

	return mappedBooks, nil
}

// generic map function
func Map[T any, U any](arr []T, f func(T) U) []U {
	var newArr []U
	for _, v := range arr {
		newArr = append(newArr, f(v))
	}
	return newArr
}
