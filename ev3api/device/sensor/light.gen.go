// Light (Light Sensor):
// LEGO NXT Light Sensor
// ev3dev documentation: http://www.ev3dev.org/docs/sensors/lego-nxt-light-sensor/
// sysfs naming convention: lego-nxt-light/sensor{0}
package sensor

import "github.com/fuzzycow/ev32go/ev3api/device"

type Light struct{ *device.Sensor }

func (_ Light) SystemClassName() string            { return "lego-nxt-light" }
func (_ Light) SystemDeviceNameConvention() string { return "sensor{0}" }

var LightPropertyNames = []string{}

func (light Light) PropertyNames() []string {

	return append(device.SensorPropertyNames, LightPropertyNames...)

}

// Reflected light. LED on
// "static method": returns a "constant", does not alter device state
func (_ Light) Mode_REFLECT() string { return "REFLECT" }

// Ambient light. LED off
// "static method": returns a "constant", does not alter device state
func (_ Light) Mode_AMBIENT() string { return "AMBIENT" }
