package types

import (
	"time"
)

type Blog struct {
	BlogID       string    `gorm:"primaryKey;not null" json:"blog_id"`
	Subject      string    `gorm:"not null" json:"subject"`
	ContentPath  string    `gorm:"not null" json:"content_path"`
	CreatedAt    time.Time `gorm:"not null;" json:"created_at"`
	LastModified time.Time `gorm:"not null;" json:"last_modified"`
}

type Project struct {
	ProjectID    string    `gorm:"primaryKey;not null" json:"project_id"`
	Name         string    `gorm:"not null" json:"name"`
	ContentPath  string    `gorm:"not null" json:"content_path"`
	CreatedAt    time.Time `gorm:"not null;" json:"created_at"`
	LastModified time.Time `gorm:"not null;" json:"last_modified"`
}

type BlogItem struct {
	Subject     string `json:"subject"`
	ContentPath string `json:"content_path"`
}

type ProjectItem struct {
	Name        string `json:"name"`
	ContentPath string `json:"content_path"`
}
