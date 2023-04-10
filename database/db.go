package database

import (
	"challenge-3-chapter-2/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "123456"
	port     = 5432
	dbname   = "db-book-sql"
)

var (
	db  *sql.DB
	err error
)

func StartDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully connected to database")
}

func GetBookAllDB(book []models.Book) (books []models.Book, err error) {
	sqlStatement := `SELECT * from books`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var getBook models.Book

		err = rows.Scan(&getBook.ID, &getBook.NameBook, &getBook.Author)

		if err != nil {
			log.Print("Error query db")
			return
		}

		books = append(books, getBook)
	}

	return
}

func GetBookIdDB(bookID int, book models.Book) (bookDatas models.Book, err error) {
	sqlStatement := `SELECT * FROM books WHERE id = $1`
	err = db.QueryRow(sqlStatement, bookID).Scan(&bookDatas.ID, &bookDatas.NameBook, &bookDatas.Author)
	if err != nil {
		return
	}

	return
}

func CreateBookDB(book models.Book) (newBook models.Book, err error) {
	var bookID int

	getID := `SELECT id FROM books ORDER BY id DESC LIMIT 1`
	err = db.QueryRow(getID).Scan(&bookID)
	if err != nil {
		book.ID = 1
	}

	book.ID = bookID + 1

	sqlStatement := `
	INSERT INTO books (id, name_book, author)
	VALUES ($1, $2, $3)
	Returning *
	`

	_, err = db.Exec(sqlStatement, book.ID, book.NameBook, book.Author)
	if err != nil {
		return
	}

	newBook = book
	return
}

func UpdateBookDB(bookID int, updateBook models.Book) (book models.Book, err error) {

	findId := `SELECT id FROM books WHERE id = $1`
	err = db.QueryRow(findId, bookID).Scan(&bookID)
	if err != nil {
		return
	}

	sqlStatement := `UPDATE books SET name_book = $2, author = $3 WHERE id = $1`
	_, err = db.Exec(sqlStatement, bookID, updateBook.NameBook, updateBook.Author)
	if err != nil {
		return
	}

	book = updateBook
	book.ID = bookID
	return
}

func DeleteBookDB(bookID int) (err error) {

	findId := `SELECT id FROM books WHERE id = $1`
	err = db.QueryRow(findId, bookID).Scan(&bookID)
	if err != nil {
		return
	}

	sqlStatement := `DELETE FROM books WHERE id = $1`
	_, err = db.Exec(sqlStatement, bookID)
	if err != nil {
		panic(err)
	}

	return

}
