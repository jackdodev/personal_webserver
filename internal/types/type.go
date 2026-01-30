package types

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	PostID       string         `gorm:"primaryKey;not null" json:"post_id"`
	AuthorID     string         `gorm:"not null" json:"author_id"`
	Title        string         `gorm:"not null" json:"title"`
	CreatedAt    time.Time      `gorm:"not null;" json:"created_at"`
	LastModified time.Time      `gorm:"not null;" json:"last_modified"`
	Tags         pq.StringArray `gorm:"type:text[];not null;default:'{}'" json:"tags"`
}

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
	AuthorID     string         `gorm:"not null" json:"author_id"`
	ProjectName  string         `gorm:"not null" json:"project_name"`
	CreatedAt    time.Time      `gorm:"not null;" json:"created_at"`
	LastModified time.Time      `gorm:"not null;" json:"last_modified"`
	Tags         pq.StringArray `gorm:"type:text[];not null;default:'{}'" json:"tags"`
}

type PostItem struct {
	AuthorID     string    `json:"author_id"`
	Title        string    `json:"subject"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
	Tags         []string  `json:"tags"`
}

type BlogItem struct {
	AuthorID     string    `json:"author_id"`
	Subject      string    `json:"subject"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
	Tags         []string  `json:"tags"`
}

type ProjectItem struct {
	AuthorID     string    `json:"author_id"`
	ProjectName  string    `json:"project_name"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
	Tags         []string  `json:"tags"`
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
