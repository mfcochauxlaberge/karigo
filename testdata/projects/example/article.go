package main

import "time"

type Article struct {
	ID string `api:"users"`

	// Attribute
	Title       string    `json:"title" api:"attr"`
	Content     string    `json:"content" api:"attr"`
	PublishedAt time.Time `json:"published-at" api:"attr"`

	// Relationship
	Author string `json:"author" api:"rel,users,articles"`
}
