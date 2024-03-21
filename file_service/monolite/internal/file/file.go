package file

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
)

type File struct {
	FilePath   string
	buffer     *bytes.Buffer
	OutputFile *os.File
}

func NewFile() *File {
	return &File{
		FilePath:   "",
		buffer:     &bytes.Buffer{},
		OutputFile: nil,
	}
}

func (f *File) SetFile(filename, path string, logger *slog.Logger) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		logger.Error(err.Error())
		return err
	}

	f.FilePath = filepath.Join(path, filename)

	file, err := os.Create(f.FilePath)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	f.OutputFile = file
	return nil
}

func (f *File) Write(chunk []byte) error {
	if f.OutputFile == nil {
		return nil
	}

	_, err := f.OutputFile.Write(chunk)
	return err
}

func (f *File) Close() error {
	return f.OutputFile.Close()
}
