package clip

import (
	"github.com/fuzzycow/ev32go/ev3api/device"
	"github.com/fuzzycow/ev32go/drivers/sysfs"
)

func NewSensor(port string) *device.Sensor {
	drv := sysfs.NewDriver()
	dev := &device.Sensor{Device: drv}
	drv.SetPort(port)
	drv.SetPathFilter(dev.SystemClassName(), dev.SystemDeviceNameConvention())
	return dev
}

func NewMotor(port string) *device.Motor {
	drv := sysfs.NewDriver()
	dev := &device.Motor{Device: drv}
	drv.SetPort(port)
	drv.SetPathFilter(dev.SystemClassName(), dev.SystemDeviceNameConvention())
	return dev
}

func NewDCMotor(port string) *device.DCMotor {
	drv := sysfs.NewDriver()
	dev := &device.DCMotor{Device: drv}
	drv.SetPort(port)
	drv.SetPathFilter(dev.SystemClassName(), dev.SystemDeviceNameConvention())
	return dev
}

func NewServoMotor(port string) *device.ServoMotor {
	drv := sysfs.NewDriver()
	dev := &device.ServoMotor{Device: drv}
	drv.SetPort(port)
	drv.SetPathFilter(dev.SystemClassName(), dev.SystemDeviceNameConvention())
	return dev
}


func NewLegoPort(port string) *device.LegoPort {
	drv := sysfs.NewDriver()
	dev := &device.LegoPort{Device: drv}
	drv.SetPort(port)
	// FIX for spec issue
	drv.SetPathFilter("lego-port", "*")
	return dev
}

func NewLED(led string) *device.LED {
	drv := sysfs.NewDriver()
	dev := &device.LED{Device: drv}
	drv.SetPathFilter(dev.SystemClassName(),led)
	return dev
}


func NewPowerSupply() *device.PowerSupply {
	drv := sysfs.NewDriver()
	dev := &device.PowerSupply{Device: drv}
	//FIXME: Ugly hack
	drv.SetPathFilter(dev.SystemClassName(),"*")
	return dev
}

