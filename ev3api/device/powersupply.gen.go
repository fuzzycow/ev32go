// PowerSupply (Power Supply):
// A generic interface to read data from the system's power_supply class.
// Uses the built-in legoev3-battery if none is specified.
// ev3dev documentation:
// sysfs naming convention: power_supply/
package device

import "github.com/fuzzycow/ev32go/ev3api"

type PowerSupply struct{ ev3api.Device }

func (_ PowerSupply) SystemClassName() string            { return "power_supply" }
func (_ PowerSupply) SystemDeviceNameConvention() string { return "" }

var PowerSupplyPropertyNames = []string{

	"current_now",
	"voltage_now",
	"voltage_max_design",
	"voltage_min_design",
	"technology",
	"type",
}

func (powerSupply PowerSupply) PropertyNames() []string {

	return PowerSupplyPropertyNames

}

// The measured current that the battery is supplying (in microamps)
// sysfs filename: current_now
func (powerSupply *PowerSupply) MeasuredCurrent() int {
	return powerSupply.GetAttrInt("current_now")
}

// "static method": returns a "constant", does not alter device state
// func (_ PowerSupply) Property_MeasuredCurrent() string { return "current_now" }

// The measured voltage that the battery is supplying (in microvolts)
// sysfs filename: voltage_now
func (powerSupply *PowerSupply) MeasuredVoltage() int {
	return powerSupply.GetAttrInt("voltage_now")
}

// "static method": returns a "constant", does not alter device state
// func (_ PowerSupply) Property_MeasuredVoltage() string { return "voltage_now" }

// sysfs filename: voltage_max_design
func (powerSupply *PowerSupply) MaxVoltage() int {
	return powerSupply.GetAttrInt("voltage_max_design")
}

// "static method": returns a "constant", does not alter device state
// func (_ PowerSupply) Property_MaxVoltage() string { return "voltage_max_design" }

// sysfs filename: voltage_min_design
func (powerSupply *PowerSupply) MinVoltage() int {
	return powerSupply.GetAttrInt("voltage_min_design")
}

// "static method": returns a "constant", does not alter device state
// func (_ PowerSupply) Property_MinVoltage() string { return "voltage_min_design" }

// sysfs filename: technology
func (powerSupply *PowerSupply) Technology() string {
	return powerSupply.GetAttrString("technology")
}

// "static method": returns a "constant", does not alter device state
// func (_ PowerSupply) Property_Technology() string { return "technology" }

// sysfs filename: type
func (powerSupply *PowerSupply) Type() string {
	return powerSupply.GetAttrString("type")
}

// "static method": returns a "constant", does not alter device state
// func (_ PowerSupply) Property_Type() string { return "type" }
