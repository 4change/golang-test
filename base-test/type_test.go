package base_test

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	// float64 --> int
	fmt.Println(int(1.0))

	fmt.Println(fmt.Printf("%.18f", 0.9))
}
