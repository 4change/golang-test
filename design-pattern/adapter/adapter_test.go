package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	F := NewForwards("F")
	C := NewCenter("C")
	T := NewTranslator("T")

	F.attack()
	C.defense()
	T.attack()
	T.defense()
}
