// Gyro (Gyro Sensor):
// LEGO EV3 gyro sensor.
// ev3dev documentation: http://www.ev3dev.org/docs/sensors/lego-ev3-gyro-sensor/
// sysfs naming convention: lego-sensor/sensor{0}
package sensor

import "github.com/fuzzycow/ev32go/ev3api/device"

type Gyro struct{ *device.Sensor }

func (_ Gyro) SystemClassName() string            { return "lego-sensor" }
func (_ Gyro) SystemDeviceNameConvention() string { return "sensor{0}" }

var GyroPropertyNames = []string{}

func (gyro Gyro) PropertyNames() []string {

	return append(device.SensorPropertyNames, GyroPropertyNames...)

}

// Angle
// "static method": returns a "constant", does not alter device state
func (_ Gyro) Mode_GYRO_ANG() string { return "GYRO-ANG" }

// Rotational speed
// "static method": returns a "constant", does not alter device state
func (_ Gyro) Mode_GYRO_RATE() string { return "GYRO-RATE" }

// Raw sensor value
// "static method": returns a "constant", does not alter device state
func (_ Gyro) Mode_GYRO_FAS() string { return "GYRO-FAS" }

// Angle and rotational speed
// "static method": returns a "constant", does not alter device state
func (_ Gyro) Mode_GYRO_GA() string { return "GYRO-G&A" }

// Calibration ???
// "static method": returns a "constant", does not alter device state
func (_ Gyro) Mode_GYRO_CAL() string { return "GYRO-CAL" }
