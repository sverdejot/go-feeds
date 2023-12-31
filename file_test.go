package file

import (
	"io/fs"
	"slices"
	"strings"
	"testing"
	"testing/fstest"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Case      string
		FilePaths []string
		Expected  []string
	}{
		{
			Case:				"files only on root directory",
			FilePaths:	[]string{"file1", "file2"},
			Expected:		[]string{"file1", "file2"},
		},
		{
			Case:      "files on subdirectory with no other directories",
			FilePaths:	[]string{"dir/file1", "dir/file2"},
			Expected:		[]string{"dir/file1", "dir/file2"},
		},
		{
			Case:      "3-level nested files",
			FilePaths:	[]string{"dir/subdir/file1", "dir/subdir2/file2"},
			Expected:		[]string{"dir/subdir/file1", "dir/subdir2/file2"},
		},
		{
			Case:				"nested empty directory",
			FilePaths:	[]string{ "dir/" },
			Expected:		[]string{},
		},
		{
			Case:				"two nested directories and one empty",
			FilePaths:	[]string{ "dir/child1/file", "dir/child2/", },
			Expected:		[]string{ "dir/child1/file" },
		},
		{
			Case:				"three double nested directories, with two empty",
			FilePaths:	[]string{ "dir/child1/file", "dir/child2/file", "dir/child3/", "dir/child1/subchild/" },
			Expected:		[]string{ "dir/child1/file", "dir/child2/file" },
		},
	}

	for _, test := range cases {
		t.Run(test.Case, func(t *testing.T) {
			fileSystem := createFs(test.FilePaths)

			got := Walk(fileSystem)

			if !slices.Equal(got, test.Expected) {
				t.Errorf("got %q want %q", got, test.Expected)
			}
		})
	}
}

func createFs(paths []string) fstest.MapFS {
	pathFiles := map[string]*fstest.MapFile{}

	for _, path := range paths {
		if !strings.HasSuffix(path, "/") {
			pathFiles[path] = &fstest.MapFile{Data: []byte("")}
		} else {
			pathFiles[path] = &fstest.MapFile{Mode: fs.ModeDir}
		}

	}
	return fstest.MapFS(pathFiles)
}
