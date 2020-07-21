package models

import (
	"html"
	"strings"
	"time"
)

//
type Comments struct {
	Content   string `gorm:"size:255;not null;" json:"content"`
	CreatedAt time.Time
	Author    User `json:"author"`
}

//
func (c *Comments) Prepare() {
	c.Content = html.EscapeString(strings.TrimSpace(c.Content))
	c.Author = User{}
}
