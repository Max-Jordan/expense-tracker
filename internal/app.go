// Package internal describe the app
package internal

import (
	"database/sql"
	"errors"
	"finance-tracker/internal/models"
	"finance-tracker/internal/storage"
	"fmt"
	"strings"
)

type App struct {
	categoryStorage storage.DBManager[models.Category]
}

func NewApp(catStorage storage.DBManager[models.Category]) *App {
	return &App{categoryStorage: catStorage}
}

func (a *App) AddExpense(category string, amount string) error {
	return nil
}

func (a *App) AddIncome(category string, amount string) error {
	return nil
}

func (a *App) AddCategory(name string) error {
	exist, err := a.checkCategory(name)
	if err != nil {
		return err
	}
	if exist {
		fmt.Println("Category already exist")
	} else {
		cat := models.Category{Name: strings.ToLower(name)}
		if err := saveEntity(a.categoryStorage, cat); err != nil {
			return fmt.Errorf("saving category failed: %w", err)
		}
	}
	return nil
}

func (a *App) checkCategory(name string) (bool, error) {
	_, err := a.categoryStorage.Get(name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func saveEntity[E storage.Entities](repo storage.DBManager[E], e E) error {
	if err := repo.Save(e); err != nil {
		return fmt.Errorf("savig %T failed: %w", e, err)
	}
	return nil
}


















