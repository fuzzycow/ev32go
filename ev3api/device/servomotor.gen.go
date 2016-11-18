// ServoMotor (Servo Motor):
// The servo motor class provides a uniform interface for using hobby type
// servo motors.
// ev3dev documentation: http://www.ev3dev.org/docs/drivers/servo-motor-class/
// sysfs naming convention: servo-motor/motor{0}
package device

import "github.com/fuzzycow/ev32go/ev3api"

type ServoMotor struct{ ev3api.Device }

func (_ ServoMotor) SystemClassName() string            { return "servo-motor" }
func (_ ServoMotor) SystemDeviceNameConvention() string { return "motor{0}" }

var ServoMotorPropertyNames = []string{

	"command",
	"driver_name",
	"max_pulse_sp",
	"mid_pulse_sp",
	"min_pulse_sp",
	"polarity",
	"port_name",
	"position_sp",
	"rate_sp",
	"state",
}

func (servoMotor ServoMotor) PropertyNames() []string {

	return ServoMotorPropertyNames

}

// Sets the command for the servo. Valid values are `run` and `float`. Setting
// to `run` will cause the servo to be driven to the position_sp set in the
// `position_sp` attribute. Setting to `float` will remove power from the motor.
// sysfs filename: command
func (servoMotor *ServoMotor) SetCommand(value string) {
	servoMotor.SetAttrString("command", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ ServoMotor) Property_Command() string { return "command" }

// Returns the name of the motor driver that loaded this device. See the list
// of [supported devices] for a list of drivers.
// sysfs filename: driver_name
func (servoMotor *ServoMotor) DriverName() string {
	return servoMotor.GetAttrString("driver_name")
}

// "static method": returns a "constant", does not alter device state
// func (_ ServoMotor) Property_DriverName() string { return "driver_name" }

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the maximum (clockwise) position_sp. Default value is 2400.
// Valid values are 2300 to 2700. You must write to the position_sp attribute for
// changes to this attribute to take effect.
// sysfs filename: max_pulse_sp
func (servoMotor *ServoMotor) MaxPulseSP() int {
	return servoMotor.GetAttrInt("max_pulse_sp")
}

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the maximum (clockwise) position_sp. Default value is 2400.
// Valid values are 2300 to 2700. You must write to the position_sp attribute for
// changes to this attribute to take effect.
// sysfs filename: max_pulse_sp
func (servoMotor *ServoMotor) SetMaxPulseSP(value int) {
	servoMotor.SetAttrInt("max_pulse_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ ServoMotor) Property_MaxPulseSP() string { return "max_pulse_sp" }

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the mid position_sp. Default value is 1500. Valid
// values are 1300 to 1700. For example, on a 180 degree servo, this would be
// 90 degrees. On continuous rotation servo, this is the 'neutral' position_sp
// where the motor does not turn. You must write to the position_sp attribute for
// changes to this attribute to take effect.
// sysfs filename: mid_pulse_sp
func (servoMotor *ServoMotor) MidPulseSP() int {
	return servoMotor.GetAttrInt("mid_pulse_sp")
}

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the mid position_sp. Default value is 1500. Valid
// values are 1300 to 1700. For example, on a 180 degree servo, this would be
// 90 degrees. On continuous rotation servo, this is the 'neutral' position_sp
// where the motor does not turn. You must write to the position_sp attribute for
// changes to this attribute to take effect.
// sysfs filename: mid_pulse_sp
func (servoMotor *ServoMotor) SetMidPulseSP(value int) {
	servoMotor.SetAttrInt("mid_pulse_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ ServoMotor) Property_MidPulseSP() string { return "mid_pulse_sp" }

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the miniumum (counter-clockwise) position_sp. Default value
// is 600. Valid values are 300 to 700. You must write to the position_sp
// attribute for changes to this attribute to take effect.
// sysfs filename: min_pulse_sp
func (servoMotor *ServoMotor) MinPulseSP() int {
	return servoMotor.GetAttrInt("min_pulse_sp")
}

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the miniumum (counter-clockwise) position_sp. Default value
// is 600. Valid values are 300 to 700. You must write to the position_sp
// attribute for changes to this attribute to take effect.
// sysfs filename: min_pulse_sp
func (servoMotor *ServoMotor) SetMinPulseSP(value int) {
	servoMotor.SetAttrInt("min_pulse_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ ServoMotor) Property_MinPulseSP() string { return "min_pulse_sp" }

// Sets the polarity of the servo. Valid values are `normal` and `inversed`.
// Setting the value to `inversed` will cause the position_sp value to be
// inversed. i.e `-100` will correspond to `max_pulse_sp`, and `100` will
// correspond to `min_pulse_sp`.
// sysfs filename: polarity
func (servoMotor *ServoMotor) Polarity() string {
	return servoMotor.GetAttrString("polarity")
}

// Sets the polarity of the servo. Valid values are `normal` and `inversed`.
// Setting the value to `inversed` will cause the position_sp value to be
// inversed. i.e `-100` will correspond to `max_pulse_sp`, and `100` will
// correspond to `min_pulse_sp`.
// sysfs filename: polarity
func (servoMotor *ServoMotor) SetPolarity(value string) {
	servoMotor.SetAttrString("polarity", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ ServoMotor) Property_Polarity() string { return "polarity" }

// Returns the name of the port that the motor is connected to.
// sysfs filename: port_name
func (servoMotor *ServoMotor) PortName() string {
	return servoMotor.GetAttrString("port_name")
}

// "static method": returns a "constant", does not alter device state
// func (_ ServoMotor) Property_PortName() string { return "port_name" }

// Reading returns the current position_sp of the servo. Writing instructs the
// servo to move to the specified position_sp. Units are percent. Valid values
// are -100 to 100 (-100% to 100%) where `-100` corresponds to `min_pulse_sp`,
// `0` corresponds to `mid_pulse_sp` and `100` corresponds to `max_pulse_sp`.
// sysfs filename: position_sp
func (servoMotor *ServoMotor) PositionSP() int {
	return servoMotor.GetAttrInt("position_sp")
}

// Reading returns the current position_sp of the servo. Writing instructs the
// servo to move to the specified position_sp. Units are percent. Valid values
// are -100 to 100 (-100% to 100%) where `-100` corresponds to `min_pulse_sp`,
// `0` corresponds to `mid_pulse_sp` and `100` corresponds to `max_pulse_sp`.
// sysfs filename: position_sp
func (servoMotor *ServoMotor) SetPositionSP(value int) {
	servoMotor.SetAttrInt("position_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ ServoMotor) Property_PositionSP() string { return "position_sp" }

// Sets the rate_sp at which the servo travels from 0 to 100.0% (half of the full
// range of the servo). Units are in milliseconds. Example: Setting the rate_sp
// to 1000 means that it will take a 180 degree servo 2 second to move from 0
// to 180 degrees. Note: Some servo controllers may not support this in which
// case reading and writing will fail with `-EOPNOTSUPP`. In continuous rotation
// servos, this value will affect the rate_sp at which the speed ramps up or down.
// sysfs filename: rate_sp
func (servoMotor *ServoMotor) RateSP() int {
	return servoMotor.GetAttrInt("rate_sp")
}

// Sets the rate_sp at which the servo travels from 0 to 100.0% (half of the full
// range of the servo). Units are in milliseconds. Example: Setting the rate_sp
// to 1000 means that it will take a 180 degree servo 2 second to move from 0
// to 180 degrees. Note: Some servo controllers may not support this in which
// case reading and writing will fail with `-EOPNOTSUPP`. In continuous rotation
// servos, this value will affect the rate_sp at which the speed ramps up or down.
// sysfs filename: rate_sp
func (servoMotor *ServoMotor) SetRateSP(value int) {
	servoMotor.SetAttrInt("rate_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ ServoMotor) Property_RateSP() string { return "rate_sp" }

// Returns a list of flags indicating the state of the servo.
// Possible values are:
// * `running`: Indicates that the motor is powered.
// sysfs filename: state
func (servoMotor *ServoMotor) State() []string {
	return servoMotor.GetAttrStringArray("state")
}

// "static method": returns a "constant", does not alter device state
// func (_ ServoMotor) Property_State() string { return "state" }

// Drive servo to the position set in the `position_sp` attribute.
// "static method": returns a "constant", does not alter device state
func (_ ServoMotor) Command_Run() string { return "run" }

// Drive servo to the position set in the `position_sp` attribute.
func (servoMotor *ServoMotor) Run() {
	servoMotor.SetCommand(servoMotor.Command_Run())
}

// Remove power from the motor.
// "static method": returns a "constant", does not alter device state
func (_ ServoMotor) Command_Float() string { return "float" }

// Remove power from the motor.
func (servoMotor *ServoMotor) Float() {
	servoMotor.SetCommand(servoMotor.Command_Float())
}

// With `normal` polarity, a positive duty cycle will
// cause the motor to rotate clockwise.
// "static method": returns a "constant", does not alter device state
func (_ ServoMotor) Polarity_Normal() string { return "normal" }

// With `inversed` polarity, a positive duty cycle will
// cause the motor to rotate counter-clockwise.
// "static method": returns a "constant", does not alter device state
func (_ ServoMotor) Polarity_Inversed() string { return "inversed" }
