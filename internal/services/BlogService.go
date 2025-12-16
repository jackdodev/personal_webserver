package services

import (
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	types "go_webserv/internal/types"
)

type BlogService struct {

}

func InitBlogService() *BlogService {
	return &BlogService{}
}

func Sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func (b *BlogService) CreateNewBlog(db *gorm.DB, newBlog types.BlogItem) error {
	err := db.Create(&types.Blog{
		BlogID: 	 Sha256Hex(fmt.Sprintf("%s-%d", newBlog.Subject, time.Now().UnixNano()))[:16],
		Subject:     newBlog.Subject,
		ContentPath: newBlog.ContentPath,
		CreatedAt:   time.Now(),
		LastModified: time.Now(),
	}).Error

	if err != nil {
		println("Error creating blog:", err)
		return err
	}

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

func (b *BlogService) GetUploadLink(db *gorm.DB, req types.UploadLinkRequest) (*types.UploadLinkResponse, error) {	
	key := fmt.Sprintf("blogs:%s:%s", req.AuthorID, req.BlogId)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-esat-2"),
	})

	if err != nil {
		return nil, err
	}

	svc := s3.New(sess);

	s3Req, err := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("jackdodev-webpage-posts"),
		Key:    aws.String(key),
		Body:	  strings.NewReader("This is the content of the object"),
	})

	if err != nil {
		println(err)

	str, err := s3Req.Presign(5 * time.Minute)

	if err != nil {
		println(err)
	}

	return &types.UploadLinkResponse{
		UploadURL: str,
		Key:       key,
	}, err
}
