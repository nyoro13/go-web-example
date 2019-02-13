package file

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

type FileReader struct {
	cache map[string][]byte
}

func NewFileReader() *FileReader {
	return &FileReader{cache: make(map[string][]byte)}
}

func MakeFileReader() FileReader {
	return FileReader{cache: make(map[string][]byte)}
}

func (cacher *FileReader) Read(path string) ([]byte, error) {
	data, err := cacher.getCache(path)
	if err == nil {
		return data, nil
	}

	return cacher.readFromFile(path)
}

func (cacher *FileReader) Clear() {
	// for now
	cacher.cache = make(map[string][]byte)
}

func (cacher *FileReader) Delete(path string) {
	delete(cacher.cache, makePathHash(path))
}

func (cacher *FileReader) readFromFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cacher.putCache(path, &data)
	return data, nil
}

func (cacher *FileReader) getCache(path string) ([]byte, error) {
	if cacher.cache == nil {
		cacher.cache = make(map[string][]byte)
	}
	data, ok := cacher.cache[makePathHash(path)]
	if !ok {
		return nil, fmt.Errorf("Not Found Cache. %s", path)
	}

	return data, nil
}

func (cacher *FileReader) putCache(path string, data *[]byte) {
	hash := makePathHash(path)
	if cacher.cache == nil {
		cacher.cache = make(map[string][]byte)
	}

	cacher.cache[hash] = *data
}

func makePathHash(path string) string {
	hash := sha256.Sum256([]byte(path))
	return hex.EncodeToString(hash[:])
}
