package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	// Interfaces to work with database
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

// Method Example
func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return Category{}, err
	}
	// Nil -> Error empty
	return Category{ID: id, Name: name, Description: description}, nil
}
