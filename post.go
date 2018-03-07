package main

import "github.com/jinzhu/gorm"

type post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}
