package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {

	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {

	rc, err := zip.OpenReader(self.absPath) // Open zip file
	if err != nil {
		return nil, nil, err
	}

	defer rc.Close()

	for _, file := range rc.File {

		if file.Name == className {

			result, err := file.Open()
			if err != nil {
				return nil, nil, err
			}

			defer result.Close()

			data, err := ioutil.ReadAll(result)
			if err != nil {
				return nil, nil, err
			}

			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
