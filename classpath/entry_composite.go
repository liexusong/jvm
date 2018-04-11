package classpath

import (
	"errors"
	"strings"
)

type CompsiteEntry []Entry

func newCompsiteEntry(pathList string) CompsiteEntry {

	entries := CompsiteEntry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		entries = append(entries, entry)
	}

	return entries
}

func (self CompsiteEntry) readClass(className string) ([]byte, Entry, error) {

	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (self CompsiteEntry) String() string {

	strs := make([]string, len(self))

	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
