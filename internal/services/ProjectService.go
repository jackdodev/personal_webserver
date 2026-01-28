package services

import (
	"time"

	"gorm.io/gorm"

	types "go_webserv/internal/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ProjectService struct {
}

func InitProjectService() *ProjectService {
	return &ProjectService{}
}

func (p *ProjectService) CreateNewProject(db *gorm.DB, newProject types.ProjectItem, pId string) error {
	err := db.Create(&types.Project{
		ProjectID:    pId,
		AuthorID:     newProject.AuthorID,
		ProjectName:  newProject.ProjectName,
		CreatedAt:    time.Now(),
		LastModified: time.Now(),
	}).Error

	if err != nil {
		println("Error creating project:", err)
		return err
	}

	return nil
}

func (p *ProjectService) QueryProject(db *gorm.DB, projectId string) (*types.ProjectItem, error) {
	project := types.Project{}

	println("Querying project with ID:", projectId)

	result := db.First(&project, "project_id = ?", projectId)
	if result.Error != nil {
		println("Error querying project:", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	projectItem := &types.ProjectItem{
		AuthorID:     project.AuthorID,
		ProjectName:  project.ProjectName,
		CreatedAt:    project.CreatedAt,
		LastModified: project.LastModified,
		Tags:         project.Tags,
	}

	return projectItem, nil
}

func (p *ProjectService) QueryAllProjects(db *gorm.DB) ([]types.Project, error) {
	var projects []types.Project
	result := db.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}

func (p *ProjectService) GetDownloadLink(db *gorm.DB, req types.DownloadLinkRequest) (*types.DownloadLinkResponse, error) {
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

func (p *ProjectService) GetUploadLink(db *gorm.DB, req types.UploadLinkRequest) (*types.UploadLinkResponse, error) {
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
