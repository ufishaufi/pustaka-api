package main

import (
	"fmt"
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

	// CRUD

	// ===========
	// CREATE DATA
	// ===========
	/*
		book := book.Book{}
		book.Title = "Atomic habits"
		book.Price = 120000
		book.Discount = 15
		book.Rating = 4
		book.Description = "Buku self development tentang kebiasaan baik dan menghilangkan kebiasaan buruk"

		err = db.Create(&book).Error

		if err != nil {
			fmt.Println("==========================")
			fmt.Println("Error creating book record")
			fmt.Println("==========================")
		}
	*/

	// ========
	// GET DATA
	// ========

	var book book.Book

	err = db.Debug().Where("id = ?", 1).First(&book).Error

	if err != nil {
		fmt.Println("=========================")
		fmt.Println("Error finding book record")
		fmt.Println("=========================")
	}

	book.Title = "Man Tiger (Revised edition)"
	err = db.Save(&book).Error

	if err != nil {
		fmt.Println("=========================")
		fmt.Println("Error updating book record")
		fmt.Println("=========================")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)

	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8888")
}
