// LegoPort (Lego Port):
// The `lego-port` class provides an interface for working with input and
// output ports that are compatible with LEGO MINDSTORMS RCX/NXT/EV3, LEGO
// WeDo and LEGO Power Functions sensors and motors. Supported devices include
// the LEGO MINDSTORMS EV3 Intelligent Brick, the LEGO WeDo USB hub and
// various sensor multiplexers from 3rd party manufacturers.
//
// Some types of ports may have multiple modes of operation. For example, the
// input ports on the EV3 brick can communicate with sensors using UART, I2C
// or analog validate signals - but not all at the same time. Therefore there
// are multiple modes available to connect to the different types of sensors.
//
// In most cases, ports are able to automatically detect what type of sensor
// or motor is connected. In some cases though, this must be manually specified
// using the `mode` and `set_device` attributes. The `mode` attribute affects
// how the port communicates with the connected device. For example the input
// ports on the EV3 brick can communicate using UART, I2C or analog voltages,
// but not all at the same time, so the mode must be set to the one that is
// appropriate for the connected sensor. The `set_device` attribute is used to
// specify the exact type of sensor that is connected. Note: the mode must be
// correctly set before setting the sensor type.
//
// Ports can be found at `/sys/class/lego-port/port<N>` where `<N>` is
// incremented each time a new port is registered. Note: The number is not
// related to the actual port at all - use the `port_name` attribute to find
// a specific port.
// ev3dev documentation:
// sysfs naming convention: lego_port/
package device

import "github.com/fuzzycow/ev32go/ev3api"

type LegoPort struct{ ev3api.Device }

func (_ LegoPort) SystemClassName() string            { return "lego_port" }
func (_ LegoPort) SystemDeviceNameConvention() string { return "" }

var LegoPortPropertyNames = []string{

	"driver_name",
	"modes",
	"mode",
	"port_name",
	"set_device",
	"status",
}

func (legoPort LegoPort) PropertyNames() []string {

	return LegoPortPropertyNames

}

// Returns the name of the driver that loaded this device. You can find the
// complete list of drivers in the [list of port drivers].
// sysfs filename: driver_name
func (legoPort *LegoPort) DriverName() string {
	return legoPort.GetAttrString("driver_name")
}

// "static method": returns a "constant", does not alter device state
// func (_ LegoPort) Property_DriverName() string { return "driver_name" }

// Returns a list of the available modes of the port.
// sysfs filename: modes
func (legoPort *LegoPort) Modes() []string {
	return legoPort.GetAttrStringArray("modes")
}

// "static method": returns a "constant", does not alter device state
// func (_ LegoPort) Property_Modes() string { return "modes" }

// Reading returns the currently selected mode. Writing sets the mode.
// Generally speaking when the mode changes any sensor or motor devices
// associated with the port will be removed new ones loaded, however this
// this will depend on the individual driver implementing this class.
// sysfs filename: mode
func (legoPort *LegoPort) Mode() string {
	return legoPort.GetAttrString("mode")
}

// Reading returns the currently selected mode. Writing sets the mode.
// Generally speaking when the mode changes any sensor or motor devices
// associated with the port will be removed new ones loaded, however this
// this will depend on the individual driver implementing this class.
// sysfs filename: mode
func (legoPort *LegoPort) SetMode(value string) {
	legoPort.SetAttrString("mode", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ LegoPort) Property_Mode() string { return "mode" }

// Returns the name of the port. See individual driver documentation for
// the name that will be returned.
// sysfs filename: port_name
func (legoPort *LegoPort) PortName() string {
	return legoPort.GetAttrString("port_name")
}

// "static method": returns a "constant", does not alter device state
// func (_ LegoPort) Property_PortName() string { return "port_name" }

// For modes that support it, writing the name of a driver will cause a new
// device to be registered for that driver and attached to this port. For
// example, since NXT/Analog sensors cannot be auto-detected, you must use
// this attribute to load the correct driver. Returns -EOPNOTSUPP if setting a
// device is not supported.
// sysfs filename: set_device
func (legoPort *LegoPort) SetSetDevice(value string) {
	legoPort.SetAttrString("set_device", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ LegoPort) Property_SetDevice() string { return "set_device" }

// In most cases, reading status will return the same value as `mode`. In
// cases where there is an `auto` mode additional values may be returned,
// such as `no-device` or `error`. See individual port driver documentation
// for the full list of possible values.
// sysfs filename: status
func (legoPort *LegoPort) Status() string {
	return legoPort.GetAttrString("status")
}

// "static method": returns a "constant", does not alter device state
// func (_ LegoPort) Property_Status() string { return "status" }
