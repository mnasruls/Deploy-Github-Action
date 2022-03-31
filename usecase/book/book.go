package book

import (
	_entities "be7/layered/entities"
	_bookRepository "be7/layered/repository/book"
)

type BookUseCase struct {
	bookRepository _bookRepository.BookRepositoryInterface
}

func NewBookUseCase(bookRepo _bookRepository.BookRepositoryInterface) BookUseCaseInterface {
	return &BookUseCase{
		bookRepository: bookRepo,
	}
}

func (buc *BookUseCase) GetAll() ([]_entities.Book, error) {
	books, err := buc.bookRepository.GetAll()
	return books, err
}
func (buc *BookUseCase) GetBook(id int) (_entities.Book, int, error) {
	books, rows, err := buc.bookRepository.GetBook(id)
	return books, rows, err
}
func (buc *BookUseCase) DeleteBook(id int) (_entities.Book, error) {
	books, err := buc.bookRepository.DeleteBook(id)
	return books, err
}
func (buc *BookUseCase) CreateBook(book _entities.Book) (_entities.Book, error) {
	books, err := buc.bookRepository.CreateBook(book)
	return books, err
}
func (buc *BookUseCase) UpdatedBook(books _entities.Book, id int) (_entities.Book, error) {
	books.ID = uint(id)
	books, err := buc.bookRepository.UpdatedBook(books)
	return books, err
}
