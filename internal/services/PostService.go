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

func (b *PostService) CreateNewPost(db *gorm.DB, newBlog types.PostItem, id string) error {
	blog := types.Post{
		PostID:       "blg:" + id,
		AuthorID:     newBlog.AuthorID,
		Title:        newBlog.Title,
		CreatedAt:    time.Now(),
		LastModified: time.Now(),
		Tags:         newBlog.Tags,
	}

	err := b.dbService.doInsert(db, &blog)

	if err != nil {
		println("Error creating post:", err)
		return err
	}

	return nil
}

func (b *PostService) QueryPost(db *gorm.DB, id string) (*types.PostItem, error) {
	post := types.Post{}

	err := b.dbService.doQueryByID(db, &post, id)
	if err != nil {
		return nil, err
	}

	if post.PostID == "" {
		return nil, gorm.ErrRecordNotFound
	}

	blogItem := &types.PostItem{
		AuthorID:  post.AuthorID,
		Title:     post.Title,
		CreatedAt: post.CreatedAt,
		Tags:      post.Tags,
	}

	return blogItem, nil
}

func (b *PostService) QueryAllBlogs(db *gorm.DB) ([]types.PostItem, error) {
	var posts []types.Post
	result := db.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	

	return posts, nil
}

func (b *PostService) GetDownloadLink(db *gorm.DB, req types.DownloadLinkRequest) (*types.DownloadLinkResponse, error) {
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

func (b *PostService) GetUploadLink(db *gorm.DB, req types.UploadLinkRequest) (*types.UploadLinkResponse, error) {
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
