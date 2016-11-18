package sysfs

import (
	"log"
	"io"
	"bytes"
	"os"
	"path"
	"sync"
	"syscall"

)


const READ_BUF_SIZE = 128

var DefaultCacheSize = 16

type CachedDir struct {
	reading  *FdCache
	writing  *FdCache
	dirname string
	mu       sync.Mutex
}

func NewCachedDir(path string) *CachedDir {
	return &CachedDir{
		reading:  NewCache(DefaultCacheSize),
		writing:  NewCache(DefaultCacheSize),
		dirname: path,
	}
}

func (cdir *CachedDir) path() string {
	return cdir.dirname
}


func (cdir *CachedDir) getFile(filename string, writing bool, reopen bool) (*os.File, error) {
	var (
		cache  *FdCache
		mode int
	)
	if writing {
		cache = cdir.writing
		mode = os.O_WRONLY // | os.O_CREATE
	} else {
		cache = cdir.reading
		mode = os.O_RDONLY
	}

	fullname := path.Join(cdir.dirname, filename)

	//log.Printf("checking if file %s is already cached", fullname)
	if file, ok := cache.Get(fullname); ok {
		//log.Printf("file %v already in cache", fullname)
		if reopen {
			if f,ok := cache.Remove(fullname); ok {
				f.Close()
			}
		} else {
			return file, nil
		}
	}

	//log.Printf("opening file %s: %s (mode: %d)", filename,fullname, mode)
	file, err := os.OpenFile(fullname,mode,0666)

	if err != nil {
		log.Printf("failed to open file %s: %v", fullname, err)
		return nil, err
	}

	//log.Printf("adding file %s to cache", file.Name())
	if evfile, eviction := cache.Add(file); eviction {
		//log.Printf("evicted %s",evfile.Name())
		evfile.Close()
	}

	return file, nil
}


func (cdir *CachedDir) ReadFile(filename string) ([]byte, error) {
	b := make([]byte, READ_BUF_SIZE)
	reopen := false
	var err error
	cdir.mu.Lock()
	defer cdir.mu.Unlock()
	for i := 0; i < 2; i ++ {
		file, err := cdir.getFile(filename, false, reopen)
		if err != nil {
			return nil, err
		}

		// WARNING: May cut long lines, but x2 faster then other methods
		n, err := file.ReadAt(b, 0)
		_ = n
		if ( err == nil || err == io.EOF ) {
			return bytes.Trim(b[:n], "\n"), nil
		}

		//MAYBE_RETRY:
		if err2, ok := err.(*os.PathError); ok && err2.Err == syscall.ENODEV {
			log.Printf("file %s has gone away - retrying", filename)
			reopen = true
			continue
		}
		break
	}
	return nil, err
}


func (cdir *CachedDir) WriteFile(filename string, b []byte) error {
	var err error
	reopen := false
	cdir.mu.Lock()
	defer cdir.mu.Unlock()
	for i := 0; i < 2; i ++ {
		file, err := cdir.getFile(filename, true, reopen)
		if err != nil {
			return err
		}
		_, err = file.WriteAt(b, 0)
		if  err == nil {
			return nil
		}
		if err2, ok := err.(*os.PathError); ok && err2.Err == syscall.ENODEV {
			log.Printf("file %s has gone away - retrying", filename)
			reopen = true
			continue
		}
		break
	}
	return err
}



func (cdir *CachedDir) Close() {
	cdir.mu.Lock()
	defer cdir.mu.Unlock()

	caches := []*FdCache{cdir.reading, cdir.writing}

	for _, c := range caches {
		//c.Show()
		for ! c.Empty() {
			file :=  c.Pop()
			if file == nil {
				log.Fatalf("got null file !")
			}
			//log.Printf("closing file : %s", file.Name())
			filename := file.Name()

			err := file.Close()
			if err != nil {
				log.Fatalf("failed to close file %s", filename)
			}
		}
	}
}


