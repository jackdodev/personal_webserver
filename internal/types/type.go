package types

import (
	"time"

	"github.com/lib/pq"
)

type Blog struct {
	BlogID       string         `gorm:"primaryKey;not null" json:"blog_id"`
	AuthorID     string         `gorm:"not null" json:"author_id"`
	Subject      string         `gorm:"not null" json:"subject"`
	CreatedAt    time.Time      `gorm:"not null;" json:"created_at"`
	LastModified time.Time      `gorm:"not null;" json:"last_modified"`
	Tags         pq.StringArray `gorm:"type:text[];not null;default:'{}'" json:"tags"`
}

type Project struct {
	ProjectID    string         `gorm:"primaryKey;not null" json:"project_id"`
	Name         string         `gorm:"not null" json:"name"`
	ContentPath  string         `gorm:"not null" json:"content_path"`
	CreatedAt    time.Time      `gorm:"not null;" json:"created_at"`
	LastModified time.Time      `gorm:"not null;" json:"last_modified"`
	Tags         pq.StringArray `gorm:"type:text[];not null;default:'{}'" json:"tags"`
}

type BlogItem struct {
	AuthorID     string    `json:"author_id"`
	Subject      string    `json:"subject"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
	Tags         []string  `json:"tags"`
}

type ProjectItem struct {
	Name        string `json:"name"`
	ContentPath string `json:"content_path"`
}

type DownloadLinkRequest struct {
	Key string `json:"key"`
}

type DownloadLinkResponse struct {
	DownloadURL string `json:"download_url"`
	Key         string `json:"key"`
}

type UploadLinkRequest struct {
	Key string `json:"key"`
}

type UploadLinkResponse struct {
	UploadURL string `json:"upload_url"`
	Key       string `json:"key"`
}
