package book_test

import (
	"context"
	bookRepository "github.com/bludot/tempmee/inventory/internal/db/repositories/book"
	"github.com/bludot/tempmee/inventory/internal/services/book"
	"github.com/bludot/tempmee/inventory/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestInventory_FindAllBooks(t *testing.T) {
	t.Run("should return all books", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()
		ctx := context.Background()

		bookRepo := mocks.NewMockBookRepositoryImpl(ctrl)

		bookRepo.EXPECT().FindAll().Return(&[]bookRepository.Book{
			{
				Title:  "title",
				Author: "author",
				Price:  1000,
			},
		}, nil)

		inventory := book.NewBookService(bookRepo)
		books, err := inventory.FindAllBooks(ctx)
		a.NoError(err)
		a.Len(*books, 1)
		a.Equal("title", (*books)[0].Title)
		a.Equal("author", (*books)[0].Author)
		a.Equal(int64(1000), (*books)[0].Price)
	})

	t.Run("should return empty array if no books", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		ctx := context.Background()

		bookRepo := mocks.NewMockBookRepositoryImpl(ctrl)

		bookRepo.EXPECT().FindAll().Return(&[]bookRepository.Book{}, nil)

		inventory := book.NewBookService(bookRepo)
		books, err := inventory.FindAllBooks(ctx)
		a.NoError(err)
		a.Len(*books, 0)

	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		ctx := context.Background()

		bookRepo := mocks.NewMockBookRepositoryImpl(ctrl)

		bookRepo.EXPECT().FindAll().Return(nil, assert.AnError)

		inventory := book.NewBookService(bookRepo)
		books, err := inventory.FindAllBooks(ctx)
		a.Error(err)
		a.Nil(books)

	})

}

func TestInventory_FindBookById(t *testing.T) {
	t.Run("should return book", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		ctx := context.Background()

		bookRepo := mocks.NewMockBookRepositoryImpl(ctrl)

		bookRepo.EXPECT().FindById("id").Return(&bookRepository.Book{
			Title:  "title",
			Author: "author",
			Price:  1000,
		}, nil)

		inventory := book.NewBookService(bookRepo)
		book, err := inventory.FindBookById(ctx, "id")
		a.NoError(err)
		a.Equal("title", book.Title)
		a.Equal("author", book.Author)
		a.Equal(int64(1000), book.Price)

	})

	t.Run("should return error if no book with id", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		ctx := context.Background()

		bookRepo := mocks.NewMockBookRepositoryImpl(ctrl)

		bookRepo.EXPECT().FindById("id").Return(nil, assert.AnError)

		inventory := book.NewBookService(bookRepo)
		book, err := inventory.FindBookById(ctx, "id")
		a.Error(err)
		a.Nil(book)

	})
}

func TestInventory_FindBookByTitle(t *testing.T) {
	t.Run("should return book", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		ctx := context.Background()

		bookRepo := mocks.NewMockBookRepositoryImpl(ctrl)

		bookRepo.EXPECT().FindByTitle("title").Return(&bookRepository.Book{
			Title:  "title",
			Author: "author",
			Price:  1000,
		}, nil)

		inventory := book.NewBookService(bookRepo)
		book, err := inventory.FindBookByTitle(ctx, "title")
		a.NoError(err)
		a.Equal("title", book.Title)
		a.Equal("author", book.Author)
		a.Equal(int64(1000), book.Price)

	})

	t.Run("should return error if no book with title", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		ctx := context.Background()

		bookRepo := mocks.NewMockBookRepositoryImpl(ctrl)

		bookRepo.EXPECT().FindByTitle("title").Return(nil, assert.AnError)

		inventory := book.NewBookService(bookRepo)
		book, err := inventory.FindBookByTitle(ctx, "title")
		a.Error(err)
		a.Nil(book)

	})
}

func TestInventory_FindBooksByAuthor(t *testing.T) {
	t.Run("should return books", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		ctx := context.Background()

		bookRepo := mocks.NewMockBookRepositoryImpl(ctrl)

		bookRepo.EXPECT().FindByAuthor("author").Return(&[]bookRepository.Book{
			{
				Title:  "title",
				Author: "author",
				Price:  1000,
			},
		}, nil)

		inventory := book.NewBookService(bookRepo)
		books, err := inventory.FindBooksByAuthor(ctx, "author")
		a.NoError(err)
		a.Len(*books, 1)
		a.Equal("title", (*books)[0].Title)
		a.Equal("author", (*books)[0].Author)

	})

	t.Run("should return error if no books with author", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		ctx := context.Background()

		bookRepo := mocks.NewMockBookRepositoryImpl(ctrl)

		bookRepo.EXPECT().FindByAuthor("author").Return(nil, assert.AnError)

		inventory := book.NewBookService(bookRepo)
		books, err := inventory.FindBooksByAuthor(ctx, "author")
		a.Error(err)
		a.Nil(books)

	})

}
