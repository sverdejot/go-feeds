package scan

import (
	"io/fs"
	"strings"
)

func Walk(fileSystem fs.FS) []string {
	return readDirRecursive(fileSystem, ".")
}

func getPath(parent, child string) string {
	return strings.TrimPrefix(parent+"/"+child, "./")
}

func readDirRecursive(fileSystem fs.FS, path string) []string {
	childs, _ := fs.ReadDir(fileSystem, path)

	var childPaths []string
	for _, child := range childs {
		if child.IsDir() {
			childEntries := readDirRecursive(fileSystem, getPath(path, child.Name()))
			for _, entry := range childEntries {
				childPaths = append(childPaths, entry)
			}
		} else {
			childPaths = append(childPaths, getPath(path, child.Name()))
		}
	}
	return childPaths
}
