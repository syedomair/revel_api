package controllers

import (
    "github.com/revel/revel"
    "myapp8/app/models"
    "fmt"
)

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {
var err error

	_, err = insertBook("Test 1")

	if err != nil {
	//	panic(err)
	}


    books, err := allBooks()
	if err != nil {
		panic(err)
	}

    return c.Render(books)
    //return c.Render()
}	

func allBooks() ([]models.Book, error) {
	//Retrieve
	books := []models.Book{}

	rows, err := db.Query(`SELECT id, name FROM books order by id`)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id int
			var name string

			err = rows.Scan(&id, &name)
			if err == nil {
				currentBook := models.Book{ID: id, Name: name}
				books = append(books, currentBook)
			} else {
				return books, err
			}
		}
	} else {
		return books, err
	}

	return books, err
}

func insertBook(name string) (int, error) {
	//Create
	var bookID int
	err := db.QueryRow(`INSERT INTO books(name) VALUES($1) RETURNING id`).Scan(&bookID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", bookID)
	return bookID, err
}
