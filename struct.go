package main

import (
	"time"
)

type (
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	}

	Article struct {
		Source      Source    `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		Content     string    `json:"content"`
	}

	Results struct {
		Status       string    `json:"status"`
		TotalResults int       `json:"totalResults"`
		Articles     []Article `json:"articles"`
	}
)

type (
	Search struct {
		SearchKey  string
		NextPage   int
		TotalPages int
		Results    Results
	}
)
