package reader

import "os"

// FileReader is an interface that defines the ReadFile method
// to abstract the file reading process from local fs or s3
type FileReader interface {
	ReadFile(filename string) (*os.File, error)
}

type FsType string

const (
	S3    FsType = "s3"
	Local FsType = "local"
)

func NewFileReader(fsType FsType) FileReader {
	switch fsType {
	case S3:
		config := S3FileReaderConfig{
			Region: os.Getenv("AWS_REGION"),
			Bucket: os.Getenv("S3_BUCKET"),
		}
		return NewS3FileReader(config)
	case Local:
		return NewLocalFileReader()
	default:
		return nil
	}
}
