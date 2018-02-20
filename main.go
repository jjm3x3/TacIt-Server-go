package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	fmt.Println("Hello, World")
	defaultHost := "localhost"
	defaultPort := "5432"
	defaultUser := "gorm"
	defaultDb := "tacit_db"

	db, err := gorm.Open("postgres", "host="+defaultHost+" port="+defaultPort+" user="+defaultUser+" dbname="+defaultDb+" sslmode=disable") // TODO:: enable ssl
	defer db.Close()

	if err != nil {
		fmt.Println("There was an error opeing the db: ", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
}
