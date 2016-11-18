package bench

import (
	"testing"
	"strconv"
	"fmt"
	"strings"
)

func BenchmarkBaselineStrconvAtoi(b *testing.B) {
	for i:=0;i<b.N;i++ {
		_,_ = strconv.Atoi("1234")
	}
}

func BenchmarkBaselineStrconvItoa(b *testing.B) {
	for i:=0;i<b.N;i++ {
		_ = strconv.Itoa(i)
	}
}

func BenchmarkBaselineFmtSprintf(b *testing.B) {
	for i:=0;i<b.N;i++ {
		_ = fmt.Sprintf("%d",i)
	}
}


func BenchmarkBaselineStringJoin(b *testing.B) {
	var s string
	for i:=0;i<b.N;i++ {
		s = strings.Join([]string{"ab","cd","ef"},"")
		_ = s
	}
}

func BenchmarkBaselineStringPlus(b *testing.B) {
	for i:=0;i<b.N;i++ {
		_ = "ab" + "cd" + "ef"
	}
}




