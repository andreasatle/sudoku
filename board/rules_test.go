package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test that two '8' are added to the final values for naked singles
func TestNakedSingles(t *testing.T) {
	inputBoard := []byte(nakedSinglesParam)
	b := NewBoard(inputBoard)

	assert.NotEqual(t, '8', b.finalValue[5])
	assert.NotEqual(t, '8', b.finalValue[11])

	b.NakedSingles()

	assert.Equal(t, '8', b.finalValue[5])
	assert.Equal(t, '8', b.finalValue[11])
}

// Test that three final values {1,4,7} are set for hidden singles
func TestHiddenSingles(t *testing.T) {
	inputBoard := []byte(hiddenSinglesParam)
	b := NewBoard(inputBoard)

	assert.NotEqual(t, '7', b.finalValue[2])
	assert.NotEqual(t, '4', b.finalValue[13])
	assert.NotEqual(t, '1', b.finalValue[17])

	b.HiddenSingles()

	assert.Equal(t, '7', b.finalValue[2])
	assert.Equal(t, '4', b.finalValue[13])
	assert.Equal(t, '1', b.finalValue[17])
}

// Test that 4 candidates are removed for naked pairs
func TestNakedPairs(t *testing.T) {
	inputBoard := []byte(nakedPairsParam)
	b := NewBoard(inputBoard)

	assert.True(t, b.candidates[64].Contains('C'))
	assert.True(t, b.candidates[73].Contains('G'))
	assert.True(t, b.candidates[78].Contains('B'))
	assert.True(t, b.candidates[78].Contains('G'))

	b.NakedPairs()

	assert.True(t, !b.candidates[64].Contains('C'))
	assert.True(t, !b.candidates[73].Contains('G'))
	assert.True(t, !b.candidates[78].Contains('B'))
	assert.True(t, !b.candidates[78].Contains('G'))
}

// Test that 18 candidates are removed for pointing pairs
func TestPointingPairs(t *testing.T) {
	inputBoard := []byte(pointingPairsParam)
	b := NewBoard(inputBoard)

	assert.True(t, b.candidates[18].Contains('1'))
	assert.True(t, b.candidates[20].Contains('1'))

	assert.True(t, b.candidates[32].Contains('2'))
	assert.True(t, b.candidates[41].Contains('2'))
	assert.True(t, b.candidates[50].Contains('2'))
	assert.True(t, b.candidates[68].Contains('2'))

	assert.True(t, b.candidates[32].Contains('3'))
	assert.True(t, b.candidates[41].Contains('3'))
	assert.True(t, b.candidates[50].Contains('3'))
	assert.True(t, b.candidates[68].Contains('3'))

	assert.True(t, b.candidates[36].Contains('4'))
	assert.True(t, b.candidates[37].Contains('4'))

	assert.True(t, b.candidates[39].Contains('6'))

	assert.True(t, b.candidates[71].Contains('7'))

	assert.True(t, b.candidates[69].Contains('8'))
	assert.True(t, b.candidates[70].Contains('8'))
	assert.True(t, b.candidates[71].Contains('8'))

	assert.True(t, b.candidates[57].Contains('9'))

	b.PointingPairs()

	assert.True(t, !b.candidates[18].Contains('1'))
	assert.True(t, !b.candidates[20].Contains('1'))

	assert.True(t, !b.candidates[32].Contains('2'))
	assert.True(t, !b.candidates[41].Contains('2'))
	assert.True(t, !b.candidates[50].Contains('2'))
	assert.True(t, !b.candidates[68].Contains('2'))

	assert.True(t, !b.candidates[32].Contains('3'))
	assert.True(t, !b.candidates[41].Contains('3'))
	assert.True(t, !b.candidates[50].Contains('3'))
	assert.True(t, !b.candidates[68].Contains('3'))

	assert.True(t, !b.candidates[36].Contains('4'))
	assert.True(t, !b.candidates[37].Contains('4'))

	assert.True(t, !b.candidates[39].Contains('6'))

	assert.True(t, !b.candidates[71].Contains('7'))

	assert.True(t, !b.candidates[69].Contains('8'))
	assert.True(t, !b.candidates[70].Contains('8'))
	assert.True(t, !b.candidates[71].Contains('8'))

	assert.True(t, !b.candidates[57].Contains('9'))
}

// Test that 6 candidates are removed for X-Wings
func TestXWings(t *testing.T) {
	inputBoard := []byte(xWingsParam)
	b := NewBoard(inputBoard)

	assert.True(t, b.candidates[1].Contains('4'))
	assert.True(t, b.candidates[4].Contains('4'))
	assert.True(t, b.candidates[10].Contains('4'))
	assert.True(t, b.candidates[13].Contains('4'))
	assert.True(t, b.candidates[46].Contains('4'))
	assert.True(t, b.candidates[49].Contains('4'))

	b.XWings()

	assert.True(t, !b.candidates[1].Contains('4'))
	assert.True(t, !b.candidates[4].Contains('4'))
	assert.True(t, !b.candidates[10].Contains('4'))
	assert.True(t, !b.candidates[13].Contains('4'))
	assert.True(t, !b.candidates[46].Contains('4'))
	assert.True(t, !b.candidates[49].Contains('4'))
}

// Test that 1 candidate is removed for hidden pairs
func TestHiddenPairs(t *testing.T) {
	inputBoard := []byte(hiddenPairsParam)
	b := NewBoard(inputBoard)

	assert.True(t, b.candidates[44].Contains('6'))

	b.HiddenPairs()

	assert.True(t, !b.candidates[44].Contains('6'))
}
