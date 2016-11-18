// I2C (I2C Sensor):
// A generic interface to control I2C-type EV3 sensors.
// ev3dev documentation:
// sysfs naming convention: lego-sensor/sensor{0}
package sensor

import "github.com/fuzzycow/ev32go/ev3api/device"

type I2C struct{ *device.Sensor }

func (_ I2C) SystemClassName() string            { return "lego-sensor" }
func (_ I2C) SystemDeviceNameConvention() string { return "sensor{0}" }

var I2CPropertyNames = []string{

	"fw_version",
	"poll_ms",
}

func (i2c I2C) PropertyNames() []string {

	return append(device.SensorPropertyNames, I2CPropertyNames...)

}

// Returns the firmware version of the sensor if available. Currently only
// I2C/NXT sensors support this.
// sysfs filename: fw_version
func (i2c *I2C) FWVersion() string {
	return i2c.GetAttrString("fw_version")
}

// "static method": returns a "constant", does not alter device state
// func (_ I2C) Property_FWVersion() string { return "fw_version" }

// Returns the polling period of the sensor in milliseconds. Writing sets the
// polling period. Setting to 0 disables polling. Minimum value is hard
// coded as 50 msec. Returns -EOPNOTSUPP if changing polling is not supported.
// Currently only I2C/NXT sensors support changing the polling period.
// sysfs filename: poll_ms
func (i2c *I2C) PollMS() int {
	return i2c.GetAttrInt("poll_ms")
}

// Returns the polling period of the sensor in milliseconds. Writing sets the
// polling period. Setting to 0 disables polling. Minimum value is hard
// coded as 50 msec. Returns -EOPNOTSUPP if changing polling is not supported.
// Currently only I2C/NXT sensors support changing the polling period.
// sysfs filename: poll_ms
func (i2c *I2C) SetPollMS(value int) {
	i2c.SetAttrInt("poll_ms", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ I2C) Property_PollMS() string { return "poll_ms" }
