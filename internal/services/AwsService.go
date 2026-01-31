package services

import (
	types "go_webserv/internal/types"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)
type AwsService struct {
}

func InitAwsService() *AwsService {
	return &AwsService{}
}

func (a *AwsService) GetDownloadLink(req types.DownloadLinkRequest) (*types.DownloadLinkResponse, error) {
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

	if err != nil {
		return nil, err
	}

	return &types.DownloadLinkResponse{
		DownloadURL: str,
		Key:         key,
	}, nil
}

func (a *AwsService) GetUploadLink(req types.UploadLinkRequest) (*types.UploadLinkResponse, error) {
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

	if err != nil {
		return nil, err
	}

	return &types.UploadLinkResponse{
		UploadURL: str,
		Key:       key,
	}, nil
}
