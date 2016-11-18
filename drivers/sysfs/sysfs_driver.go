package sysfs

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"log"
	"io/ioutil"
	"fmt"
	"path/filepath"
)



const PORT_NAME_FILE = "port_name"

var DefaultSysfsClassDir = "/sys/class"

type SysfsDriver struct {
	PathFilter []string
	port string
	id string
	caching bool
	sync.Mutex
	dir Dir
	err error
	tracing bool
}

func NewDriver() *SysfsDriver {
	return &SysfsDriver{caching: true}
}

func NewNonCachingDriver() *SysfsDriver {
	return &SysfsDriver{caching: false}
}

// FIXME - make work for all devices !!!
func (drv *SysfsDriver) Id() string {
	return "port=" + drv.port
}

func (drv *SysfsDriver) SetPort(port string) {
	drv.port = port
}

func (drv *SysfsDriver) SetPathFilter(path ...string) {
	drv.PathFilter = path
}

func (drv *SysfsDriver) SetTracing(enabled bool) {
	drv.Lock()
	drv.tracing = enabled
	drv.Unlock()
}

func (drv *SysfsDriver) Open() error {
	drv.Lock()
	defer drv.Unlock()
	devPath,err  := drv.findPath()
	if err != nil {
		return err
	}
	var dir Dir
	if drv.caching {
		dir = NewCachedDir(devPath)
	} else {
		dir = NewDirectDir(devPath)
	}
	drv.dir = dir
	return nil
}

func (drv *SysfsDriver) Close(){
	drv.Lock()
	defer drv.Unlock()
	dir := drv.dir
	if dir != nil {
		dir.Close()
	}
}

func (drv *SysfsDriver) Err() error {
	var err error
	drv.Lock()
	err = drv.err
	drv.Unlock()
	return err
}

func (drv *SysfsDriver) setErr(err error) error {
	if err != nil {
		drv.err = err
	}
	return err
}


func (drv *SysfsDriver) findPath() (string, error) {
	pat := make([]string,len(drv.PathFilter))
	for i,p := range drv.PathFilter {
		pat[i] = strings.Replace(p,"{0}","*",1)
	}
	pattern := filepath.Join(DefaultSysfsClassDir, filepath.Join(pat...))
	if drv.port != "" {
		pattern = filepath.Join(pattern,PORT_NAME_FILE)
	}

	matches, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalf("Failed to glob file pattern '%s': %v", pattern,err)
	}

	if len(matches) == 0 {
		return "", errors.New("port path not found using expression: " + pattern)
	}

	var drvPath string

	if drv.port != "" {
		MATCHING: for _, file := range matches {
			thisPort, err := readFile(file)
			if err == nil {
				if thisPort == drv.port {
					drvPath = filepath.Dir(file)
					break MATCHING
				}
			}
		}
	} else {
		drvPath = matches[0]
	}

	if drvPath == "" {
		return "",errors.New("port path not found (2)")
	}

	if _, err = ioutil.ReadDir(drvPath); err != nil {
		log.Fatalf("failed to read port dir %v", drvPath)
		return "", errors.New("port not found (3)")
	}

	return drvPath, nil
}


func (drv *SysfsDriver) String() string {
	return fmt.Sprintf("SysfsDriver port=%s,pathfilter=%+v",drv.port,drv.PathFilter)
}

func (drv *SysfsDriver) readString(attr string) (string, error) {
	if drv.dir == nil {
		return "", errors.New("sysfs dir is nil")
	}
	b, err := drv.dir.ReadFile(attr)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (drv *SysfsDriver) readInt(attr string) (int, error) {
	if drv.dir == nil {
		return 0, errors.New("sysfs dir is nil")
	}
	b, err := drv.dir.ReadFile(attr)
	if err != nil {
		return 0, err
	}
	if len(b) == 0 {
		return 0, errors.New("can not convert empty string to int")
	}
	i, err := strconv.Atoi(string(b));
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (drv *SysfsDriver) writeString(attr, s string) error {
	if drv.dir == nil {
		return errors.New("sysfs dir is nil")
	}
	return drv.dir.WriteFile(attr, []byte(s))
}

func (drv *SysfsDriver) writeInt(attr string, i int) error {
	if drv.dir == nil {
		return errors.New("sysfs dir is nil")
	}
	return drv.dir.WriteFile(attr, []byte(strconv.Itoa(i)))
}


func (drv *SysfsDriver) GetAttrString(attr string) string{
	drv.Lock()
	value, err := drv.readString(attr)
	drv.setErr(err)
	if drv.tracing {
		log.Println("TRACE: GetAttrInt(" + attr + ")=" + value + "\terr=",err)
	}
	drv.Unlock()
	return value
}

func (drv *SysfsDriver) SetAttrString(attr string, value string){
	drv.Lock()
	err :=  drv.writeString(attr,value)
	drv.setErr(err)
	if drv.tracing {
		log.Println("TRACE: SetAttrString(" + attr + ")=" + value + ")\terr=",err)
	}
	drv.Unlock()
}

func (drv *SysfsDriver) GetAttrInt(attr string) int {
	drv.Lock()
	value, err := drv.readInt(attr)
	drv.setErr(err)
	if drv.tracing {
		log.Println("TRACE: GetAttrInt(" + attr + ")=" + strconv.Itoa(value) + ")\terr=",err)
	}
	drv.Unlock()
	return value
}

func (drv *SysfsDriver) SetAttrInt(attr string, value int)  {
	drv.Lock()
	err := drv.writeInt(attr,value)
	drv.setErr(err)
	if drv.tracing {
		log.Println("TRACE: SetAttrInt(" + attr + "," + strconv.Itoa(value) + ")\terr=",err)
	}
	drv.Unlock()
}

func (drv *SysfsDriver) GetAttrStringArray(attr string) []string {
	val := drv.GetAttrString(attr)
	if drv.Err() != nil || len(val) == 0 {
		return []string{}
	}
	list := strings.Split(val, " ")
	return list
}

func (drv *SysfsDriver) MGetAttrString(attrs []string) map[string]string {
	m := make(map[string]string)
	for _,a := range attrs {
		value := drv.GetAttrString(a)
		if drv.Err() != nil {
			break
		} else {
			m[a] = value
		}
	}
	return m
}

func (drv *SysfsDriver) MSetAttrString(avps map[string]string) {
	for a,value := range avps {
		drv.SetAttrString(a,value)
		if drv.Err() != nil {
			break
		}
	}
}

