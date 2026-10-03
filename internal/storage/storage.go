// Package storage describe app storage
package storage

import (
	"database/sql"
	"finance-tracker/internal/models"
	"fmt"
)

const initCategoryStorage = "create table if not exists categories (id integer not null primary key, name text)"

type Entities interface {
	models.Category | models.Entire
}

type DBManager[E Entities] interface {
	Save(E) error
	Get(string) (E, error)
}

type AppStorage struct {
	categoryStorage *categoryStorage
}

func NewAppStorage(catStorage *categoryStorage) *AppStorage {
	return &AppStorage{categoryStorage: catStorage}
}

type categoryStorage struct {
	db *sql.DB
}

func NewCategoryStorage(db *sql.DB) (*categoryStorage, error) {
	if _, err := db.Exec(initCategoryStorage); err != nil {
		return nil, fmt.Errorf("failed to init category storage: %w", err)
	}
	return &categoryStorage{db: db}, nil
}

func (cs categoryStorage) Save(cat models.Category) error {
	query := "insert into categories (name) values ($1)"
	if _, err := cs.db.Exec(query, cat.Name); err != nil {
		return fmt.Errorf("failed to save category: %w", err)
	}
	return nil
}

func (cs categoryStorage) Get(name string) (models.Category, error) {
	query := "select * from categories where name = $1"
	var cat models.Category
	if err := cs.db.QueryRow(query, name).Scan(&cat.ID, &cat.Name); err != nil {
		return models.Category{}, fmt.Errorf("failed to get a category: %w", err)
	}
	return cat, nil
}
