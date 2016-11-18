package telemetry

type Client interface {
	Open() error
	Close() error
	NewMessage(...string) Message
	Write([]byte) (int,error)

}

type Message interface {
	String() string
	Send() error
	SendTo(Client) error
	Reset() Message
	Add(string,interface{}) Message
	AddString(string,string) Message
	AddInt(string,int) Message
	AddFloat64(string,float64) Message
	AddInt64(string,int64) Message
	Timestamp() Message
}



