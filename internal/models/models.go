// Package models describe entities in the app
package models

type Type string

const (
	Expense = "expense"
	Income  = "income"
)

type Category struct {
	ID   int8
	Name string
}

type CatSaver interface {
	Save()
}

type Entire struct {
	ID       int16
	Type     Type
	Category Category
	Amount   int16
	Date     string
}
