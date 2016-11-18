package influxtl

import (
	"bytes"
	"strconv"
	"fmt"
	"strings"
	"time"
	"github.com/fuzzycow/ev32go/robotics/telemetry"
	"log"
)

type Message struct {
	tlc telemetry.Client
	buf bytes.Buffer
	sealed bool
	fields int
}

func newClientMessage(tlc telemetry.Client, subject ...string)  telemetry.Message {
	m := &Message{tlc: tlc}
	m.buf.WriteString(strings.Join(subject,","))
	m.buf.WriteByte(' ')
	return m
}

func (m *Message) String() string {
	return string(m.buf.Bytes())
}

func (m *Message) Reset() telemetry.Message {
	m.buf.Reset()
	m.sealed = false
	m.fields = 0
	return m
}

func (m *Message) Add(attr string, value interface{}) telemetry.Message {
	switch value := value.(type) {
	case string:
		return m.AddString(attr,value)
	case int:
		return m.AddInt(attr,value)
	case int64:
		return m.AddInt64(attr, value)
	case float64:
		return m.AddFloat64(attr,value)
	default:
		return m.AddString(attr,fmt.Sprintf("%v",value))
	}
	return m
}

func (m *Message) AddString(attr string, value string)  telemetry.Message  {
	if m.sealed {
		return m
	}
	if m.fields > 0 {
		m.buf.WriteByte(',')
	}
	m.buf.WriteString(attr + "=" + value)
	m.fields += 1
	return m
}

func (m *Message) AddInt(attr string, i int)  telemetry.Message  {
	return m.AddString(attr, strconv.Itoa(i) + "i")
}

func (m *Message) AddFloat64(attr string, f float64)  telemetry.Message  {
	return m.AddString(attr,strconv.FormatFloat(f, 'f', 4, 64))
}

func (m *Message) AddInt64(attr string, i int64)  telemetry.Message  {
	return m.AddString(attr,strconv.FormatInt(i,10) + "i")
}

func (m *Message) Send() error {
	return m.SendTo(m.tlc)
}

func (m *Message) SendTo(tlc telemetry.Client) error {
	if tlc == nil {
		//return fmt.Errorf("Client not defined")
		log.Fatalf("attempt to send telemetry message to nil client")
	}
	m.seal()

	_,err :=  tlc.Write(m.Bytes())
	m.Reset()
	return err
}

func (m *Message) Bytes() []byte {
	return m.buf.Bytes()
}

func (m *Message) seal() {
	if ! m.sealed {
		m.buf.WriteByte('\n')
	}
}

func (m *Message) Timestamp()  telemetry.Message {
	if ! m.sealed {
		m.buf.WriteByte(' ')
		m.buf.WriteString(strconv.FormatInt(time.Now().UnixNano(),10))
		m.seal()
	}
	return m
}


