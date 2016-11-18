package sysfs

import (
	"io/ioutil"
	"testing"
	"os"
	"path/filepath"
	"sync/atomic"
	"path"
	"strconv"
	"fmt"
	"bytes"
)

var TEMP_DIR = "/tmp/"
var TEMP_PREFIX_1 = "ev32go-test-"

type logErrorf func(string,...interface{})

func testWriteRead(logFunc logErrorf,cdir *CachedDir, n int) {
	str := strconv.Itoa(n)
	file, err := ioutil.TempFile(cdir.path(), str );
	if err != nil {
		panic(fmt.Sprintf("failed to create temp file: %v",err))
	}
	//logFunc("using temporary file %v", file.Name())
	fileName := file.Name()
	baseName := filepath.Base(fileName)
	err = file.Close()
	if err != nil {
		logFunc("could not close temp file - bad tmp dir ?: %v",err)
	}
	defer os.Remove(fileName)
	buf1 := []byte(str)
	err = cdir.WriteFile(baseName, buf1)
	if err != nil {
		logFunc("failed to write to file: %v", err)
	}
	buf2, err := cdir.ReadFile(baseName)
	if err != nil {
		logFunc("failed to write to file: %v", err)
	}
	if ! bytes.Equal(buf1,buf2) {
		logFunc("read/write inconsistent")
	}
}

func TestRW(t *testing.T) {
	tempdir, err := ioutil.TempDir(os.TempDir(), TEMP_PREFIX_1)
	if err != nil {
		t.Errorf("failed to create temp dir: %v", err)
	}
	defer os.Remove(tempdir)
	base := "golang-test-file"
	defer os.Remove(path.Join(tempdir,base))
	t.Logf("using temporary directory %s", tempdir)
	cdir := NewCachedDir(tempdir)
	defer cdir.Close()
	for i := 0; i < 10; i ++ {
		testWriteRead(t.Errorf,cdir,i)
	}
}



func BenchmarkTmpDirRWP(b *testing.B) {
	tempdir, err := ioutil.TempDir(os.TempDir(), TEMP_PREFIX_1)
	if err != nil {
		b.Errorf("failed to create temp dir: %v", err)
	}
	defer os.Remove(tempdir)
	b.Logf("using temporary directory %s", tempdir)

	var counter uint32 = 0

	//b.SetParallelism(32)
	b.RunParallel(func(pb *testing.PB) {
		i := int(atomic.AddUint32(&counter,1))
		cdir := NewCachedDir(tempdir)
		defer cdir.Close()
		b.Logf("starting worker %d", i)
		for pb.Next() {
			testWriteRead(b.Errorf,cdir,i)
		}
	})
}
