package inventory

import "github.com/bludot/tempmee/inventory/internal/db/repositories/book"

type InventoryImpl interface {
	FindAllBooks() (*[]book.Book, error)
	FindBookById(id string) (*book.Book, error)
	FindBookByTitle(title string) (*book.Book, error)
	FindBooksByAuthor(author string) (*[]book.Book, error)
}

type Inventory struct {
	Repo book.BookRepositoryImpl
}

func NewInventoryService(repo book.BookRepositoryImpl) InventoryImpl {
	return &Inventory{
		Repo: repo,
	}
}

func (i *Inventory) FindAllBooks() (*[]book.Book, error) {
	return i.Repo.FindAll()
}

func (i *Inventory) FindBookById(id string) (*book.Book, error) {
	return i.Repo.FindById(id)
}

func (i *Inventory) FindBookByTitle(title string) (*book.Book, error) {
	return i.Repo.FindByTitle(title)
}

func (i *Inventory) FindBooksByAuthor(author string) (*[]book.Book, error) {
	return i.Repo.FindByAuthor(author)
}
