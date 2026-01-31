package services

import (
	"time"

	"gorm.io/gorm"

	"crypto/sha256"
	"encoding/hex"

	types "go_webserv/internal/types"
)

type PostService struct {
	dbService *DbService
}

func InitPostService(dbService *DbService) *PostService {
	return &PostService{
		dbService: dbService,
	}
}

func Sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func (b *PostService) CreateNewPost(newPost types.PostItem, id string) error {
	post := types.Post{
		PostID:       "blg:" + id,
		AuthorID:     newPost.AuthorID,
		Title:        newPost.Title,
		CreatedAt:    time.Now(),
		LastModified: time.Now(),
		Tags:         newPost.Tags,
	}

	err := b.dbService.doInsert(&post)

	if err != nil {
		println("Error creating post:", err)
		return err
	}

	return nil
}

func (b *PostService) QueryPost(id string, postType types.PostType) (*types.PostItem, error) {
	post := types.Post{}

	err := b.dbService.doQueryByID(&post, id)
	if err != nil {
		return nil, err
	}

	if post.PostID == "" {
		return nil, gorm.ErrRecordNotFound
	}

	postItem := &types.PostItem{
		AuthorID:  post.AuthorID,
		Title:     post.Title,
		CreatedAt: post.CreatedAt,
		Tags:      post.Tags,
	}

	return postItem, nil
}

func (b *PostService) QueryAllPosts(postType types.PostType) ([]types.PostItem, error) {
	var posts []types.Post
	err := b.dbService.doQueryAll(posts, postType)
	if err != nil {
		return nil, err
	}

	return convertToPostItem(posts), nil
}

func convertToPostItem(posts []types.Post) []types.PostItem {
	var postItems []types.PostItem
	
	for _, post := range posts {
		postItem := types.PostItem{
			AuthorID:  post.AuthorID,
			Title:     post.Title,
			CreatedAt: post.CreatedAt,
			Tags:      post.Tags,
		}
		postItems = append(postItems, postItem)
	}
	
	return postItems
}
