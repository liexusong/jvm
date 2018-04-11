package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompsiteEntry {

	baseDir := path[:len(path)-1] // Remove "*"

	entries := CompsiteEntry{}

	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") {
			jarEntry := newZipEntry(path)
			entries = append(entries, jarEntry)
		}

		return nil
	})

	return entries
}
