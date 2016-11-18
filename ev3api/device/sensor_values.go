package device

import (
	"strconv"
)

func (dev *Sensor) Value() int {
	return dev.GetAttrInt("value0")
}

func (dev *Sensor) ValueN(n int) int {
	attr := "value" + strconv.Itoa(n)
	return dev.GetAttrInt(attr)
}

