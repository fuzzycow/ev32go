package sysfs

import (
	"os"
	"fmt"
	"strconv"
)


type FdCache []*os.File

func (c *FdCache ) pop() *os.File {
	d := (*c)[len(*c) - 1]
	(*c) = (*c)[:len(*c) - 1]
	return d
}

func (c *FdCache) removeN(n int) *os.File {
	d := (*c)[n]
	copy((*c)[n:], (*c)[n + 1:])
	(*c)[len(*c) - 1] = nil
	(*c) = (*c)[:len(*c) - 1]
	return d
}

func (c *FdCache) append(f *os.File) {
	(*c) = append((*c), f)
}

func (c *FdCache) insert(f *os.File, pos int) {
	(*c) = append((*c), f) // fixme
	copy((*c)[pos + 1:], (*c)[pos:])
	(*c)[pos] = f
}

func (c *FdCache) push(f *os.File) (*os.File,bool) {
	l := len(*c)
	switch {
	case l == 0:
		(*c) = append((*c), f)
		return nil,false
	case l < cap(*c):
		c.insert(f, 0)
		return nil,false
	case l == cap(*c):
		r := c.pop()
		c.insert(f, 0)
		return r,true
	default:
		panic("unexpected case in cache push")
	}
	return nil,false
}

func (c *FdCache) refresh(n int) {
	if n == 0 {
		return
	}
	t := (*c)[n]
	copy((*c)[1:n+1], (*c)[0:n])
	(*c)[0] = t
}

func (c FdCache) pos(name string) (int, bool) {
	for i, f := range c {
		if f.Name() == name {
			return i, true
		}
	}
	return 0, false
}

func NewCache(size int) *FdCache {
	c := FdCache(make([]*os.File, 0,size))
	return &c
}


func (c *FdCache ) Empty() bool { return len(*c) == 0 }

func (c *FdCache) Get(name string) (*os.File, bool) {
	if n, ok := c.pos(name); ok {
		f := (*c)[n]
		c.refresh(n)
		return f, true
	}
	return nil, false
}

func (c *FdCache) Remove(name string) (*os.File, bool) {
	if n, ok := c.pos(name); ok {
		r := c.removeN(n)
		return r, true
	}
	return nil, false
}

func (c *FdCache) Add(f *os.File) (*os.File, bool) {
	r,evicted := c.push(f)
	if evicted {
		return r, true
	} else {
		return nil, false
	}
}

func (c *FdCache) Pop() (*os.File) {
	return c.pop()
}

func (c FdCache) String() string {
	var t string
	for i, s := range c {
		t += strconv.Itoa(i) + ":  " + s.Name() + " | "
	}
	return fmt.Sprintf("len %d, cap %d: %s, raw %#v", len(c), cap(c), t,c)

}
