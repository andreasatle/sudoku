package board

import (
	"fmt"
	"testing"
)

func TestRules(t *testing.T) {
	TestNakedSingles(t)
	TestHiddenSingles(t)
	TestNakedPairs(t)
	TestPointingPairs(t)
	TestXWings(t)
	TestHiddenPairs(t)
}

func TestNakedSingles(t *testing.T) {
	inputBoard := []byte(nakedSinglesParam)
	b := NewBoard(inputBoard)
	finalBefore := b.finalValuesLeft
	b.NakedSingles()
	finalAfter := b.finalValuesLeft

	// Gauss-Seidel like update, gives more entries removed...
	if finalBefore != 53 || finalAfter != 46 {
		errMsg := fmt.Sprintf("NakedSingles fail (Before:53,After:46): (%d,%d)", finalBefore, finalAfter)
		t.Errorf(errMsg)
	}
}

func TestHiddenSingles(t *testing.T) {
	inputBoard := []byte(hiddenSinglesParam)
	b := NewBoard(inputBoard)
	finalBefore := b.finalValuesLeft
	b.HiddenSingles()
	finalAfter := b.finalValuesLeft

	// Gauss-Seidel like update, gives more entries removed...
	if finalBefore != 53 || finalAfter != 42 {
		errMsg := fmt.Sprintf("HiddenSingles fail (Before:53,After:42): (%d,%d)", finalBefore, finalAfter)
		t.Errorf(errMsg)
	}
}

func TestNakedPairs(t *testing.T) {
	inputBoard := []byte(nakedPairsParam)
	b := NewBoard(inputBoard)
	candidatesBefore := b.CountCandidates()
	b.NakedPairs()
	candidatesAfter := b.CountCandidates()

	// Gauss-Seidel like update, gives more entries removed...
	if candidatesBefore != 59 || candidatesAfter != 55 {
		errMsg := fmt.Sprintf("NakedPairs fail (Before:57,After:55): (%d,%d)", candidatesBefore, candidatesAfter)
		t.Errorf(errMsg)
	}
}

func TestPointingPairs(t *testing.T) {
	inputBoard := []byte(pointingPairsParam)
	b := NewBoard(inputBoard)
	candidatesBefore := b.CountCandidates()
	b.PointingPairs()
	candidatesAfter := b.CountCandidates()

	// Gauss-Seidel like update, gives more entries removed...
	if candidatesBefore != 198 || candidatesAfter != 180 {
		errMsg := fmt.Sprintf("PointingPairs fail (Before:198,After:180): (%d,%d)", candidatesBefore, candidatesAfter)
		t.Errorf(errMsg)
	}
}


func TestXWings(t *testing.T) {
	inputBoard := []byte(xWingsParam)
	b := NewBoard(inputBoard)
	candidatesBefore := b.CountCandidates()
	b.XWings()
	candidatesAfter := b.CountCandidates()

	// Gauss-Seidel like update, gives more entries removed...
	if candidatesBefore != 77 || candidatesAfter != 71 {
		errMsg := fmt.Sprintf("XWings fail (Before:77,After:71): (%d,%d)", candidatesBefore, candidatesAfter)
		t.Errorf(errMsg)
	}
}

func TestHiddenPairs(t *testing.T) {
	inputBoard := []byte(hiddenPairsParam)
	b := NewBoard(inputBoard)
	candidatesBefore := b.CountCandidates()
	b.HiddenPairs()
	candidatesAfter := b.CountCandidates()

	// Gauss-Seidel like update, gives more entries removed...
	if candidatesBefore != 68 || candidatesAfter != 67 {
		errMsg := fmt.Sprintf("HiddenPairs fail (Before:68,After:67): (%d,%d)", candidatesBefore, candidatesAfter)
		t.Errorf(errMsg)
	}
}
