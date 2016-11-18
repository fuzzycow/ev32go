// Motor (Motor):
// The motor class provides a uniform interface for using motors with
// positional and directional feedback such as the EV3 and NXT motors.
// This feedback allows for precise control of the motors. This is the
// most common type of motor, so we just call it `motor`.
// ev3dev documentation: http://www.ev3dev.org/docs/drivers/tacho-motor-class/
// sysfs naming convention: tacho-motor/motor{0}
package device

import "github.com/fuzzycow/ev32go/ev3api"

type Motor struct{ ev3api.Device }

func (_ Motor) SystemClassName() string            { return "tacho-motor" }
func (_ Motor) SystemDeviceNameConvention() string { return "motor{0}" }

var MotorPropertyNames = []string{

	"command",
	"commands",
	"count_per_rot",
	"driver_name",
	"duty_cycle",
	"duty_cycle_sp",
	"encoder_polarity",
	"polarity",
	"port_name",
	"position",
	"hold_pid/Kp",
	"hold_pid/Ki",
	"hold_pid/Kd",
	"position_sp",
	"speed",
	"speed_sp",
	"ramp_up_sp",
	"ramp_down_sp",
	"speed_regulation",
	"speed_pid/Kp",
	"speed_pid/Ki",
	"speed_pid/Kd",
	"state",
	"stop_command",
	"stop_commands",
	"time_sp",
}

func (motor Motor) PropertyNames() []string {

	return MotorPropertyNames

}

