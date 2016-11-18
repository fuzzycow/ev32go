package clip

import (
	"github.com/fuzzycow/ev32go/ev3api/device"
	"github.com/fuzzycow/ev32go/ev3api/device/sensor"
	"github.com/fuzzycow/ev32go/drivers/sysfs"
)

func NewColorSensor(port string) *sensor.Color {
	drv := sysfs.NewDriver()
	dev := &sensor.Color{Sensor: &device.Sensor{Device: drv}}
	drv.SetPort(port)
	drv.SetPathFilter(dev.SystemClassName(), dev.SystemDeviceNameConvention())
	return dev
}

func NewGyroSensor(port string) *sensor.Gyro {
	drv := sysfs.NewDriver()
	dev := &sensor.Gyro{Sensor: &device.Sensor{Device: drv}}
	drv.SetPort(port)
	drv.SetPathFilter(dev.SystemClassName(), dev.SystemDeviceNameConvention())
	return dev
}

func NewI2CSensor(port string) *sensor.I2C {
	drv := sysfs.NewDriver()
	dev := &sensor.I2C{Sensor: &device.Sensor{Device: drv}}
	drv.SetPort(port)
	drv.SetPathFilter(dev.SystemClassName(), dev.SystemDeviceNameConvention())
	return dev
}

func NewInfraredSensor(port string) *sensor.Infrared {
	drv := sysfs.NewDriver()
	dev := &sensor.Infrared{Sensor: &device.Sensor{Device: drv}}
	drv.SetPort(port)
	drv.SetPathFilter(dev.SystemClassName(), dev.SystemDeviceNameConvention())
	return dev
}

func NewLightSensor(port string) *sensor.Light {
	drv := sysfs.NewDriver()
	dev := &sensor.Light{Sensor: &device.Sensor{Device: drv}}
	drv.SetPort(port)
	drv.SetPathFilter(dev.SystemClassName(), dev.SystemDeviceNameConvention())
	return dev
}

func NewUltrasonicSensor(port string) *sensor.Ultrasonic {
	drv := sysfs.NewDriver()
	dev := &sensor.Ultrasonic{Sensor: &device.Sensor{Device: drv}}
	drv.SetPort(port)
	drv.SetPathFilter(dev.SystemClassName(), dev.SystemDeviceNameConvention())
	return dev
}

