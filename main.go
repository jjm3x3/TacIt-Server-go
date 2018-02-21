package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

var (
	db *gorm.DB
)

func main() {
	fmt.Println("Hello, World")
	defaultHost := "localhost"
	defaultPort := "5432"
	defaultUser := "gorm"
	defaultDb := "tacit_db"

	var err error
	db, err = gorm.Open("postgres", "host="+defaultHost+" port="+defaultPort+" user="+defaultUser+" dbname="+defaultDb+" sslmode=disable") // TODO:: enable ssl
	defer db.Close()

	if err != nil {
		fmt.Println("There was an error opeing the db: ", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/note", makePost)

	r.Run()
}

func makePost(c *gin.Context) {
	var aPost post
	err := c.BindJSON(&aPost)
	if err != nil {
		fmt.Println("has headers: ", c.GetHeader("Content-Type"))
		var body []byte
		num, err := c.Request.Body.Read(body)
		if num <= 0 { // not sure if this is really an error
			fmt.Println("There was no body provided")
		} else if err != nil {
			fmt.Println("There was an error reading the body: ", err)
		}
		fmt.Println("There was an error binding to aPost: ")
		c.JSON(400, gin.H{"Error": "There was an error with what you provided"})
		return
	}
	fmt.Printf("Here is the result: '%v'\n", aPost)
	db.AutoMigrate(&post{}) // probably doesn't need to happen every time
	db.Create(&aPost)
	c.JSON(200, gin.H{"status": "success"})
}
