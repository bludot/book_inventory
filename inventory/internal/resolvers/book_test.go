package resolvers_test

import (
	"github.com/bludot/tempmee/inventory/internal/db/repositories/book"
	"github.com/bludot/tempmee/inventory/internal/resolvers"
	"github.com/bludot/tempmee/inventory/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

type BookTest struct {
	ID     string
	Title  string
	Author string
	Price  int
}

func TestBookByID(t *testing.T) {
	t.Run("should return book by id", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBookById(gomock.Any(), "1").Return(&book.Book{
			Title:  "title",
			Author: "author",
			Price:  1000,
		}, nil)

		book, err := resolvers.BookByID(nil, bookService, "1")
		a.NoError(err)
		a.Equal("title", book.Title)
		a.Equal("author", book.Author)
		a.Equal(1000, book.Price)

	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBookById(gomock.Any(), "1").Return(nil, assert.AnError)

		book, err := resolvers.BookByID(nil, bookService, "1")
		a.Error(err)
		a.Nil(book)
	})
}

func TestBookByTitle(t *testing.T) {
	t.Run("should return book by title", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBookByTitle(gomock.Any(), "title").Return(&book.Book{
			Title:  "title",
			Author: "author",
			Price:  1000,
		}, nil)

		book, err := resolvers.BookByTitle(nil, bookService, "title")
		a.NoError(err)
		a.Equal("title", book.Title)
		a.Equal("author", book.Author)
		a.Equal(1000, book.Price)
	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBookByTitle(gomock.Any(), "title").Return(nil, assert.AnError)

		book, err := resolvers.BookByTitle(nil, bookService, "title")
		a.Error(err)
		a.Nil(book)

	})
}

func TestBooksByAuthor(t *testing.T) {
	t.Run("should return books by author", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBooksByAuthor(gomock.Any(), "author").Return(&[]book.Book{
			{
				Title:  "title",
				Author: "author",
				Price:  1000,
			},
		}, nil)

		books, err := resolvers.BooksByAuthor(nil, bookService, "author")
		a.NoError(err)
		a.Len(books, 1)
		a.Equal("title", books[0].Title)
		a.Equal("author", books[0].Author)
		a.Equal(1000, books[0].Price)

	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBooksByAuthor(gomock.Any(), "author").Return(nil, assert.AnError)

		books, err := resolvers.BooksByAuthor(nil, bookService, "author")
		a.Error(err)
		a.Nil(books)

	})
}

func TestMap(t *testing.T) {
	t.Run("should map", func(t *testing.T) {
		a := assert.New(t)

		books := []book.Book{
			{
				ID:     "1",
				Title:  "title",
				Author: "author",
				Price:  1000,
			},
			{
				ID:     "2",
				Title:  "title2",
				Author: "author2",
				Price:  2000,
			},
		}

		mappedBooks := resolvers.Map[book.Book, *BookTest](books, func(b book.Book) *BookTest {
			return &BookTest{
				ID:     b.ID,
				Title:  b.Title,
				Author: b.Author,
				Price:  int(b.Price),
			}
		})

		a.Len(mappedBooks, 2)
		a.Equal("1", mappedBooks[0].ID)
		a.Equal("title", mappedBooks[0].Title)
		a.Equal("author", mappedBooks[0].Author)
		a.Equal(1000, mappedBooks[0].Price)
		a.Equal("2", mappedBooks[1].ID)
		a.Equal("title2", mappedBooks[1].Title)

	})
}

func TestBooks(t *testing.T) {
	t.Run("should return books", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindAllBooks(gomock.Any()).Return(&[]book.Book{
			{
				Title:  "title",
				Author: "author",
				Price:  1000,
			},
		}, nil)

		books, err := resolvers.Books(nil, bookService)
		a.NoError(err)
		a.Len(books, 1)
		a.Equal("title", books[0].Title)
		a.Equal("author", books[0].Author)
		a.Equal(1000, books[0].Price)
	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindAllBooks(gomock.Any()).Return(nil, assert.AnError)

		books, err := resolvers.Books(nil, bookService)
		a.Error(err)
		a.Nil(books)

	})
}
