// Ultrasonic (Ultrasonic Sensor):
// LEGO EV3 ultrasonic sensor.
// ev3dev documentation: http://www.ev3dev.org/docs/sensors/lego-ev3-ultrasonic-sensor/
// sysfs naming convention: lego-sensor/sensor{0}
package sensor

import "github.com/fuzzycow/ev32go/ev3api/device"

type Ultrasonic struct{ *device.Sensor }

func (_ Ultrasonic) SystemClassName() string            { return "lego-sensor" }
func (_ Ultrasonic) SystemDeviceNameConvention() string { return "sensor{0}" }

var UltrasonicPropertyNames = []string{}

func (ultrasonic Ultrasonic) PropertyNames() []string {

	return append(device.SensorPropertyNames, UltrasonicPropertyNames...)

}

// Continuous measurement in centimeters.
// LEDs: On, steady
// "static method": returns a "constant", does not alter device state
func (_ Ultrasonic) Mode_US_DIST_CM() string { return "US-DIST-CM" }

// Continuous measurement in inches.
// LEDs: On, steady
// "static method": returns a "constant", does not alter device state
func (_ Ultrasonic) Mode_US_DIST_IN() string { return "US-DIST-IN" }

// Listen.  LEDs: On, blinking
// "static method": returns a "constant", does not alter device state
func (_ Ultrasonic) Mode_US_LISTEN() string { return "US-LISTEN" }

// Single measurement in centimeters.
// LEDs: On momentarily when mode is set, then off
// "static method": returns a "constant", does not alter device state
func (_ Ultrasonic) Mode_US_SI_CM() string { return "US-SI-CM" }

// Single measurement in inches.
// LEDs: On momentarily when mode is set, then off
// "static method": returns a "constant", does not alter device state
func (_ Ultrasonic) Mode_US_SI_IN() string { return "US-SI-IN" }
