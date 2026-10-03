package main

import (
	"database/sql"
	"finance-tracker/internal"
	"finance-tracker/internal/storage"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const helpMessage = `
Finance tracker usage:

	ftrack <command> [argument]

Available commands:
	add [-e add expense | -i add income | -c add category]`

func main() {
	db, err := sql.Open("sqlite3", "./tracker.db")
	if err != nil {
		log.Fatalf("open database failed: %v", err)
	}
	flag.Usage = func() {
		fmt.Println(helpMessage)
	}
	categoryStorage, err := storage.NewCategoryStorage(db)
	if err != nil {
		log.Fatalf("create category storage failed: %v", err)
	}
	app := internal.NewApp(categoryStorage)
	if len(os.Args) > 1 && os.Args[1] == "add" {
		fs := flag.NewFlagSet("add", flag.ContinueOnError)
		var expense = fs.Bool("e", false, "expense")
		var income = fs.Bool("i", false, "income")
		var category = fs.Bool("c", false, "category")
		var name = fs.String("N", "", "category name")
		var amount = fs.String("amount", "", "expense/income amount")
		err := fs.Parse(os.Args[2:])
		if err != nil {
			log.Fatalf("failing to parse arguments: %v", err)
		}
		switch {
		case *expense:
			app.AddExpense(*name, *amount)
		case *income:
			app.AddIncome(*name, *amount)
		case *category:
			if err := app.AddCategory(*name); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		fmt.Println("Try -help to get more information")
	}
}
