// LED (LED):
// Any device controlled by the generic LED driver.
// See https://www.kernel.org/doc/Documentation/leds/leds-class.txt
// for more details.
// ev3dev documentation:
// sysfs naming convention: leds/
package device

import "github.com/fuzzycow/ev32go/ev3api"

type LED struct{ ev3api.Device }

func (_ LED) SystemClassName() string            { return "leds" }
func (_ LED) SystemDeviceNameConvention() string { return "" }

var LEDPropertyNames = []string{

	"max_brightness",
	"brightness",
	"trigger",
	"trigger",
	"delay_on",
	"delay_off",
}

func (led LED) PropertyNames() []string {

	return LEDPropertyNames

}

// Returns the maximum allowable brightness value.
// sysfs filename: max_brightness
func (led *LED) MaxBrightness() int {
	return led.GetAttrInt("max_brightness")
}

// "static method": returns a "constant", does not alter device state
// func (_ LED) Property_MaxBrightness() string { return "max_brightness" }

// Sets the brightness level. Possible values are from 0 to `max_brightness`.
// sysfs filename: brightness
func (led *LED) Brightness() int {
	return led.GetAttrInt("brightness")
}

// Sets the brightness level. Possible values are from 0 to `max_brightness`.
// sysfs filename: brightness
func (led *LED) SetBrightness(value int) {
	led.SetAttrInt("brightness", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ LED) Property_Brightness() string { return "brightness" }

// Returns a list of available triggers.
// sysfs filename: trigger
func (led *LED) Triggers() []string {
	return led.GetAttrStringArray("trigger")
}

// "static method": returns a "constant", does not alter device state
// func (_ LED) Property_Triggers() string { return "trigger" }

// Sets the led trigger. A trigger
// is a kernel based source of led events. Triggers can either be simple or
// complex. A simple trigger isn't configurable and is designed to slot into
// existing subsystems with minimal additional code. Examples are the `ide-disk` and
// `nand-disk` triggers.
//
// Complex triggers whilst available to all LEDs have LED specific
// parameters and work on a per LED basis. The `timer` trigger is an example.
// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `on` and `off` time can
// be specified via `delay_{on,off}` attributes in milliseconds.
// You can change the brightness value of a LED independently of the timer
// trigger. However, if you set the brightness value to 0 it will
// also disable the `timer` trigger.
// sysfs filename: trigger
func (led *LED) Trigger() string {
	return led.GetAttrString("trigger")
}

// Sets the led trigger. A trigger
// is a kernel based source of led events. Triggers can either be simple or
// complex. A simple trigger isn't configurable and is designed to slot into
// existing subsystems with minimal additional code. Examples are the `ide-disk` and
// `nand-disk` triggers.
//
// Complex triggers whilst available to all LEDs have LED specific
// parameters and work on a per LED basis. The `timer` trigger is an example.
// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `on` and `off` time can
// be specified via `delay_{on,off}` attributes in milliseconds.
// You can change the brightness value of a LED independently of the timer
// trigger. However, if you set the brightness value to 0 it will
// also disable the `timer` trigger.
// sysfs filename: trigger
func (led *LED) SetTrigger(value string) {
	led.SetAttrString("trigger", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ LED) Property_Trigger() string { return "trigger" }

// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `on` time can
// be specified via `delay_on` attribute in milliseconds.
// sysfs filename: delay_on
func (led *LED) DelayOn() int {
	return led.GetAttrInt("delay_on")
}

// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `on` time can
// be specified via `delay_on` attribute in milliseconds.
// sysfs filename: delay_on
func (led *LED) SetDelayOn(value int) {
	led.SetAttrInt("delay_on", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ LED) Property_DelayOn() string { return "delay_on" }

// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `off` time can
// be specified via `delay_off` attribute in milliseconds.
// sysfs filename: delay_off
func (led *LED) DelayOff() int {
	return led.GetAttrInt("delay_off")
}

// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `off` time can
// be specified via `delay_off` attribute in milliseconds.
// sysfs filename: delay_off
func (led *LED) SetDelayOff(value int) {
	led.SetAttrInt("delay_off", value)
}

// "static method": returns a "constant", does not alter device state
// func (_ LED) Property_DelayOff() string { return "delay_off" }
