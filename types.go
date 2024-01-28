package main

import (
	"math/rand"
	"time"
)

type Blog struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	HeroImage string `json:"heroImage"`
	Date      string `json:"date"`
}

func NewBlog(title, content, heroImage string) *Blog {
	t := time.Now()
	return &Blog{
		Id:        rand.Intn(1000),
		Title:     title,
		Content:   content,
		HeroImage: heroImage,
		Date:      t.UTC().String(),
	}
}
