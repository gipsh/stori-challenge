package reader

import "os"

type LocalFileReader struct {
}

func NewLocalFileReader() *LocalFileReader {
	return &LocalFileReader{}
}

func (l *LocalFileReader) ReadFile(filename string) (*os.File, error) {
	return os.Open(filename)
}
