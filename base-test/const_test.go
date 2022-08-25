package base

import (
	"fmt"
	"testing"
)

const (
	X  =  100
	Y = 10
	Z = X / Y
)

func TestConst(t *testing.T)  {
	fmt.Println(Z)
}

func TestSprintf(t *testing.T) {
	sprintf := fmt.Sprintf("%s", make(map[string]interface{}))
	fmt.Println(sprintf)
}

