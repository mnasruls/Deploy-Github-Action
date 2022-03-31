package book

import (
	_entities "be7/layered/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepository{})
		data, err := bookUseCase.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, "Alchemist", data[0].Title)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepositoryError{})
		data, err := bookUseCase.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestGetBook(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepository{})
		data, rows, err := bookUseCase.GetBook(1)
		assert.Nil(t, err)
		assert.Equal(t, "Alchemist", data.Title)
		assert.Equal(t, "Paulo Choelo", data.Author)
		assert.Equal(t, "Mizan", data.Publisher)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetUserError", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepositoryError{})
		data, rows, err := bookUseCase.GetBook(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.Book{}, data)
	})
}

func TestDeleteBook(t *testing.T) {
	t.Run("TestDeleteBookSuccess", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepository{})
		data, err := bookUseCase.DeleteBook(1)
		assert.Nil(t, err)
		assert.Equal(t, "deleted", data.Title)
	})

	t.Run("TestDeleteBookError", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepositoryError{})
		data, err := bookUseCase.DeleteBook(1)
		assert.NotNil(t, err)
		assert.Equal(t, "Alchemist", data.Title)
	})
}

func TestCreateBook(t *testing.T) {
	t.Run("TestCreateBookSuccess", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepository{})
		data, err := bookUseCase.CreateBook(_entities.Book{
			Title: "Alchemist", Author: "Paulo Choelo", Publisher: "Mizan",
		})
		assert.Nil(t, err)
		assert.Equal(t, "Alchemist", data.Title)
		assert.Equal(t, "Paulo Choelo", data.Author)
		assert.Equal(t, "Mizan", data.Publisher)
	})

	t.Run("TestCreateBookError", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepositoryError{})
		data, err := bookUseCase.CreateBook(_entities.Book{
			Title: "Alchemist", Author: "Paulo Choelo", Publisher: "Mizan",
		})
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Book{}, data)
	})
}

func TestUpdatedBook(t *testing.T) {
	t.Run("TestUpdatedBookSuccess", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepository{})
		data, err := bookUseCase.UpdatedBook(_entities.Book{
			Title: "Alchemist", Author: "Paulo Choelo", Publisher: "Mizan",
		}, 1)
		assert.Nil(t, err)
		assert.Equal(t, "Alchemist", data.Title)
		assert.Equal(t, "Paulo Choelo", data.Author)
		assert.Equal(t, "Mizan", data.Publisher)
	})

	t.Run("TestUpdatedBookError", func(t *testing.T) {
		bookUseCase := NewBookUseCase(mockBookRepositoryError{})
		data, err := bookUseCase.UpdatedBook(_entities.Book{
			Title: "Alchemist", Author: "Paulo Choelo", Publisher: "Mizan",
		}, 1)
		assert.NotNil(t, err)
		assert.Equal(t, _entities.Book{
			Title: "Dunia Sophie", Author: "Jostein Gaarder", Publisher: "Mizan",
		}, data)
	})
}

// mock succes
type mockBookRepository struct{}

func (m mockBookRepository) GetAll() ([]_entities.Book, error) {
	return []_entities.Book{
		{Title: "Alchemist", Author: "Paulo Choelo", Publisher: "Mizan"},
	}, nil
}

func (m mockBookRepository) GetBook(id int) (_entities.Book, int, error) {
	return _entities.Book{
		Title: "Alchemist", Author: "Paulo Choelo", Publisher: "Mizan",
	}, 1, nil
}

func (m mockBookRepository) DeleteBook(id int) (_entities.Book, error) {
	return _entities.Book{
		Title: "deleted", Author: "deleted", Publisher: "deleted",
	}, nil
}

func (m mockBookRepository) UpdatedBook(user _entities.Book) (_entities.Book, error) {
	return _entities.Book{
		Title: "Alchemist", Author: "Paulo Choelo", Publisher: "Mizan",
	}, nil
}
func (m mockBookRepository) CreateBook(user _entities.Book) (_entities.Book, error) {
	return _entities.Book{
		Title: "Alchemist", Author: "Paulo Choelo", Publisher: "Mizan",
	}, nil
}

//  mock error

type mockBookRepositoryError struct{}

func (m mockBookRepositoryError) GetAll() ([]_entities.Book, error) {
	return nil, fmt.Errorf("error")
}
func (m mockBookRepositoryError) GetBook(id int) (_entities.Book, int, error) {
	return _entities.Book{}, 0, fmt.Errorf("error get data user")
}
func (m mockBookRepositoryError) DeleteBook(id int) (_entities.Book, error) {
	return _entities.Book{
		Title: "Alchemist", Author: "Paulo Choelo", Publisher: "Mizan",
	}, fmt.Errorf("error delete data user")
}
func (m mockBookRepositoryError) CreateBook(_entities.Book) (_entities.Book, error) {
	return _entities.Book{}, fmt.Errorf("error create data user")
}
func (m mockBookRepositoryError) UpdatedBook(_entities.Book) (_entities.Book, error) {
	return _entities.Book{
		Title: "Dunia Sophie", Author: "Jostein Gaarder", Publisher: "Mizan",
	}, fmt.Errorf("error updated data user")
}
