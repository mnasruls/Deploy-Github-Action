package book

import (
	_entities "be7/layered/entities"

	"gorm.io/gorm"
)

type BookRepository struct {
	database *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		database: db,
	}
}

func (br *BookRepository) GetAll() ([]_entities.Book, error) {
	var books []_entities.Book
	tx := br.database.Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return books, nil
}

func (br *BookRepository) GetBook(id int) (_entities.Book, int, error) {
	var books _entities.Book
	tx := br.database.Find(&books, id)
	if tx.Error != nil {
		return books, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return books, 0, tx.Error
	}
	return books, int(tx.RowsAffected), nil
}
func (br *BookRepository) DeleteBook(id int) (_entities.Book, error) {
	var books _entities.Book
	tx := br.database.Delete(&books, id)
	if tx.Error != nil {
		return books, tx.Error
	}
	if tx.RowsAffected == 0 {
		return books, tx.Error
	}
	return books, nil
}
func (br *BookRepository) CreateBook(book _entities.Book) (_entities.Book, error) {
	var books _entities.Book
	books = book
	tx := br.database.Save(&books)
	if tx.Error != nil {
		return books, tx.Error
	}
	if tx.RowsAffected == 0 {
		return books, tx.Error
	}
	return books, nil
}

func (br *BookRepository) UpdatedBook(books _entities.Book) (_entities.Book, error) {
	tx := br.database.Save(&books)
	if tx.Error != nil {
		return books, tx.Error
	}
	if tx.RowsAffected == 0 {
		return books, tx.Error
	}
	return books, nil
}
