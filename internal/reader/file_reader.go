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
