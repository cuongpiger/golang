package service

import (
	"bytes"
	"os"
	"path/filepath"
)

type MemoryFile struct {
	buffer *bytes.Buffer
}

func NewMemoryFile() *MemoryFile {
	return &MemoryFile{
		buffer: &bytes.Buffer{},
	}
}

func (s *MemoryFile) Write(chunk []byte) error {
	if s.buffer == nil {
		return nil
	}
	_, err := s.buffer.Write(chunk)
	return err
}

func (s *MemoryFile) WriteTo(filePath string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	// Create and write file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = s.buffer.WriteTo(file)
	return err
}

func (s *MemoryFile) Size() int {
	return s.buffer.Len()
}

func (s *MemoryFile) Reset() {
	s.buffer.Reset()
}

func (s *MemoryFile) DataAsBytes() []byte {
	return s.buffer.Bytes()
}

func (s *MemoryFile) Close() error {
	// No resources to close for in-memory file
	return nil
}
