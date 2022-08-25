package main

import (
	"errors"
)

func main() {
	myIndex := 4
	ia := [3]int{1,2,3}
	_ = ia[myIndex]
}

func outerFunc() {
	innerFunc()
}

func innerFunc() {
	panic(errors.New("an intended fatal error"))
}
