package services

import (
	"time"

	"gorm.io/gorm"

	"crypto/sha256"
	"encoding/hex"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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

func (b *BlogService) CreateNewBlog(db *gorm.DB, newBlog types.BlogItem, bId string) error {
	err := db.Create(&types.Blog{
		BlogID:       bId,
		AuthorID:     newBlog.AuthorID,
		Subject:      newBlog.Subject,
		CreatedAt:    time.Now(),
		LastModified: time.Now(),
		Tags:         newBlog.Tags,
	}).Error

	if err != nil {
		println("Error creating blog:", err)
		return err
	}

	return nil
}

func (b *BlogService) QueryBlog(db *gorm.DB, blogId string) (*types.BlogItem, error) {
	blog := types.Blog{}

	println("Querying blog with ID:", blogId)

	result := db.First(&blog, "blog_id = ?", blogId)
	if result.Error != nil {
		println("Error querying blog:", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	blogItem := &types.BlogItem{
		AuthorID:  blog.AuthorID,
		Subject:   blog.Subject,
		CreatedAt: blog.CreatedAt,
		Tags:      blog.Tags,
	}

	return blogItem, nil
}

func (b *BlogService) QueryAllBlogs(db *gorm.DB) ([]types.Blog, error) {
	var blogs []types.Blog
	result := db.Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}

	return blogs, nil
}

func (b *BlogService) GetDownloadLink(db *gorm.DB, req types.DownloadLinkRequest) (*types.DownloadLinkResponse, error) {
	key := req.Key
	creds := credentials.NewSharedCredentials("/app/.aws/credentials", "default")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: creds,
	})

	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	getReq, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("jackdodev-webpage-posts"),
		Key:    aws.String(key),
	})

	str, err := getReq.Presign(5 * time.Minute)

	println("Generated presigned URL:", str)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return &types.DownloadLinkResponse{
		DownloadURL: str,
		Key:         key,
	}, nil
}

func (b *BlogService) GetUploadLink(db *gorm.DB, req types.UploadLinkRequest) (*types.UploadLinkResponse, error) {
	key := req.Key
	creds := credentials.NewSharedCredentials("/app/.aws/credentials", "default")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: creds,
	})

	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	putReq, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("jackdodev-webpage-posts"),
		Key:    aws.String(key),
	})

	str, err := putReq.Presign(5 * time.Minute)

	println("Generated presigned URL:", str)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return &types.UploadLinkResponse{
		UploadURL: str,
		Key:       key,
	}, nil
}
