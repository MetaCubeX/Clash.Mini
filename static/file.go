package static

import (
	"io/fs"
	"time"
)

type FakeFile struct {
	name string
	isDir bool
	data *[]byte
}

func (f FakeFile) Name() string {
	return f.name
}

func (f FakeFile) Size() int64 {
	if f.data == nil {
		return -1
	}
	return int64(len(*(f.data)))
}

func (f FakeFile) Mode() fs.FileMode {
	return fs.ModeIrregular
}

func (f FakeFile) ModTime() time.Time {
	return time.Now()
}

func (f FakeFile) IsDir() bool {
	return f.isDir
}

func (f FakeFile) Sys() interface{} {
	return f.data
}
