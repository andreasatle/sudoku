package board

import (
	"testing"
	"strconv"
)

func Test(t *testing.T) {
	TestBoard(t)
	TestRules(t)
}

func TestBoard(t *testing.T) {
	TestSizes(t)
	TestBlocks(t)
	TestSets(t)
}

func TestSizes(t *testing.T) {
	inputBoard := []byte(xwingsParam)
	b := NewBoard(inputBoard)
	if b.CountFinalValuesLeft() != 29 {
		t.Errorf("Wrong number of final values left (29): " + strconv.Itoa(b.CountFinalValuesLeft()))
	}

	if b.CountCandidates() != 77 {
		t.Errorf("Wrong number of possible values (77): " + strconv.Itoa(b.CountCandidates()))
	}
}

func TestBlocks(t *testing.T) {
	inputBoard := []byte(small6x6Param)
	b := NewBoard(inputBoard)

	var flag bool
	correctRows := [][]int{{0,1,2,3,4,5},{6,7,8,9,10,11},{12,13,14,15,16,17},{18,19,20,21,22,23},{24,25,26,27,28,29},{30,31,32,33,34,35}}
	correctCols := [][]int{{0,6,12,18,24,30},{1,7,13,19,25,31},{2,8,14,20,26,32},{3,9,15,21,27,33},{4,10,16,22,28,34},{5,11,17,23,29,35}}
	correctBoxs := [][]int{{0,1,2,6,7,8},{3,4,5,9,10,11},{12,13,14,18,19,20},{15,16,17,21,22,23},{24,25,26,30,31,32},{27,28,29,33,34,35}}
	flag = true
	for i, line := range b.rowPositions {
		for j, _ := range line {
			if b.rowPositions[i][j] != correctRows[i][j] ||
			   b.colPositions[i][j] != correctCols[i][j] ||
			   b.boxPositions[i][j] != correctBoxs[i][j] {
				flag = false
			}
		}
	}
	if !flag {
		t.Errorf("Wrong definition of rows")
	}
}

func TestSets(t *testing.T) {
	inputBoard := []byte(small6x6Param)
	b := NewBoard(inputBoard)

	if b.size != 6  {
		t.Errorf("Wrong size of board (6): "+strconv.Itoa(b.size))
	}

	if b.rowBlockSize != 3 {
		t.Errorf("Wrong number of row blocks in board (3): "+strconv.Itoa(b.rowBlockSize))
	}

	if b.colBlockSize != 2 {
		t.Errorf("Wrong number of column blocks in board (2): "+strconv.Itoa(b.colBlockSize))
	}
}

