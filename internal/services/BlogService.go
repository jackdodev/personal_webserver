package services

import (
	"gorm.io/gorm"
	"strconv"

	types "go_webserv/internal/types"
)

type BlogService struct {

}

func InitBlogService() *BlogService {
	return &BlogService{}
}

func (b *BlogService) CreateNewBlog(db *gorm.DB) error {
	return nil
}

func (b *BlogService) QueryBlog(db *gorm.DB, blogId string) (*types.BlogItem, error) {
	blog := types.Blog{}
	blogIdInInt, err := strconv.Atoi(blogId)
	if err != nil {
		return nil, err
	}
	db.First(&blog, "blog_id = ?", blogIdInInt)
	return nil, nil	
}

func (b *BlogService) QueryAllBlogs(db *gorm.DB) ([]types.Blog, error) {
	var blogs []types.Blog
	result := db.Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}

	return blogs, nil
}
