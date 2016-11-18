package ev3api

import (
	"github.com/fuzzycow/ev32go/drivers"
)

type Device interface {
	drivers.AvpDevice
}

/*
type Opener interface {
	Open() error
	Close()
	SetSelector(*DeviceSelector)
	Selector() *DeviceSelector
}


type PropertyNamer interface {
	PropertyNames() []string
	SetPropertyNames([]string)
	AddPropertyNames([]string)
}
*/