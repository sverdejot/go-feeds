package handle

import (
	"io"
	"io/fs"
	"path/filepath"
	"fmt"
)

type FileFactory interface {
	Create(name string) (io.Writer, error)
}

type Handler struct {
	FileSystem		fs.FS
	FileFactory		FileFactory
	LetPath				string
	RejectPath		string
}

func (h Handler) Reject(path string) error {
	// pls, proper error handling
	targetDirectory := filepath.Base(path)
	writer, err := h.FileFactory.Create(fmt.Sprintf("%s/%s", h.RejectPath, targetDirectory))

	if err != nil {
		return err
	}

	content, err := readContent(h.FileSystem, path)
	 
	if err != nil {
		return err
	}

	_, err = writer.Write(content)

	return err
}

func (h Handler) Let(path string) error {
	// pls, proper error handling
	targetDirectory := filepath.Base(path)
	writer, err := h.FileFactory.Create(fmt.Sprintf("%s/%s", h.LetPath, targetDirectory))
	
	if err != nil {
		return err
	}

	content, err := readContent(h.FileSystem, path)
	 
	if err != nil {
		return err
	}

	_, err = writer.Write(content)

	return err
}

func readContent(fileSystem fs.FS, sourcePath string) ([]byte, error) {
	file, err := fileSystem.Open(sourcePath)
	defer file.Close()
	
	if err != nil {
		return nil, err
	}

	return io.ReadAll(file)
}
