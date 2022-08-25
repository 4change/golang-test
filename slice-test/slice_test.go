package slice_test

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Compare_Slice_With_Reflect(t *testing.T) {
	a := []string{"A", "B"}
	b := []string{"A", "B"}

	fmt.Println(reflect.DeepEqual(a, b))
}

func Test_Compare_Slice_With_Loop(t *testing.T) {
	a := []string{"A", "B"}
	b := []string{"A", "B"}

	fmt.Println(CompareSliceWithLoop(a, b))
}

func CompareSliceWithLoop(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func BenchmarkEqual(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CompareSliceWithLoop(sa, sb)
	}
}

func BenchmarkDeepEqual(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		reflect.DeepEqual(sa, sb)
	}
}

func CompareSliceWithBCELoop(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func BenchmarkEqualBCE(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CompareSliceWithBCELoop(sa, sb)
	}
}
