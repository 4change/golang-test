package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	type MyStruct struct {
		N int
		M string
	}
	n := MyStruct{
		N: 1,
	}
	// get
	immutable := reflect.ValueOf(n)
	val := immutable.FieldByName("N").Int()
	str := immutable.FieldByName("M").String()
	fmt.Printf("N=%d, %v\n", val, str) // prints 1

	// set
	mutable := reflect.ValueOf(&n).Elem()
	mutable.FieldByName("N").SetInt(7)
	fmt.Printf("N=%d\n", n.N) // prints 7
}
