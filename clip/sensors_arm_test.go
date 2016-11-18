// +build arm

package clip

import (
	"testing"
	"runtime"
)


func TestLegoPortTracing(t *testing.T) {
	if runtime.GOARCH != "arm" {
		t.Skip("skipping test on non-ev3dev platform")
	}
	dev := NewLegoPort("in1")
	dev.Device.SetTracing(true)
	if err := dev.Open(); err != nil {
		t.Fatalf("Failed to open lego port: %v",err)
	}
	defer dev.Close()
	_ = dev.Mode()
	if dev.Err() != nil {
		t.Fatalf("failed to read from device: %v",dev.Err())
	}
}

/*
func OnlyOnArm(b *testing.B) {
	if runtime.GOARCH != "arm" {
		b.Skip("skipping test on non-ev3dev platform")
	}
}*/

func BenchmarkSysfsAccessorInt(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		dev := NewPowerSupply()
		if err := dev.Open(); err != nil {
			b.Fatalf("Failed to open lego port: %v", err)
		}
		defer dev.Close()
		for pb.Next() {
			_ = dev.MeasuredVoltage()
			if dev.Err() != nil {
				b.Fatalf("failed to read from device: %v", dev.Err())
			}
		}
	})
}

func BenchmarkSysfsAccessorString(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		dev := NewPowerSupply()
		if err := dev.Open(); err != nil {
			b.Fatalf("Failed to open lego port: %v", err)
		}
		defer dev.Close()
		for pb.Next() {
			_ = dev.Type()
			if dev.Err() != nil {
				b.Fatalf("failed to read from device: %v", dev.Err())
			}
		}
	})
}

func BenchmarkSysfsGetAttrString(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		dev := NewPowerSupply()
		if err := dev.Open(); err != nil {
			b.Fatalf("Failed to open lego port: %v", err)
		}
		defer dev.Close()
		for pb.Next() {
			_ = dev.GetAttrString("type")
			if dev.Err() != nil {
				b.Fatalf("failed to read from device: %v", dev.Err())
			}
		}
	})
}


