package sysfs

import (
	"io/ioutil"
	"path/filepath"
	"bytes"
)

type DirectDir struct {
	dirname string
}


func NewDirectDir(path string) *DirectDir {
	return &DirectDir{
		dirname: path,
	}
}

func (dir *DirectDir) ReadFile(filename string) ([]byte, error) {
	fullname := filepath.Join(dir.dirname, filename)
	b,err := ioutil.ReadFile(fullname)
	if err != nil {
		return nil,err
	}
	return bytes.TrimRight(b, "\n"), nil
}

func (dir *DirectDir) WriteFile(filename string, b []byte) error {
	fullname := filepath.Join(dir.dirname, filename)
	return ioutil.WriteFile(fullname,b,0644)
}

func (dir *DirectDir) Close() {
	// noop
}
