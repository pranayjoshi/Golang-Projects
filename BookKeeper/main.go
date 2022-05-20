package main

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model
	Title      string
	Author     string
	CallNumber string `gorm:"unique_index"`
	PersonID   int
}
