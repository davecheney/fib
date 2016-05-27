package fib

import (
	"fmt"
	"testing"
)

var fibTests = []struct {
	n        int // input
	expected int // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func BenchmarkFibVV(b *testing.B) {
	benchSizes := []int{1, 2, 3, 10, 20, 40}
	for _, i := range benchSizes {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Fib(i)
			}
		})
	}
}

var result int

func BenchmarkFibComplete(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = Fib(10)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
