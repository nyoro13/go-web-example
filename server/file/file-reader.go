package file

import (
	"io/ioutil"
	"sort"
	"sync"
	"time"
)

type FileReadingInfo struct {
	lastRead time.Time
	path     string
	body     []byte
}

func newFileReadingInfo(filePath string, body []byte) *FileReadingInfo {
	return &FileReadingInfo{
		lastRead: time.Now(),
		path:     filePath,
		body:     body,
	}
}

func (info *FileReadingInfo) touch() {
	info.lastRead = time.Now()
}

func (info *FileReadingInfo) getSize() int {
	return len(info.body)
}

type FileReader struct {
	//cache map[string][]byte
	cache sync.Map
	mutex sync.Mutex
	limit uint
}

func NewFileReader(limit uint) *FileReader {
	return &FileReader{
		cache: sync.Map{},
		mutex: sync.Mutex{},
		limit: limit,
	}
}

func MakeFileReader(limit uint) FileReader {
	return *NewFileReader(limit)
}

func (reader *FileReader) Read(filePath string) ([]byte, error) {
	info, ok := reader.getCache(filePath)
	if !ok {
		body, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		info = reader.putCache(filePath, body)
	}

	return info.body, nil
}

func (reader *FileReader) getCache(filePath string) (*FileReadingInfo, bool) {
	raw, ok := reader.cache.Load(filePath)
	if ok {
		return raw.(*FileReadingInfo), true
	} else {
		return nil, false
	}
}

func (reader *FileReader) putCache(filePath string, body []byte) *FileReadingInfo {
	info := newFileReadingInfo(filePath, body)
	reader.cache.Store(filePath, info)

	reader.gcCache()

	return info
}

func (reader *FileReader) gcCache() {
	reader.mutex.Lock()
	defer reader.mutex.Unlock()

	if reader.limit == 0 || reader.limit >= reader.getCacheSize() {
		return
	}

	reader.gcCacheNoLock()
}

func (reader *FileReader) gcCacheNoLock() {
	if reader.limit == 0 || reader.limit >= reader.getCacheSize() {
		return
	}

	reader.removeCache(reader.countCache()/2 + 1)
	reader.gcCacheNoLock()
}

func (reader *FileReader) removeCache(count int) {
	if count <= 0 {
		return
	}

	cacheCount := reader.countCache()
	if count >= cacheCount {
		reader.cache = sync.Map{}
		return
	}

	files := make([]*FileReadingInfo, cacheCount)
	idx := 0
	reader.cache.Range(func(_, value interface{}) bool {
		files[idx] = value.(*FileReadingInfo)
		idx++
		return true
	})

	sort.Slice(files, func(i, j int) bool {
		return files[i].lastRead.Before(files[j].lastRead)
	})

	for i := 0; i < count; i++ {
		reader.cache.Delete(files[i].path)
	}
}

func (reader *FileReader) countCache() int {
	count := 0
	reader.cache.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}

func (reader *FileReader) getCacheSize() uint {
	var size uint
	reader.cache.Range(func(_, v interface{}) bool {
		size += uint(v.(*FileReadingInfo).getSize())
		return true
	})
	return size
}
