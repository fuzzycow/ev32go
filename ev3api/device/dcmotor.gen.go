// DCMotor (DC Motor):
// The DC motor class provides a uniform interface for using regular DC motors
// with no fancy controls or feedback. This includes LEGO MINDSTORMS RCX motors
// and LEGO Power Functions motors.
// ev3dev documentation: http://www.ev3dev.org/docs/drivers/dc-motor-class/
// sysfs naming convention: dc-motor/motor{0}
package device

import "github.com/fuzzycow/ev32go/ev3api"

type DCMotor struct{ ev3api.Device }

func (_ DCMotor) SystemClassName() string            { return "dc-motor" }
func (_ DCMotor) SystemDeviceNameConvention() string { return "motor{0}" }

var DCMotorPropertyNames = []string{

	"command",
	"commands",
	"driver_name",
	"duty_cycle",
	"duty_cycle_sp",
	"polarity",
	"port_name",
	"ramp_down_sp",
	"ramp_up_sp",
	"state",
	"stop_command",
	"stop_commands",
}

func (dcMotor DCMotor) PropertyNames() []string {

	return DCMotorPropertyNames

}

// Sets the command for the motor. Possible values are `run-forever`, `run-timed` and
// `stop`. Not all commands may be supported, so be sure to check the contents
// of the `commands` attribute.
// sysfs filename: command
func (dcMotor *DCMotor) SetCommand(value string) {
	dcMotor.SetAttrString("command", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_Command() string { return "command" }

// Returns a list of commands supported by the motor
// controller.
// sysfs filename: commands
func (dcMotor *DCMotor) Commands() []string {
	return dcMotor.GetAttrStringArray("commands")
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_Commands() string { return "commands" }

// Returns the name of the motor driver that loaded this device. See the list
// of [supported devices] for a list of drivers.
// sysfs filename: driver_name
func (dcMotor *DCMotor) DriverName() string {
	return dcMotor.GetAttrString("driver_name")
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_DriverName() string { return "driver_name" }

// Shows the current duty cycle of the PWM signal sent to the motor. Values
// are -100 to 100 (-100% to 100%).
// sysfs filename: duty_cycle
func (dcMotor *DCMotor) DutyCycle() int {
	return dcMotor.GetAttrInt("duty_cycle")
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_DutyCycle() string { return "duty_cycle" }

// Writing sets the duty cycle setpoint of the PWM signal sent to the motor.
// Valid values are -100 to 100 (-100% to 100%). Reading returns the current
// setpoint.
// sysfs filename: duty_cycle_sp
func (dcMotor *DCMotor) DutyCycleSP() int {
	return dcMotor.GetAttrInt("duty_cycle_sp")
}

// Writing sets the duty cycle setpoint of the PWM signal sent to the motor.
// Valid values are -100 to 100 (-100% to 100%). Reading returns the current
// setpoint.
// sysfs filename: duty_cycle_sp
func (dcMotor *DCMotor) SetDutyCycleSP(value int) {
	dcMotor.SetAttrInt("duty_cycle_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_DutyCycleSP() string { return "duty_cycle_sp" }

// Sets the polarity of the motor. Valid values are `normal` and `inversed`.
// sysfs filename: polarity
func (dcMotor *DCMotor) Polarity() string {
	return dcMotor.GetAttrString("polarity")
}

// Sets the polarity of the motor. Valid values are `normal` and `inversed`.
// sysfs filename: polarity
func (dcMotor *DCMotor) SetPolarity(value string) {
	dcMotor.SetAttrString("polarity", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_Polarity() string { return "polarity" }

// Returns the name of the port that the motor is connected to.
// sysfs filename: port_name
func (dcMotor *DCMotor) PortName() string {
	return dcMotor.GetAttrString("port_name")
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_PortName() string { return "port_name" }

// Sets the time in milliseconds that it take the motor to ramp down from 100%
// to 0%. Valid values are 0 to 10000 (10 seconds). Default is 0.
// sysfs filename: ramp_down_sp
func (dcMotor *DCMotor) RampDownSP() int {
	return dcMotor.GetAttrInt("ramp_down_sp")
}

// Sets the time in milliseconds that it take the motor to ramp down from 100%
// to 0%. Valid values are 0 to 10000 (10 seconds). Default is 0.
// sysfs filename: ramp_down_sp
func (dcMotor *DCMotor) SetRampDownSP(value int) {
	dcMotor.SetAttrInt("ramp_down_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_RampDownSP() string { return "ramp_down_sp" }

// Sets the time in milliseconds that it take the motor to up ramp from 0% to
// 100%. Valid values are 0 to 10000 (10 seconds). Default is 0.
// sysfs filename: ramp_up_sp
func (dcMotor *DCMotor) RampUpSP() int {
	return dcMotor.GetAttrInt("ramp_up_sp")
}

// Sets the time in milliseconds that it take the motor to up ramp from 0% to
// 100%. Valid values are 0 to 10000 (10 seconds). Default is 0.
// sysfs filename: ramp_up_sp
func (dcMotor *DCMotor) SetRampUpSP(value int) {
	dcMotor.SetAttrInt("ramp_up_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_RampUpSP() string { return "ramp_up_sp" }

// Gets a list of flags indicating the motor status. Possible
// flags are `running` and `ramping`. `running` indicates that the motor is
// powered. `ramping` indicates that the motor has not yet reached the
// `duty_cycle_sp`.
// sysfs filename: state
func (dcMotor *DCMotor) State() []string {
	return dcMotor.GetAttrStringArray("state")
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_State() string { return "state" }

// Sets the stop command that will be used when the motor stops. Read
// `stop_commands` to get the list of valid values.
// sysfs filename: stop_command
func (dcMotor *DCMotor) SetStopCommand(value string) {
	dcMotor.SetAttrString("stop_command", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_StopCommand() string { return "stop_command" }

// Gets a list of stop commands. Valid values are `coast`
// and `brake`.
// sysfs filename: stop_commands
func (dcMotor *DCMotor) StopCommands() []string {
	return dcMotor.GetAttrStringArray("stop_commands")
}

// "static method": returns a "constant", does not alter device state
// func (_ DCMotor) Property_StopCommands() string { return "stop_commands" }

// Run the motor until another command is sent.
// "static method": returns a "constant", does not alter device state
func (_ DCMotor) Command_RunForever() string { return "run-forever" }

// Run the motor until another command is sent.
func (dcMotor *DCMotor) RunForever() {
	dcMotor.SetCommand(dcMotor.Command_RunForever())
}

// Run the motor for the amount of time specified in `time_sp`
// and then stop the motor using the command specified by `stop_command`.
// "static method": returns a "constant", does not alter device state
func (_ DCMotor) Command_RunTimed() string { return "run-timed" }

// Run the motor for the amount of time specified in `time_sp`
// and then stop the motor using the command specified by `stop_command`.
func (dcMotor *DCMotor) RunTimed() {
	dcMotor.SetCommand(dcMotor.Command_RunTimed())
}

// Run the motor at the duty cycle specified by `duty_cycle_sp`.
// Unlike other run commands, changing `duty_cycle_sp` while running *will*
// take effect immediately.
// "static method": returns a "constant", does not alter device state
func (_ DCMotor) Command_RunDirect() string { return "run-direct" }

// Run the motor at the duty cycle specified by `duty_cycle_sp`.
// Unlike other run commands, changing `duty_cycle_sp` while running *will*
// take effect immediately.
func (dcMotor *DCMotor) RunDirect() {
	dcMotor.SetCommand(dcMotor.Command_RunDirect())
}

// Stop any of the run commands before they are complete using the
// command specified by `stop_command`.
// "static method": returns a "constant", does not alter device state
func (_ DCMotor) Command_Stop() string { return "stop" }

// Stop any of the run commands before they are complete using the
// command specified by `stop_command`.
func (dcMotor *DCMotor) Stop() {
	dcMotor.SetCommand(dcMotor.Command_Stop())
}

// With `normal` polarity, a positive duty cycle will
// cause the motor to rotate clockwise.
// "static method": returns a "constant", does not alter device state
func (_ DCMotor) Polarity_Normal() string { return "normal" }

// With `inversed` polarity, a positive duty cycle will
// cause the motor to rotate counter-clockwise.
// "static method": returns a "constant", does not alter device state
func (_ DCMotor) Polarity_Inversed() string { return "inversed" }

// Power will be removed from the motor and it will freely coast to a stop.
// "static method": returns a "constant", does not alter device state
func (_ DCMotor) StopCommand_Coast() string { return "coast" }

// Power will be removed from the motor and a passive electrical load will
// be placed on the motor. This is usually done by shorting the motor terminals
// together. This load will absorb the energy from the rotation of the motors and
// cause the motor to stop more quickly than coasting.
// "static method": returns a "constant", does not alter device state
func (_ DCMotor) StopCommand_Brake() string { return "brake" }
