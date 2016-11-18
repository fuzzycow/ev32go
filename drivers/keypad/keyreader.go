package keypad

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"errors"
	"flag"
	"log"
	"os"
	"syscall"
	"time"
	"unsafe"
)

/*
Credits to:
	github.com/goburrow/modbus
	github.com/gvalkov/golang-evdev

*/

const INPUT_DEVICE = "/dev/input/event0"

type KeyCode uint16
type KeyValue uint32

const (
	KEY_UP    KeyCode = 103
	KEY_DOWN          = 108
	KEY_LEFT          = 105
	KEY_RIGHT         = 106
	KEY_ENTER         = 28
	KEY_ESC           = 14
	KEY_MAX           = 0x2ff

	KEY_PRESS   KeyValue = 1
	KEY_RELEASE          = 0
)

type InputEvent struct {
	Time syscall.Timeval // time in seconds since epoch at which event occurred
	//Sec  uint32
	//Usec uint32
	Type  uint16   // event type - one of ecodes.EV_*
	Code  KeyCode  // event code related to the event type
	Value KeyValue // event value related to the event type

}

var ErrReadTimeout = errors.New("read timeout")
var eventsize = int(unsafe.Sizeof(InputEvent{}))
var ReadTimeout = 500 * time.Millisecond

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func (ev *InputEvent) String() string {
	return fmt.Sprintf("event at %d.%d, code %02d, type %02d, val %02d",
		ev.Time.Sec, ev.Time.Usec, ev.Code, ev.Type, ev.Value)
}

// fdGet returns index and offset of fd in fds
func fdGet(fd int, fds *syscall.FdSet) (index, offset int) {
	index = fd / (syscall.FD_SETSIZE / len(fds.Bits)) % len(fds.Bits)
	offset = fd % (syscall.FD_SETSIZE / len(fds.Bits))
	return
}

// fdIsSet implements FD_ISSET macro
func fdIsSet(fd int, fds *syscall.FdSet) bool {
	idx, pos := fdGet(fd, fds)
	return fds.Bits[idx]&(1<<uint(pos)) != 0
}

// fdSet implements FD_SET macro
func fdSet(fd int, fds *syscall.FdSet) {
	idx, pos := fdGet(fd, fds)
	fds.Bits[idx] = 1 << uint(pos)
}

// Read reads from serial port, blocked until data received or timeout after Timeout.
func ReadWithTimeout(b []byte, file *os.File, nsec int64) (int, error) {
	var timeout *syscall.Timeval
	var tv syscall.Timeval

	timeout = nil
	if nsec > 0 {
		tv = syscall.NsecToTimeval(nsec)
		timeout = &tv
	}

	var n int
	var err error

	var rfds syscall.FdSet
	fd := int(file.Fd())
	fdSet(fd, &rfds)
	//te.Sec = int32(timeout.Nanoseconds() / 1E9)
	//te.Usec = int32((timeout.Nanoseconds() % 1E9) / 1E3)

	if _, err = syscall.Select(fd+1, &rfds, nil, nil, timeout); err != nil {
		log.Printf("select error")
		return 0, err
	}

	if fdIsSet(fd, &rfds) {
		n, err = file.Read(b)
		return n, err
	} else {
		return 0, ErrReadTimeout
	}
}

// Read and return a single input event.
func ReadInputEvent(file *os.File, timeout time.Duration) (*InputEvent, error) {
	event := InputEvent{}

	//log.Printf("event size: %d", eventsize)
	buffer := make([]byte, eventsize)

	n, err := ReadWithTimeout(buffer, file, timeout.Nanoseconds())
	if err != nil {
		return nil, err
	} else if n < eventsize {
		log.Fatalf("read only %d from kbd, expecting %d", n, eventsize)
	}
	b := bytes.NewBuffer(buffer)
	err = binary.Read(b, binary.LittleEndian, &event)
	if err != nil {
		return nil, err
	}

	//event.Type = binary.LittleEndian.Uint16(buffer[8:10])
	//event.Code = binary.LittleEndian.Uint16(buffer[10:12])
	//event.Value = binary.LittleEndian.Uint32(buffer[12:16])
	log.Printf("FIXME: %#v", event)
	return &event, err
}

func FollowEvents(quit <-chan struct{}) {
	file, err := os.Open(INPUT_DEVICE)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open keypad at %f:", err)
	}
	log.Printf("opened keypad at %s", file.Name())
LOOP:
	for {
		select {
		case _, ok := <-quit:
			if !ok {
				break LOOP
			}
		default:
			ev, err := ReadInputEvent(file, ReadTimeout)
			if err == nil {
				log.Printf("GOT: %s", ev)
			} else if err == ErrReadTimeout {
				continue
			} else {
				log.Fatalf("ERROR: ", err)
			}
		}
	}
	log.Printf("stopped")
}

func main() {
	q := make(chan struct{})
	defer close(q)
	go FollowEvents(q)
	for i := 0; i < 15; i++ {
		time.Sleep(time.Second)
		log.Info("...")
	}
}
