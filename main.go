package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type post struct {
	Title string `json:"title"`
}

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
	r.POST("/note", func(c *gin.Context) {
		var aPost post
		err := c.BindJSON(&aPost)
		if err != nil {
			fmt.Println("has headers: ", c.GetHeader("Content-Type"))
			fmt.Println("There was an error binding to aPost: ", c.PostForm("title"))
			c.JSON(400, gin.H{"Error": "There was an error with what you provided"})
			return
		}
		fmt.Println("Here is the result: ", aPost)
		c.JSON(200, gin.H{"status": "success"})
	})
	r.Run()
}
