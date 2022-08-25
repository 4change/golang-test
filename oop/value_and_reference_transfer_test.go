package oop

import (
    "testing"
)

type myStruct struct {
    a, b, c int64
    d, e, f string
    g, h, i float64
}

// func byValue() myStruct {
//     return myStruct{
//         a: 1, b: 1, c: 1,
//         d: "foo", e: "bar", f: "baz",
//         g: 1.0, h: 1.0, i: 1.0,
//     }
// }

// func byReference() *myStruct {
//     return &myStruct{
//         a: 1, b: 1, c: 1,
//         d: "foo", e: "bar", f: "baz",
//         g: 1.0, h: 1.0, i: 1.0,
//     }
// }

var s = myStruct{
    a: 1, b: 1, c: 1,
    d: "foo", e: "bar", f: "baz",
    g: 1.0, h: 1.0, i: 1.0,
}

func byValue() myStruct {
    return s
}

func byReference() *myStruct {
    return &s
}

func BenchmarkByValue(b *testing.B) {
    var s myStruct

    for i := 0; i < b.N; i++ {
        // make a copy of the whole struct
        // but do it through stack memory
        s = byValue()
    }

    _ = s
}

func BenchmarkByReference(b *testing.B) {
    var s *myStruct

    for i := 0; i < b.N; i++ {
        // allocate struct on the heap
        // and only return a pointer to it
        s = byReference()
    }

    _ = s
}