package reader

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3FileReader struct {
	client *s3.Client
	bucket string
}

type S3FileReaderConfig struct {
	Region string
	Bucket string
}

func NewS3FileReader(cfg aws.Config, bucket string) *S3FileReader {
	return &S3FileReader{
		client: s3.NewFromConfig(cfg),
		bucket: bucket,
	}
}

func (s *S3FileReader) ReadFile(filename string) (*os.File, error) {

	// Create a temp file to store the downloaded file
	outputFile, err := os.CreateTemp("", "tmp-*")
	if err != nil {
		return nil, err
	}

	log.Println("Downloading file from S3", outputFile.Name())

	localFile, err := s.downloadFile(context.Background(), filename, outputFile)
	if err != nil {
		return nil, err
	}

	return os.Open(localFile)
}

func (s *S3FileReader) downloadFile(ctx context.Context, key string, outputFile *os.File) (string, error) {

	downloader := manager.NewDownloader(s.client)
	defer outputFile.Close()

	_, err := downloader.Download(ctx, outputFile, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}

	return outputFile.Name(), nil

}
