package interface_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInterfaceNil(t *testing.T) {
	var a interface{} = nil         // tab = nil, data = nil
	var b interface{} = (*int)(nil) // tab 包含 *int 类型信息, data = nil

	fmt.Println(a == nil)
	fmt.Println(b == nil)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	fmt.Println("IsNil(a)-------------------------------------------------------------------------------", IsNil(a))
	fmt.Println("IsNil(b)-------------------------------------------------------------------------------", IsNil(b))
}

func IsNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}
