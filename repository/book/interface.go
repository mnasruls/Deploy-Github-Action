package book

import (
	_entities "be7/layered/entities"
)

type BookRepositoryInterface interface {
	GetAll() ([]_entities.Book, error)
	GetBook(id int) (_entities.Book, int, error)
	DeleteBook(id int) (_entities.Book, error)
	CreateBook(book _entities.Book) (_entities.Book, error)
	UpdatedBook(books _entities.Book) (_entities.Book, error)
}
