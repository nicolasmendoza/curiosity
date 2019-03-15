package rss

import (
	"curiosity/db"
	"html/template"
)

type document struct {
	Title string
	Description template.HTML
	//Author      string
	//Category    string
	PubDate string
	Done bool
}

func (d *document) save() error{
	db.DB
}