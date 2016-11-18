package sysfs

import (
	"io/ioutil"
	"strings"
	"path/filepath"
)

func readFile(parts ...string) (string, error) {
	buf, err := ioutil.ReadFile(filepath.Join(parts...))
	if err != nil {
		return "", err
	}
	s := strings.TrimSuffix(string(buf), "\n")
	return s, nil
}

