package models

type User struct {
	FirstName string
	LastName  string
	Tel       string `default:"0123456789"`
}
