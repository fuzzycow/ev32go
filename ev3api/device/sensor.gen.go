// Sensor (Sensor):
// The sensor class provides a uniform interface for using most of the
// sensors available for the EV3. The various underlying device drivers will
// create a `lego-sensor` device for interacting with the sensors.
//
// Sensors are primarily controlled by setting the `mode` and monitored by
// reading the `value<N>` attributes. Values can be converted to floating point
// if needed by `value<N>` / 10.0 ^ `decimals`.
//
// Since the name of the `sensor<N>` device node does not correspond to the port
// that a sensor is plugged in to, you must look at the `port_name` attribute if
// you need to know which port a sensor is plugged in to. However, if you don't
// have more than one sensor of each type, you can just look for a matching
// `driver_name`. Then it will not matter which port a sensor is plugged in to - your
// program will still work.
// ev3dev documentation: http://www.ev3dev.org/docs/drivers/lego-sensor-class/
// sysfs naming convention: lego-sensor/sensor{0}
package device

import "github.com/fuzzycow/ev32go/ev3api"

type Sensor struct{ ev3api.Device }

func (_ Sensor) SystemClassName() string            { return "lego-sensor" }
func (_ Sensor) SystemDeviceNameConvention() string { return "sensor{0}" }

var SensorPropertyNames = []string{

	"command",
	"commands",
	"decimals",
	"driver_name",
	"mode",
	"modes",
	"num_values",
	"port_name",
	"units",
}

func (sensor Sensor) PropertyNames() []string {

	return SensorPropertyNames

}

// Sends a command to the sensor.
// sysfs filename: command
func (sensor *Sensor) SetCommand(value string) {
	sensor.SetAttrString("command", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Sensor) Property_Command() string { return "command" }

// Returns a list of the valid commands for the sensor.
// Returns -EOPNOTSUPP if no commands are supported.
// sysfs filename: commands
func (sensor *Sensor) Commands() []string {
	return sensor.GetAttrStringArray("commands")
}

// "static method": returns a "constant", does not alter device state
// func (_ Sensor) Property_Commands() string { return "commands" }

// Returns the number of decimal places for the values in the `value<N>`
// attributes of the current mode.
// sysfs filename: decimals
func (sensor *Sensor) Decimals() int {
	return sensor.GetAttrInt("decimals")
}

// "static method": returns a "constant", does not alter device state
// func (_ Sensor) Property_Decimals() string { return "decimals" }

// Returns the name of the sensor device/driver. See the list of [supported
// sensors] for a complete list of drivers.
// sysfs filename: driver_name
func (sensor *Sensor) DriverName() string {
	return sensor.GetAttrString("driver_name")
}

// "static method": returns a "constant", does not alter device state
// func (_ Sensor) Property_DriverName() string { return "driver_name" }

// Returns the current mode. Writing one of the values returned by `modes`
// sets the sensor to that mode.
// sysfs filename: mode
func (sensor *Sensor) Mode() string {
	return sensor.GetAttrString("mode")
}

// Returns the current mode. Writing one of the values returned by `modes`
// sets the sensor to that mode.
// sysfs filename: mode
func (sensor *Sensor) SetMode(value string) {
	sensor.SetAttrString("mode", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Sensor) Property_Mode() string { return "mode" }

// Returns a list of the valid modes for the sensor.
// sysfs filename: modes
func (sensor *Sensor) Modes() []string {
	return sensor.GetAttrStringArray("modes")
}

// "static method": returns a "constant", does not alter device state
// func (_ Sensor) Property_Modes() string { return "modes" }

// Returns the number of `value<N>` attributes that will return a valid value
// for the current mode.
// sysfs filename: num_values
func (sensor *Sensor) NumValues() int {
	return sensor.GetAttrInt("num_values")
}

// "static method": returns a "constant", does not alter device state
// func (_ Sensor) Property_NumValues() string { return "num_values" }

// Returns the name of the port that the sensor is connected to, e.g. `ev3:in1`.
// I2C sensors also include the I2C address (decimal), e.g. `ev3:in1:i2c8`.
// sysfs filename: port_name
func (sensor *Sensor) PortName() string {
	return sensor.GetAttrString("port_name")
}

// "static method": returns a "constant", does not alter device state
// func (_ Sensor) Property_PortName() string { return "port_name" }

// Returns the units of the measured value for the current mode. May return
// empty string
// sysfs filename: units
func (sensor *Sensor) Units() string {
	return sensor.GetAttrString("units")
}

// "static method": returns a "constant", does not alter device state
// func (_ Sensor) Property_Units() string { return "units" }
