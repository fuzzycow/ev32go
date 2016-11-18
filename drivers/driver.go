package drivers

type AvpDevice interface {
	AvpGetSetter
	OpenCloser
	Tracer
	Ider
}

type Ider interface {
	Id() string

}

type AvpGetSetter interface {
	GetAttrString(string) string
	GetAttrStringArray(string) []string
	GetAttrInt(string) int
	SetAttrString(string,string)
	SetAttrInt(string, int)
}

type OpenCloser interface {
	Open() error
	Close()
	Err() error
}

type PortFinder interface {
	Port() string
	SetPort(string)
	SetPathFilter(...string)
}

type Tracer interface {
	SetTracing(bool)
}

