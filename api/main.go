package main

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

// Think of it like a database in below json
var books = []book {
	{Id: "1", Title: "In Search of Lost Time", Author: "James", Quantity: 5},
	{Id: "2", Title: "The Great Gatsby", Author: "Micheal", Quantity: 5},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing query"})
		return 
	}

	book, err := getBooksById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)	
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBooksById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func checkOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing query"})
		return 
	}

	book, err := getBooksById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not found"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func getBooksById(id string) (*book, error) {
	for i, b := range books {
		if b.Id == id {
			return &books[i], nil
		}
	}
	err := errors.New("book not found")
	if err != nil {
		return nil, err
	}
	return nil, err
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)

	// Update
	router.PATCH("/checkout", checkOutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}