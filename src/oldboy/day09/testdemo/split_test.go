package testdemo

import "testing"

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mySplit("a:b:c", ":")
	}
}
