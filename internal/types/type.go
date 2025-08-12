package types

import (
	"time"
)

type Blog struct {
	BlogID uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	Subject string
	ContentPath string
}

type Project struct {
	ProjectID uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	ProjectName string
	ContentPath string
}

type BlogItem struct {
	Subject string
	ContentPath string
}

type ProjectItem struct {
	Subject string
	CointentPath string
}
