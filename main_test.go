package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareSlices(t *testing.T) {
	got := []string{"a", "b", "c"}

	want := got

	require.NoError(t, CompareSlicesA(got, want))
	require.NoError(t, CompareSlicesB(got, want))
}

func BenchmarkCompareSliceB1(b *testing.B)      { benchmarkCompareSlicesB(1, b) }
func BenchmarkCompareSliceB2(b *testing.B)      { benchmarkCompareSlicesB(2, b) }
func BenchmarkCompareSliceB4(b *testing.B)      { benchmarkCompareSlicesB(4, b) }
func BenchmarkCompareSliceB10(b *testing.B)     { benchmarkCompareSlicesB(10, b) }
func BenchmarkCompareSliceB100(b *testing.B)    { benchmarkCompareSlicesB(100, b) }
func BenchmarkCompareSliceB1000(b *testing.B)   { benchmarkCompareSlicesB(1000, b) }
func BenchmarkCompareSliceB10000(b *testing.B)  { benchmarkCompareSlicesB(10000, b) }
func BenchmarkCompareSliceB100000(b *testing.B) { benchmarkCompareSlicesB(100000, b) }

func BenchmarkCompareSliceA1(b *testing.B)       { benchmarkCompareSlicesA(1, b) }
func BenchmarkCompareSliceA2(b *testing.B)       { benchmarkCompareSlicesA(2, b) }
func BenchmarkCompareSliceA4(b *testing.B)       { benchmarkCompareSlicesA(4, b) }
func BenchmarkCompareSliceA10(b *testing.B)      { benchmarkCompareSlicesA(10, b) }
func BenchmarkCompareSliceA100(b *testing.B)     { benchmarkCompareSlicesA(100, b) }
func BenchmarkCompareSliceA1000(b *testing.B)    { benchmarkCompareSlicesA(1000, b) }
func BenchmarkCompareSliceA10000(b *testing.B)   { benchmarkCompareSlicesA(10000, b) }
func BenchmarkCompareSliceA100000(b *testing.B)  { benchmarkCompareSlicesA(100000, b) }
func BenchmarkCompareSliceA1000000(b *testing.B) { benchmarkCompareSlicesA(1000000, b) }

// Global variable here to stop the compiler optimizing our function calls away
var err error

func benchmarkCompareSlicesA(i int, b *testing.B) {
	b.StopTimer()

	got := generateSlice(i)
	want := generateSlice(i)

	b.StartTimer()
	var e error
	for n := 0; n < b.N; n++ {
		// Store the result of the function under benchmark so the
		// function cannot be eliminated.
		e = CompareSlicesA(got, want)
	}

	// Store the resutl in a package level variable, so the benchmark
	// itself is not eliminated!
	err = e
}

func benchmarkCompareSlicesB(i int, b *testing.B) {
	// Stop the timer while we generate our test inputs
	b.StopTimer()

	got := generateSlice(i)
	want := generateSlice(i)

	b.StartTimer()
	var e error
	for n := 0; n < b.N; n++ {
		e = CompareSlicesB(got, want)
	}

	err = e
}

func generateSlice(n int) []string {
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, strconv.Itoa(i))
	}

	return s
}
