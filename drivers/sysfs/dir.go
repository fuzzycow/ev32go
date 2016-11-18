package sysfs


type Dir interface {
	ReadFile(string) ([]byte, error)
	WriteFile(string,[]byte) error
	Close()
}
