package board

import (
	"fmt"
	"testing"
)

func TestRules(t *testing.T) {
	TestNakedSingles(t)

	t.Errorf("Add tests for remaining rules!")
}

func TestNakedSingles(t *testing.T) {
	inputBoard := []byte(nakedSinglesParam)
	b := NewBoard(inputBoard)
	finalBefore := b.finalValuesLeft
	b.NakedSingles()
	finalAfter := b.finalValuesLeft

	// Gauss-Seidel like update, gives more entries removed...
	if finalBefore != 53 || finalAfter != 46 {
		errMsg := fmt.Sprintf("NakedSingles fail (53,46): (%d,%d)", finalBefore, finalAfter)
		t.Errorf(errMsg)
	}
}
