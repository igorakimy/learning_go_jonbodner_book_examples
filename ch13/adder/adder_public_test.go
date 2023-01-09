package adder_test

import (
	"test_examples/adder"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	result := adder.AddNumber(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
