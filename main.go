package main

import (
	"gofeeds/scan"
	"gofeeds/validate"
	"gofeeds/handle"
	"io"
	"os"
	"log"
)

func main() {
	fileSystem := os.DirFS("samples/")
	filePaths := scan.Walk(fileSystem)
	
	conditions := []validate.Condition{
		validate.StartsWith{ Prefix: "hello-" },
	}

	validFilePaths, rejectedFilePaths := validate.ValidateFiles(filePaths, conditions)
	
	handler := handle.Handler{
		FileSystem: fileSystem, 
		FileFactory: OsFileFactory{}, 
		LetPath: "samples/accept", 
		RejectPath: "samples/reject",
	}

	for rejectedFile := range rejectedFilePaths {
		err := handler.Reject(rejectedFile)
		if err != nil {
			panic(err)
		}
	}

	for _, file := range validFilePaths {
		err := handler.Let(file)
		if err != nil {
			panic(err)
		}
	}

	log.Printf("total valid: %d", len(validFilePaths))
	log.Printf("total invalid: %d", len(rejectedFilePaths))
}

type OsFileFactory struct {
}

func (f OsFileFactory) Create(name string) (io.Writer, error) {
	return os.Create(name)
}
