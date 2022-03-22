package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {

	db := connect()
	defer db.Close()

	query := ("SELECT * FROM books")

	rows, err := db.Query(query)
	if err != nil {
		return
	}

	var book Book
	var books []Book

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}

	if len(books) != 0 {
		c.IndentedJSON(http.StatusCreated, books)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}

}

func AddBook(c *gin.Context) {

	db := connect()
	defer db.Close()

	var book Book

	err := c.Bind(&book)
	if err != nil {
		fmt.Println(err)
		return
	}

	insert, err := db.Query("INSERT INTO books(Title, Author) VALUES(?, ?)",
		book.Title, book.Author)

	if err != nil {
		panic(err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, insert)
	}

	defer insert.Close()
}

func UpdateBook(c *gin.Context) {

	db := connect()
	defer db.Close()

	var book Book

	err := c.Bind(&book)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, errQuery := db.Exec("UPDATE books SET title= ?, author= ? WHERE ID= ?", book.Title, book.Author, book.ID)

	if errQuery == nil {
		c.IndentedJSON(http.StatusOK, book)
	} else {
		c.AbortWithStatusJSON(400, "Update Failed")
	}

}

func DeleteBook(c *gin.Context) {

	db := connect()
	defer db.Close()

	id := c.Query("ID")

	_, errQuery := db.Exec(`DELETE FROM books WHERE ID=?`, id)

	if errQuery == nil {
		c.IndentedJSON(http.StatusOK, id)
	} else {
		c.AbortWithStatusJSON(400, "Delete Failed")
	}
}
