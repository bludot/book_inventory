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

func FindOrderByProducts(ctx context.Context, bookService book.BookServiceImpl, products []string) (*model.Order, error) {
	books, err := bookService.FindBooksByIDs(ctx, products)
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

	var totalPrice int
	for _, book := range mappedBooks {
		totalPrice += book.Price

	}

	return &model.Order{
		Products: products,
		Books:    mappedBooks,
	}, nil
}

func CreateBook(ctx context.Context, bookService book.BookServiceImpl, input model.CreateBookInput) (*model.Book, error) {
	book, err := bookService.CreateBook(ctx, bookRepository.Book{
		Title:  input.Title,
		Author: input.Author,
		Price:  int64(input.Price),
	})

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

func Map[T any, U any](arr []T, f func(T) U) []U {
	var newArr []U
	for _, v := range arr {
		newArr = append(newArr, f(v))
	}
	return newArr
}
