package interface_test
import "testing"

type Fooer interface {
	Foo()
}

type Bar struct {
}

func (b *Bar) Foo() {}
func (b *Bar) Bar() {}

func BenchmarkStruct(b *testing.B) {
	bar := &Bar{}
	for i:=0;i<b.N;i++ {
		bar.Foo()
	}
}

func BenchmarkInterface(b *testing.B) {
	fooer := Fooer(&Bar{})
	for i:=0;i<b.N;i++ {
		fooer.Foo()
	}
}




