package object

import "testing"

func TestStringHashKey(t *testing.T) {
	a1 := &String{Value: "a"}
	a2 := &String{Value: "a"}
	b1 := &String{Value: "b"}
	b2 := &String{Value: "b"}

	if a1.HashKey() != a2.HashKey() {
		t.Errorf("Strings with same content have different hash keys")
	}

	if b1.HashKey() != b2.HashKey() {
		t.Errorf("Strings with same content have different hash keys")
	}

	if a1.HashKey() == b1.HashKey() {
		t.Errorf("Strings with different content have same has keys")
	}
}