// Sends a command to the motor controller. See `commands` for a list of
// possible values.
// sysfs filename: command
func (motor *Motor) SetCommand(value string) {
	motor.SetAttrString("command", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_Command() string { return "command" }

// Returns a list of commands that are supported by the motor
// controller. Possible values are `run-forever`, `run-to-abs-pos`, `run-to-rel-pos`,
// `run-timed`, `run-direct`, `stop` and `reset`. Not all commands may be supported.
// `run-forever` will cause the motor to run until another command is sent.
// `run-to-abs-pos` will run to an absolute position specified by `position_sp`
// and then stop using the command specified in `stop_command`.
// `run-to-rel-pos` will run to a position relative to the current `position` value.
// The new position will be current `position` + `position_sp`. When the new
// position is reached, the motor will stop using the command specified by `stop_command`.
// `run-timed` will run the motor for the amount of time specified in `time_sp`
// and then stop the motor using the command specified by `stop_command`.
// `run-direct` will run the motor at the duty cycle specified by `duty_cycle_sp`.
// Unlike other run commands, changing `duty_cycle_sp` while running *will*
// take effect immediately.
// `stop` will stop any of the run commands before they are complete using the
// command specified by `stop_command`.
// `reset` will reset all of the motor parameter attributes to their default value.
// This will also have the effect of stopping the motor.
// sysfs filename: commands
func (motor *Motor) Commands() []string {
	return motor.GetAttrStringArray("commands")
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_Commands() string { return "commands" }

// Returns the number of tacho counts in one rotation of the motor. Tacho counts
// are used by the position and speed attributes, so you can use this value
// to convert rotations or degrees to tacho counts. In the case of linear
// actuators, the units here will be counts per centimeter.
// sysfs filename: count_per_rot
func (motor *Motor) CountPerRot() int {
	return motor.GetAttrInt("count_per_rot")
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_CountPerRot() string { return "count_per_rot" }

// Returns the name of the driver that provides this tacho motor device.
// sysfs filename: driver_name
func (motor *Motor) DriverName() string {
	return motor.GetAttrString("driver_name")
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_DriverName() string { return "driver_name" }

// Returns the current duty cycle of the motor. Units are percent. Values
// are -100 to 100.
// sysfs filename: duty_cycle
func (motor *Motor) DutyCycle() int {
	return motor.GetAttrInt("duty_cycle")
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_DutyCycle() string { return "duty_cycle" }

// Writing sets the duty cycle setpoint. Reading returns the current value.
// Units are in percent. Valid values are -100 to 100. A negative value causes
// the motor to rotate in reverse. This value is only used when `speed_regulation`
// is off.
// sysfs filename: duty_cycle_sp
func (motor *Motor) DutyCycleSP() int {
	return motor.GetAttrInt("duty_cycle_sp")
}

// Writing sets the duty cycle setpoint. Reading returns the current value.
// Units are in percent. Valid values are -100 to 100. A negative value causes
// the motor to rotate in reverse. This value is only used when `speed_regulation`
// is off.
// sysfs filename: duty_cycle_sp
func (motor *Motor) SetDutyCycleSP(value int) {
	motor.SetAttrInt("duty_cycle_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_DutyCycleSP() string { return "duty_cycle_sp" }

// Sets the polarity of the rotary encoder. This is an advanced feature to all
// use of motors that send inversed encoder signals to the EV3. This should
// be set correctly by the driver of a device. It You only need to change this
// value if you are using a unsupported device. Valid values are `normal` and
// `inversed`.
// sysfs filename: encoder_polarity
func (motor *Motor) EncoderPolarity() string {
	return motor.GetAttrString("encoder_polarity")
}

// Sets the polarity of the rotary encoder. This is an advanced feature to all
// use of motors that send inversed encoder signals to the EV3. This should
// be set correctly by the driver of a device. It You only need to change this
// value if you are using a unsupported device. Valid values are `normal` and
// `inversed`.
// sysfs filename: encoder_polarity
func (motor *Motor) SetEncoderPolarity(value string) {
	motor.SetAttrString("encoder_polarity", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_EncoderPolarity() string { return "encoder_polarity" }

// Sets the polarity of the motor. With `normal` polarity, a positive duty
// cycle will cause the motor to rotate clockwise. With `inversed` polarity,
// a positive duty cycle will cause the motor to rotate counter-clockwise.
// Valid values are `normal` and `inversed`.
// sysfs filename: polarity
func (motor *Motor) Polarity() string {
	return motor.GetAttrString("polarity")
}

// Sets the polarity of the motor. With `normal` polarity, a positive duty
// cycle will cause the motor to rotate clockwise. With `inversed` polarity,
// a positive duty cycle will cause the motor to rotate counter-clockwise.
// Valid values are `normal` and `inversed`.
// sysfs filename: polarity
func (motor *Motor) SetPolarity(value string) {
	motor.SetAttrString("polarity", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_Polarity() string { return "polarity" }

// Returns the name of the port that the motor is connected to.
// sysfs filename: port_name
func (motor *Motor) PortName() string {
	return motor.GetAttrString("port_name")
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_PortName() string { return "port_name" }

// Returns the current position of the motor in pulses of the rotary
// encoder. When the motor rotates clockwise, the position will increase.
// Likewise, rotating counter-clockwise causes the position to decrease.
// Writing will set the position to that value.
// sysfs filename: position
func (motor *Motor) Position() int {
	return motor.GetAttrInt("position")
}

// Returns the current position of the motor in pulses of the rotary
// encoder. When the motor rotates clockwise, the position will increase.
// Likewise, rotating counter-clockwise causes the position to decrease.
// Writing will set the position to that value.
// sysfs filename: position
func (motor *Motor) SetPosition(value int) {
	motor.SetAttrInt("position", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_Position() string { return "position" }

// The proportional constant for the position PID.
// sysfs filename: hold_pid/Kp
func (motor *Motor) PositionP() int {
	return motor.GetAttrInt("hold_pid/Kp")
}

// The proportional constant for the position PID.
// sysfs filename: hold_pid/Kp
func (motor *Motor) SetPositionP(value int) {
	motor.SetAttrInt("hold_pid/Kp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_PositionP() string { return "hold_pid/Kp" }

// The integral constant for the position PID.
// sysfs filename: hold_pid/Ki
func (motor *Motor) PositionI() int {
	return motor.GetAttrInt("hold_pid/Ki")
}

// The integral constant for the position PID.
// sysfs filename: hold_pid/Ki
func (motor *Motor) SetPositionI(value int) {
	motor.SetAttrInt("hold_pid/Ki", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_PositionI() string { return "hold_pid/Ki" }

// The derivative constant for the position PID.
// sysfs filename: hold_pid/Kd
func (motor *Motor) PositionD() int {
	return motor.GetAttrInt("hold_pid/Kd")
}

// The derivative constant for the position PID.
// sysfs filename: hold_pid/Kd
func (motor *Motor) SetPositionD(value int) {
	motor.SetAttrInt("hold_pid/Kd", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_PositionD() string { return "hold_pid/Kd" }

// Writing specifies the target position for the `run-to-abs-pos` and `run-to-rel-pos`
// commands. Reading returns the current value. Units are in tacho counts. You
// can use the value returned by `counts_per_rot` to convert tacho counts to/from
// rotations or degrees.
// sysfs filename: position_sp
func (motor *Motor) PositionSP() int {
	return motor.GetAttrInt("position_sp")
}

// Writing specifies the target position for the `run-to-abs-pos` and `run-to-rel-pos`
// commands. Reading returns the current value. Units are in tacho counts. You
// can use the value returned by `counts_per_rot` to convert tacho counts to/from
// rotations or degrees.
// sysfs filename: position_sp
func (motor *Motor) SetPositionSP(value int) {
	motor.SetAttrInt("position_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_PositionSP() string { return "position_sp" }

// Returns the current motor speed in tacho counts per second. Not, this is
// not necessarily degrees (although it is for LEGO motors). Use the `count_per_rot`
// attribute to convert this value to RPM or deg/sec.
// sysfs filename: speed
func (motor *Motor) Speed() int {
	return motor.GetAttrInt("speed")
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_Speed() string { return "speed" }

// Writing sets the target speed in tacho counts per second used when `speed_regulation`
// is on. Reading returns the current value.  Use the `count_per_rot` attribute
// to convert RPM or deg/sec to tacho counts per second.
// sysfs filename: speed_sp
func (motor *Motor) SpeedSP() int {
	return motor.GetAttrInt("speed_sp")
}

// Writing sets the target speed in tacho counts per second used when `speed_regulation`
// is on. Reading returns the current value.  Use the `count_per_rot` attribute
// to convert RPM or deg/sec to tacho counts per second.
// sysfs filename: speed_sp
func (motor *Motor) SetSpeedSP(value int) {
	motor.SetAttrInt("speed_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_SpeedSP() string { return "speed_sp" }

// Writing sets the ramp up setpoint. Reading returns the current value. Units
// are in milliseconds. When set to a value > 0, the motor will ramp the power
// sent to the motor from 0 to 100% duty cycle over the span of this setpoint
// when starting the motor. If the maximum duty cycle is limited by `duty_cycle_sp`
// or speed regulation, the actual ramp time duration will be less than the setpoint.
// sysfs filename: ramp_up_sp
func (motor *Motor) RampUpSP() int {
	return motor.GetAttrInt("ramp_up_sp")
}

// Writing sets the ramp up setpoint. Reading returns the current value. Units
// are in milliseconds. When set to a value > 0, the motor will ramp the power
// sent to the motor from 0 to 100% duty cycle over the span of this setpoint
// when starting the motor. If the maximum duty cycle is limited by `duty_cycle_sp`
// or speed regulation, the actual ramp time duration will be less than the setpoint.
// sysfs filename: ramp_up_sp
func (motor *Motor) SetRampUpSP(value int) {
	motor.SetAttrInt("ramp_up_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_RampUpSP() string { return "ramp_up_sp" }

// Writing sets the ramp down setpoint. Reading returns the current value. Units
// are in milliseconds. When set to a value > 0, the motor will ramp the power
// sent to the motor from 100% duty cycle down to 0 over the span of this setpoint
// when stopping the motor. If the starting duty cycle is less than 100%, the
// ramp time duration will be less than the full span of the setpoint.
// sysfs filename: ramp_down_sp
func (motor *Motor) RampDownSP() int {
	return motor.GetAttrInt("ramp_down_sp")
}

// Writing sets the ramp down setpoint. Reading returns the current value. Units
// are in milliseconds. When set to a value > 0, the motor will ramp the power
// sent to the motor from 100% duty cycle down to 0 over the span of this setpoint
// when stopping the motor. If the starting duty cycle is less than 100%, the
// ramp time duration will be less than the full span of the setpoint.
// sysfs filename: ramp_down_sp
func (motor *Motor) SetRampDownSP(value int) {
	motor.SetAttrInt("ramp_down_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_RampDownSP() string { return "ramp_down_sp" }

// Turns speed regulation on or off. If speed regulation is on, the motor
// controller will vary the power supplied to the motor to try to maintain the
// speed specified in `speed_sp`. If speed regulation is off, the controller
// will use the power specified in `duty_cycle_sp`. Valid values are `on` and
// `off`.
// sysfs filename: speed_regulation
func (motor *Motor) SpeedRegulationEnabled() string {
	return motor.GetAttrString("speed_regulation")
}

// Turns speed regulation on or off. If speed regulation is on, the motor
// controller will vary the power supplied to the motor to try to maintain the
// speed specified in `speed_sp`. If speed regulation is off, the controller
// will use the power specified in `duty_cycle_sp`. Valid values are `on` and
// `off`.
// sysfs filename: speed_regulation
func (motor *Motor) SetSpeedRegulationEnabled(value string) {
	motor.SetAttrString("speed_regulation", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_SpeedRegulationEnabled() string { return "speed_regulation" }

// The proportional constant for the speed regulation PID.
// sysfs filename: speed_pid/Kp
func (motor *Motor) SpeedRegulationP() int {
	return motor.GetAttrInt("speed_pid/Kp")
}

// The proportional constant for the speed regulation PID.
// sysfs filename: speed_pid/Kp
func (motor *Motor) SetSpeedRegulationP(value int) {
	motor.SetAttrInt("speed_pid/Kp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_SpeedRegulationP() string { return "speed_pid/Kp" }

// The integral constant for the speed regulation PID.
// sysfs filename: speed_pid/Ki
func (motor *Motor) SpeedRegulationI() int {
	return motor.GetAttrInt("speed_pid/Ki")
}

// The integral constant for the speed regulation PID.
// sysfs filename: speed_pid/Ki
func (motor *Motor) SetSpeedRegulationI(value int) {
	motor.SetAttrInt("speed_pid/Ki", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_SpeedRegulationI() string { return "speed_pid/Ki" }

// The derivative constant for the speed regulation PID.
// sysfs filename: speed_pid/Kd
func (motor *Motor) SpeedRegulationD() int {
	return motor.GetAttrInt("speed_pid/Kd")
}

// The derivative constant for the speed regulation PID.
// sysfs filename: speed_pid/Kd
func (motor *Motor) SetSpeedRegulationD(value int) {
	motor.SetAttrInt("speed_pid/Kd", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_SpeedRegulationD() string { return "speed_pid/Kd" }

// Reading returns a list of state flags. Possible flags are
// `running`, `ramping` `holding` and `stalled`.
// sysfs filename: state
func (motor *Motor) State() []string {
	return motor.GetAttrStringArray("state")
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_State() string { return "state" }

// Reading returns the current stop command. Writing sets the stop command.
// The value determines the motors behavior when `command` is set to `stop`.
// Also, it determines the motors behavior when a run command completes. See
// `stop_commands` for a list of possible values.
// sysfs filename: stop_command
func (motor *Motor) StopCommand() string {
	return motor.GetAttrString("stop_command")
}

// Reading returns the current stop command. Writing sets the stop command.
// The value determines the motors behavior when `command` is set to `stop`.
// Also, it determines the motors behavior when a run command completes. See
// `stop_commands` for a list of possible values.
// sysfs filename: stop_command
func (motor *Motor) SetStopCommand(value string) {
	motor.SetAttrString("stop_command", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_StopCommand() string { return "stop_command" }

// Returns a list of stop modes supported by the motor controller.
// Possible values are `coast`, `brake` and `hold`. `coast` means that power will
// be removed from the motor and it will freely coast to a stop. `brake` means
// that power will be removed from the motor and a passive electrical load will
// be placed on the motor. This is usually done by shorting the motor terminals
// together. This load will absorb the energy from the rotation of the motors and
// cause the motor to stop more quickly than coasting. `hold` does not remove
// power from the motor. Instead it actively try to hold the motor at the current
// position. If an external force tries to turn the motor, the motor will 'push
// back' to maintain its position.
// sysfs filename: stop_commands
func (motor *Motor) StopCommands() []string {
	return motor.GetAttrStringArray("stop_commands")
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_StopCommands() string { return "stop_commands" }

// Writing specifies the amount of time the motor will run when using the
// `run-timed` command. Reading returns the current value. Units are in
// milliseconds.
// sysfs filename: time_sp
func (motor *Motor) TimeSP() int {
	return motor.GetAttrInt("time_sp")
}

// Writing specifies the amount of time the motor will run when using the
// `run-timed` command. Reading returns the current value. Units are in
// milliseconds.
// sysfs filename: time_sp
func (motor *Motor) SetTimeSP(value int) {
	motor.SetAttrInt("time_sp", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ Motor) Property_TimeSP() string { return "time_sp" }

// Run the motor until another command is sent.
// "static method": returns a "constant", does not alter device state
func (_ Motor) Command_RunForever() string { return "run-forever" }

// Run the motor until another command is sent.
func (motor *Motor) RunForever() {
	motor.SetCommand(motor.Command_RunForever())
}

// Run to an absolute position specified by `position_sp` and then
// stop using the command specified in `stop_command`.
// "static method": returns a "constant", does not alter device state
func (_ Motor) Command_RunToAbsPos() string { return "run-to-abs-pos" }

// Run to an absolute position specified by `position_sp` and then
// stop using the command specified in `stop_command`.
func (motor *Motor) RunToAbsPos() {
	motor.SetCommand(motor.Command_RunToAbsPos())
}

// Run to a position relative to the current `position` value.
// The new position will be current `position` + `position_sp`.
// When the new position is reached, the motor will stop using
// the command specified by `stop_command`.
// "static method": returns a "constant", does not alter device state
func (_ Motor) Command_RunToRelPos() string { return "run-to-rel-pos" }

// Run to a position relative to the current `position` value.
// The new position will be current `position` + `position_sp`.
// When the new position is reached, the motor will stop using
// the command specified by `stop_command`.
func (motor *Motor) RunToRelPos() {
	motor.SetCommand(motor.Command_RunToRelPos())
}

// Run the motor for the amount of time specified in `time_sp`
// and then stop the motor using the command specified by `stop_command`.
// "static method": returns a "constant", does not alter device state
func (_ Motor) Command_RunTimed() string { return "run-timed" }

// Run the motor for the amount of time specified in `time_sp`
// and then stop the motor using the command specified by `stop_command`.
func (motor *Motor) RunTimed() {
	motor.SetCommand(motor.Command_RunTimed())
}

// Run the motor at the duty cycle specified by `duty_cycle_sp`.
// Unlike other run commands, changing `duty_cycle_sp` while running *will*
// take effect immediately.
// "static method": returns a "constant", does not alter device state
func (_ Motor) Command_RunDirect() string { return "run-direct" }

// Run the motor at the duty cycle specified by `duty_cycle_sp`.
// Unlike other run commands, changing `duty_cycle_sp` while running *will*
// take effect immediately.
func (motor *Motor) RunDirect() {
	motor.SetCommand(motor.Command_RunDirect())
}

// Stop any of the run commands before they are complete using the
// command specified by `stop_command`.
// "static method": returns a "constant", does not alter device state
func (_ Motor) Command_Stop() string { return "stop" }

// Stop any of the run commands before they are complete using the
// command specified by `stop_command`.
func (motor *Motor) Stop() {
	motor.SetCommand(motor.Command_Stop())
}

// Reset all of the motor parameter attributes to their default value.
// This will also have the effect of stopping the motor.
// "static method": returns a "constant", does not alter device state
func (_ Motor) Command_Reset() string { return "reset" }

// Reset all of the motor parameter attributes to their default value.
// This will also have the effect of stopping the motor.
func (motor *Motor) Reset() {
	motor.SetCommand(motor.Command_Reset())
}

// Sets the normal polarity of the rotary encoder.
// "static method": returns a "constant", does not alter device state
func (_ Motor) EncoderPolarity_Normal() string { return "normal" }

// Sets the inversed polarity of the rotary encoder.
// "static method": returns a "constant", does not alter device state
func (_ Motor) EncoderPolarity_Inversed() string { return "inversed" }

// With `normal` polarity, a positive duty cycle will
// cause the motor to rotate clockwise.
// "static method": returns a "constant", does not alter device state
func (_ Motor) Polarity_Normal() string { return "normal" }

// With `inversed` polarity, a positive duty cycle will
// cause the motor to rotate counter-clockwise.
// "static method": returns a "constant", does not alter device state
func (_ Motor) Polarity_Inversed() string { return "inversed" }

// The motor controller will vary the power supplied to the motor
// to try to maintain the speed specified in `speed_sp`.
// "static method": returns a "constant", does not alter device state
func (_ Motor) SpeedRegulation_On() string { return "on" }

// The motor controller will use the power specified in `duty_cycle_sp`.
// "static method": returns a "constant", does not alter device state
func (_ Motor) SpeedRegulation_Off() string { return "off" }

// Power will be removed from the motor and it will freely coast to a stop.
// "static method": returns a "constant", does not alter device state
func (_ Motor) StopCommand_Coast() string { return "coast" }

// Power will be removed from the motor and a passive electrical load will
// be placed on the motor. This is usually done by shorting the motor terminals
// together. This load will absorb the energy from the rotation of the motors and
// cause the motor to stop more quickly than coasting.
// "static method": returns a "constant", does not alter device state
func (_ Motor) StopCommand_Brake() string { return "brake" }

// Does not remove power from the motor. Instead it actively try to hold the motor
// at the current position. If an external force tries to turn the motor, the motor
// will ``push back`` to maintain its position.
// "static method": returns a "constant", does not alter device state
func (_ Motor) StopCommand_Hold() string { return "hold" }
