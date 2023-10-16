package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	r := gin.Default()

	db, err := sqlx.Connect("mysql", "test:test@(localhost:3306)/test")

	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	if err != nil {
		log.Fatalln(err)
	}

	r.GET("/list", NewsList)

	r.POST("/edit/:Id", ChangeNewForId)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// NewsList Список Новостей
func NewsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// ChangeNewForId изменение новости по ID
func ChangeNewForId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Бд mysql

// sqlx

// переменные окружения использовать для подключения к базе
