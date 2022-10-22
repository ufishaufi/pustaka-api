package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Shaufi098709@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	// bookFileRepository := book.NewFileRepository()
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	/*
		books, err := bookRepository.FindAll()

		for _, book := range books {
			fmt.Println("Title :", book.Title)
		}
	*/

	/*
		book, err := bookRepository.FindByID(2)

		fmt.Println("Title :", book.Title)
	*/

	/*
		book := book.Book{
			Title:       "$100 Startup",
			Description: "Good book",
			Price:       95000,
			Rating:      4,
			Discount:    0,
		}
	*/

	/*
		bookRequest := book.BookRequest{
			Title: "Gundam",
			Price: "100000",
		}

		// bookRepository.Create(book)
		bookService.Create(bookRequest)
	*/

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run()
}
