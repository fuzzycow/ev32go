// Infrared (Infrared Sensor):
// LEGO EV3 infrared sensor.
// ev3dev documentation: http://www.ev3dev.org/docs/sensors/lego-ev3-infrared-sensor/
// sysfs naming convention: lego-sensor/sensor{0}
package sensor

import "github.com/fuzzycow/ev32go/ev3api/device"

type Infrared struct{ *device.Sensor }

func (_ Infrared) SystemClassName() string            { return "lego-sensor" }
func (_ Infrared) SystemDeviceNameConvention() string { return "sensor{0}" }

var InfraredPropertyNames = []string{}

func (infrared Infrared) PropertyNames() []string {

	return append(device.SensorPropertyNames, InfraredPropertyNames...)

}

// Proximity
// "static method": returns a "constant", does not alter device state
func (_ Infrared) Mode_IR_PROX() string { return "IR-PROX" }

// IR Seeker
// "static method": returns a "constant", does not alter device state
func (_ Infrared) Mode_IR_SEEK() string { return "IR-SEEK" }

// IR Remote Control
// "static method": returns a "constant", does not alter device state
func (_ Infrared) Mode_IR_REMOTE() string { return "IR-REMOTE" }

// IR Remote Control. State of the buttons is coded in binary
// "static method": returns a "constant", does not alter device state
func (_ Infrared) Mode_IR_REM_A() string { return "IR-REM-A" }

// Calibration ???
// "static method": returns a "constant", does not alter device state
func (_ Infrared) Mode_IR_CAL() string { return "IR-CAL" }
