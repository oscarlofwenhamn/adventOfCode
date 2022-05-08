package main

import (
	"testing"

	"github.com/oscarlofwenhamn/adventOfCode/cmd/utils"
)

// TestGenerateDiagram validates that createDiagram correctly generates the expected diagram based on test input
func TestGenerateDiagram(t *testing.T) {

	inputArray := utils.ReadStringArrayFromFile("data/test_input")
	outputArray := utils.ReadStringArrayFromFile("data/test_output")

	diagram := createDiagram(inputArray)

	for i, line := range outputArray {
		if diagram[i] != line {
			t.Fatalf("Diagram line %d does not match desired output.\nExpected:\n%v\nActual:\n%v",
				i+1,
				line,
				diagram[i])

		}

	}
}
