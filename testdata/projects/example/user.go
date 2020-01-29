package main

import "time"

type User struct {
	ID string `api:"users"`

	// Attributes
	Username  string    `json:"username" api:"attr"`
	CreatedAt time.Time `json:"created-at" api:"attr"`
	Admin     bool      `json:"admin" api:"attr"`

	// Relationships
	Articles string `json:"articles" api:"rel,articles,author"`
}
