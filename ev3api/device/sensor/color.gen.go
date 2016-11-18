// Color (Color Sensor):
// LEGO EV3 color sensor.
// ev3dev documentation: http://www.ev3dev.org/docs/sensors/lego-ev3-color-sensor/
// sysfs naming convention: lego-sensor/sensor{0}
package sensor

import "github.com/fuzzycow/ev32go/ev3api/device"

type Color struct{ *device.Sensor }

func (_ Color) SystemClassName() string            { return "lego-sensor" }
func (_ Color) SystemDeviceNameConvention() string { return "sensor{0}" }

var ColorPropertyNames = []string{}

func (color Color) PropertyNames() []string {

	return append(device.SensorPropertyNames, ColorPropertyNames...)

}

// Reflected light. Red LED on.
// "static method": returns a "constant", does not alter device state
func (_ Color) Mode_COL_REFLECT() string { return "COL-REFLECT" }

// Ambient light. Red LEDs off.
// "static method": returns a "constant", does not alter device state
func (_ Color) Mode_COL_AMBIENT() string { return "COL-AMBIENT" }

// Color. All LEDs rapidly cycling, appears white.
// "static method": returns a "constant", does not alter device state
func (_ Color) Mode_COL_COLOR() string { return "COL-COLOR" }

// Raw reflected. Red LED on
// "static method": returns a "constant", does not alter device state
func (_ Color) Mode_REF_RAW() string { return "REF-RAW" }

// Raw Color Components. All LEDs rapidly cycling, appears white.
// "static method": returns a "constant", does not alter device state
func (_ Color) Mode_RGB_RAW() string { return "RGB-RAW" }
