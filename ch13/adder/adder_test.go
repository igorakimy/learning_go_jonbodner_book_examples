package adder

import "testing"

func Test_addNumber(t *testing.T) {
	result := addNumber(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
