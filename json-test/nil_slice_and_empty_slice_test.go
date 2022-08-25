package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type NilSliceAndEmptySlice struct {
	NilSlice				[]string
	EmptySlice 				[]string
}

func Test_Nil_Slice_And_Empty_Slice(t *testing.T) {
	nilSliceAndEmptySlice := &NilSliceAndEmptySlice{
		EmptySlice: make([]string, 0),
	}

	if nilSliceAndEmptySliceBytes, e := json.Marshal(nilSliceAndEmptySlice); e == nil {
		fmt.Println(string(nilSliceAndEmptySliceBytes))
	}
}
