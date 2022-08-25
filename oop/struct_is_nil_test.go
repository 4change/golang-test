package oop

import (
	"fmt"
	"testing"
	"reflect"
)

// type Person1 struct {
//    id   int64
//    name string
// }

// func TestStructIsNil(t *testing.T) {
//    var foo Person1
//    if (Person1{}) == foo {
//       fmt.Println("foo is empty")
//    }
//    bar := Person1{
//       id:   110,
//       name: "bar",
//    }
//    if (Person1{}) != bar {
//       fmt.Println("bar is not empty")
//    }
// }  

type Person1 struct {
    id   int64
    name string
    addr []byte
}

func TestStructIsNil(t *testing.T) {
    var foo Person1
    if reflect.DeepEqual(foo, Person1{}) {
        fmt.Println("foo is empty")
    }
    bar := Person1{
        id:   110,
        name: "bar",
    }
    if !reflect.DeepEqual(bar, Person1{}) {
        fmt.Println("bar is not empty")
    }
} 