// +build arm
package sysfs
import (
	"testing"
	"bytes"
	"runtime"
)


func OnlyOnArm(b *testing.B) {
	if runtime.GOARCH != "arm" {
		b.Skip("skipping test on non-ev3dev platform")
	}
}



func benchmarkSysfsRead(b *testing.B, dir Dir) {
	defer dir.Close()
	var prev []byte
	for i := 0; i < b.N; i++ {
		buf, err := dir.ReadFile("port0/status")
		if err != nil {
			b.Fatal(err)
		}
		if i > 0 && ! bytes.Equal(buf, prev) {
			b.Error("inconsistent reads")
		}
		prev = buf
	}

}

func BenchmarkSysfsReadCached(b *testing.B) {
	OnlyOnArm(b)
	d := "/sys/class/lego-port"
	b.Logf("using directory %s", d)
	dir := NewCachedDir(d)
	benchmarkSysfsRead(b, dir)
}

func BenchmarkSysfsReadNonCached(b *testing.B) {
	OnlyOnArm(b)
	d := "/sys/class/lego-port"
	b.Logf("using directory %s", d)
	dir := NewDirectDir(d)
	benchmarkSysfsRead(b, dir)
}

