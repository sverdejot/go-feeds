package handle

import (
	"testing"
	"testing/fstest"
	"io"
	"path/filepath"
	"fmt"
	"slices"
)

type StubFile struct {
	Data	[]byte
}

func (s *StubFile) Write(p []byte) (int, error) {
	s.Data = p
	return len(p), nil
}

type StubFileFactory struct {
	createdFiles	map[string]*StubFile
}

func (s StubFileFactory) Create(name string) (io.Writer, error) {
	stub := &StubFile{}
	s.createdFiles[name] = stub
	return stub, nil
}

func TestHandlerReject(t *testing.T) {
	cases := []struct{
		Test			string
		Filepaths []string
	}{
		{
			"all files in root directory",
			[]string{ "rejected-file-1", "rejected-file-2", "rejected-file-3" },
		},
		{
			"nested paths should not copy, just filenames",
			[]string{ "dir1/rejected-file-1", "dir2/rejected-file-2", "dir3/dir4/rejected-file-4" },
		},
		{
			"mixed root files and nested",
			[]string { "dir1/dir2/dir3/rejected-file-1", "rejected-file-2"},
		},
	}
	for _, test := range cases {
		t.Run(test.Test, func(t *testing.T) {
			stubFileFactory := StubFileFactory{make(map[string]*StubFile, 0)}	
			
			sourceFileSystem := make(fstest.MapFS, 0)
			for _, filepath := range test.Filepaths {
				sourceFileSystem[filepath] = &fstest.MapFile{ Data: []byte(filepath) }
			}
			
			handler := Handler{
				FileSystem: 	sourceFileSystem,
				FileFactory: 	stubFileFactory,
				LetPath:			"accept",
				RejectPath:		"reject",
			}

			for _, filepath := range test.Filepaths {
				handler.Reject(filepath)
			}

			assertSameFiles(t, sourceFileSystem, handler.RejectPath, stubFileFactory.createdFiles)
		})
	}
}

func TestHandlerAccept(t *testing.T) {
	// given
	acceptedFiles := StubFileFactory{make(map[string]*StubFile,0)}
	fileName 		:= "accepted-file"
	fileContent := "accepted"
	sourceFileSystem := fstest.MapFS{
		fileName: { Data: []byte(fileContent) },
	}
	
	// when
	Handler{sourceFileSystem, acceptedFiles, "accept", "reject"}.Let(fileName)
	
	// then
	file, found := acceptedFiles.createdFiles["accept/" + fileName]
	if !found || string(file.Data) != fileContent {
		t.Errorf("not equal, got %v and want %v", file, fileContent)
	}
}


func assertSameFiles(t testing.TB, sourceFileSystem fstest.MapFS, targetPath string, targetFilesMap map[string]*StubFile) {
	t.Helper()

	for srcFile := range sourceFileSystem {
		tgtFile, found := targetFilesMap[fmt.Sprintf("%s/%s", targetPath, filepath.Base(srcFile))]
		if !found || !slices.Equal(tgtFile.Data, sourceFileSystem[srcFile].Data) {
			t.Errorf("file %s should be rejected with same content to %s", srcFile, tgtFile)
		}
	}
}
