package book

import "github.com/bludot/tempmee/inventory/internal/db/repositories/book"

type BookServiceImpl interface {
	FindAllBooks() (*[]book.Book, error)
	FindBookById(id string) (*book.Book, error)
	FindBookByTitle(title string) (*book.Book, error)
	FindBooksByAuthor(author string) (*[]book.Book, error)
}

type BookService struct {
	Repo book.BookRepositoryImpl
}

func NewBookService(repo book.BookRepositoryImpl) BookServiceImpl {
	return &BookService{
		Repo: repo,
	}
}

func (i *BookService) FindAllBooks() (*[]book.Book, error) {
	return i.Repo.FindAll()
}

func (i *BookService) FindBookById(id string) (*book.Book, error) {
	return i.Repo.FindById(id)
}

func (i *BookService) FindBookByTitle(title string) (*book.Book, error) {
	return i.Repo.FindByTitle(title)
}

func (i *BookService) FindBooksByAuthor(author string) (*[]book.Book, error) {
	return i.Repo.FindByAuthor(author)
}
