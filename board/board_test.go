package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSizes(t *testing.T) {
	inputBoard := []byte(xwingsParam)
	b := NewBoard(inputBoard)
	assert.Equal(t, 29, b.CountFinalValuesLeft(), "Number of final values left.")

	assert.Equal(t, 77, b.CountCandidates(), "Number of possible values.")
}

func TestBlocks(t *testing.T) {
	inputBoard := []byte(small6x6Param)
	b := NewBoard(inputBoard)

	var flag bool
	correctRows := [][]int{{0, 1, 2, 3, 4, 5}, {6, 7, 8, 9, 10, 11}, {12, 13, 14, 15, 16, 17}, {18, 19, 20, 21, 22, 23}, {24, 25, 26, 27, 28, 29}, {30, 31, 32, 33, 34, 35}}
	correctCols := [][]int{{0, 6, 12, 18, 24, 30}, {1, 7, 13, 19, 25, 31}, {2, 8, 14, 20, 26, 32}, {3, 9, 15, 21, 27, 33}, {4, 10, 16, 22, 28, 34}, {5, 11, 17, 23, 29, 35}}
	correctBoxs := [][]int{{0, 1, 2, 6, 7, 8}, {3, 4, 5, 9, 10, 11}, {12, 13, 14, 18, 19, 20}, {15, 16, 17, 21, 22, 23}, {24, 25, 26, 30, 31, 32}, {27, 28, 29, 33, 34, 35}}
	flag = true
	for i, line := range b.rowPositions {
		for j := range line {
			if b.rowPositions[i][j] != correctRows[i][j] ||
				b.colPositions[i][j] != correctCols[i][j] ||
				b.boxPositions[i][j] != correctBoxs[i][j] {
				flag = false
			}
		}
	}
	assert.True(t, flag, "Definition of rows")
}

func TestSets(t *testing.T) {
	inputBoard := []byte(small6x6Param)
	b := NewBoard(inputBoard)

	assert.Equal(t, 6, b.size, "Size of board.")
	assert.Equal(t, 3, b.rowBlockSize, "Size of row block.")
	assert.Equal(t, 2, b.colBlockSize, "Size of column block.")
}
