package models

import "time"

type CommentRequest struct {
	ID        string `gorm:"primary_key;auto_increment" json:"id"`
	Content   string `gorm:"size:255;not null;" json:"content"`
	CreatedAt time.Time
	Author    User `json:"author"`
}
