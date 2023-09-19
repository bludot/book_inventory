package book

import (
	"encoding/json"
	"github.com/bludot/tempmee/inventory/config"
	"github.com/bludot/tempmee/inventory/http/http_utils"
	bookRepository "github.com/bludot/tempmee/inventory/internal/db/repositories/book"
	"github.com/bludot/tempmee/inventory/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindAllBooks(t *testing.T) {
	t.Run("should return all books", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		httpRecorder := httptest.NewRecorder()

		cfg := config.LoadConfigOrPanic()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindAllBooks(gomock.Any()).Return(&[]bookRepository.Book{
			{
				ID:     "1",
				Title:  "title",
				Author: "author",
				Price:  1000,
			},
		}, nil)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			BookService: bookService,
		}

		handler := FindAllBooks(httpResolver)

		handler.ServeHTTP(httpRecorder, httptest.NewRequest("GET", "/books", nil))

		a.Equal(200, httpRecorder.Code)
		// json to book
		var books []map[string]interface{}
		bodyString, _ := io.ReadAll(httpRecorder.Body)
		err := json.Unmarshal(bodyString, &books)
		a.NoError(err)
		a.Len(books, 1)
		a.Equal("1", books[0]["id"])
		a.Equal("title", books[0]["title"])
		a.Equal("author", books[0]["author"])
		a.Equal(float64(1000), books[0]["price"])
	})
}

func TestFindBookByID(t *testing.T) {
	t.Run("should return book by id", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := config.LoadConfigOrPanic()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBookById(gomock.Any(), "1").Return(&bookRepository.Book{
			ID:     "1",
			Title:  "title",
			Author: "author",
			Price:  1000,
		}, nil)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			BookService: bookService,
		}

		r := mux.NewRouter()
		RegisterBookRoutes(httpResolver, r)

		ts := httptest.NewServer(r)
		defer ts.Close()

		resp, err := http.Get(ts.URL + "/api/inventory/books/1")
		a.NoError(err)
		a.Equal(200, resp.StatusCode)

		// store response body in map
		var book map[string]interface{}
		bodyString, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(bodyString, &book)
		a.NoError(err)

		a.Equal("1", book["id"])
		a.Equal("title", book["title"])
		a.Equal("author", book["author"])

	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := config.LoadConfigOrPanic()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBookById(gomock.Any(), "1").Return(nil, assert.AnError)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			BookService: bookService,
		}

		r := mux.NewRouter()
		RegisterBookRoutes(httpResolver, r)

		ts := httptest.NewServer(r)
		defer ts.Close()

		resp, err := http.Get(ts.URL + "/api/inventory/books/1")
		a.NoError(err)
		a.Equal(500, resp.StatusCode)

	})
}

func TestFindBookByTitle(t *testing.T) {
	t.Run("should return book by title", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := config.LoadConfigOrPanic()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBookByTitle(gomock.Any(), "title").Return(&bookRepository.Book{
			ID:     "1",
			Title:  "title",
			Author: "author",
			Price:  1000,
		}, nil)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			BookService: bookService,
		}

		r := mux.NewRouter()
		RegisterBookRoutes(httpResolver, r)

		ts := httptest.NewServer(r)
		defer ts.Close()

		resp, err := http.Get(ts.URL + "/api/inventory/books/title/title")
		a.NoError(err)
		a.Equal(200, resp.StatusCode)

		// store response body in map
		var book map[string]interface{}
		bodyString, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(bodyString, &book)
		a.NoError(err)
		a.Equal("1", book["id"])
		a.Equal("title", book["title"])
		a.Equal("author", book["author"])
		a.Equal(float64(1000), book["price"])

	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := config.LoadConfigOrPanic()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBookByTitle(gomock.Any(), "title").Return(nil, assert.AnError)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			BookService: bookService,
		}

		r := mux.NewRouter()
		RegisterBookRoutes(httpResolver, r)

		ts := httptest.NewServer(r)
		defer ts.Close()

		resp, err := http.Get(ts.URL + "/api/inventory/books/title/title")
		a.NoError(err)
		a.Equal(500, resp.StatusCode)

	})
}

func TestFindBooksByAuthor(t *testing.T) {
	t.Run("should return books by author", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := config.LoadConfigOrPanic()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBooksByAuthor(gomock.Any(), "author").Return(&[]bookRepository.Book{
			{
				ID:     "1",
				Title:  "title",
				Author: "author",
				Price:  1000,
			},
			{
				ID:     "2",
				Title:  "title2",
				Author: "author",
				Price:  1000,
			},
		}, nil)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			BookService: bookService,
		}

		r := mux.NewRouter()
		RegisterBookRoutes(httpResolver, r)

		ts := httptest.NewServer(r)
		defer ts.Close()

		resp, err := http.Get(ts.URL + "/api/inventory/books/author/author")
		a.NoError(err)
		a.Equal(200, resp.StatusCode)

		var books []map[string]interface{}
		bodyString, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(bodyString, &books)
		a.NoError(err)
		a.Len(books, 2)
		a.Equal("1", books[0]["id"])
		a.Equal("title", books[0]["title"])
		a.Equal("author", books[0]["author"])
		a.Equal(float64(1000), books[0]["price"])
		a.Equal("2", books[1]["id"])
		a.Equal("title2", books[1]["title"])
		a.Equal("author", books[1]["author"])
		a.Equal(float64(1000), books[1]["price"])

	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := config.LoadConfigOrPanic()

		bookService := mocks.NewMockBookServiceImpl(ctrl)
		bookService.EXPECT().FindBooksByAuthor(gomock.Any(), "author").Return(nil, assert.AnError)

		httpResolver := &http_utils.HTTPResolver{
			Config:      cfg,
			BookService: bookService,
		}

		r := mux.NewRouter()
		RegisterBookRoutes(httpResolver, r)

		ts := httptest.NewServer(r)
		defer ts.Close()

		resp, err := http.Get(ts.URL + "/api/inventory/books/author/author")
		a.NoError(err)
		a.Equal(500, resp.StatusCode)

	})
}
