// The package contains the board configuration
// and the current position with candidates.
package board

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/andreasatle/set"
	"strconv"
	"bytes"
)

// struct containing the parameters of the board
type Board struct {
	size            int           // Size of Sudoku (normally == 9)
	rowBlockSize    int           // Number of rows blocks
	colBlockSize    int           // Number of column blocks
	runes           set.RuneSet   // All values that appear in the game
	finalValue      []rune        // The final value for each position

	candidates      []set.RuneSet // Possible values on each position

	rowPositions    [][]int       // Board positions for each row
	colPositions    [][]int       // Board positions for each col
	boxPositions    [][]int       // Board positions for each box

	finalValuesLeft int           // Number of final points left to find
}

// Count total number of possible values
func (b *Board) CountCandidates() int {
	count := 0

	for pos := 0; pos < b.size*b.size; pos++ {
		count += b.candidates[pos].Size()
	}
	return count
}

// Count have many final positions are left
func (b *Board) CountFinalValuesLeft() int {
	return b.finalValuesLeft
}

// Check if the board has an invalid configuration
func (b *Board) CheckInvalid() bool {
	for pos := 0; pos < b.size*b.size; pos++ {
		if (b.finalValue[pos] == '-' && b.candidates[pos].Size() == 0) ||
		   (b.finalValue[pos] != '-' && b.candidates[pos].Size() != 0) {
			return true
		}
	}
	return false
}

// Create a new instance of Board
func NewBoard(fileName string) *Board {
	b := new(Board)
	b.read(fileName)
	return b
}

// Mapping from index to row
func (b *Board) PositionToRow(i int) int {
	return i / b.size
}

// Mapping from index to column
func (b *Board) PositionToCol(i int) int {
	return i % b.size
}

// Mapping from index to box
func (b *Board) PositionToBox(i int) int {
	row := b.PositionToRow(i)
	col := b.PositionToCol(i)
	return (row/b.colBlockSize)*b.colBlockSize + col/b.rowBlockSize
}

// Mapping from row, column to index
func (b *Board) RowColToPosition(row int, col int) int {
	return row*b.size + col
}

// Write the final values to a string
func (b *Board) FinalValuesToString() string {

	var buf bytes.Buffer
	buf.WriteString("Current values on board:")

	for row := 0; row < b.size; row++ {
		buf.WriteString("\n")
		for col := 0; col < b.size; col++ {
			pos := b.RowColToPosition(row, col)
			buf.WriteString(string(b.finalValue[pos]))
		}
	}

	return buf.String()
}

// Write the candidates to a string
func (b *Board) CandidatesToString() string {
	var buf bytes.Buffer
	buf.WriteString("Entries left on board: ")
	buf.WriteString(strconv.Itoa(b.finalValuesLeft))
	buf.WriteString("\nNumber of possible values on board: ")
	buf.WriteString(strconv.Itoa(b.CountCandidates()))
	buf.WriteString("\nCurrent possible values on board:\n")

	for pos := 0; pos < b.size*b.size; pos++ {
		if b.candidates[pos].Empty() {
			continue
		}
		row := b.PositionToRow(pos)
		col := b.PositionToCol(pos)
		buf.WriteString("(")
		buf.WriteString(strconv.Itoa(row+1))
		buf.WriteString(",")
		buf.WriteString(strconv.Itoa(col+1))
		buf.WriteString("): ")
		buf.WriteString(b.candidates[pos].ToString())
		buf.WriteString("\n")
	}
	return buf.String()
}

// Convenient error check
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Read "header" of file containing the initial sudoku.
// First line contains size, rowbox, colbox.
// Second line contains the characters that are used, e.g. -123456789.
// Note that the first character is the value for not set.
// The third line is blank.
func (b *Board) readHeader(scanner *bufio.Scanner) {
	// Read first line of file
	scanner.Scan()
	_, err := fmt.Sscanf(scanner.Text(), "%d %d %d", &b.size, &b.rowBlockSize, &b.colBlockSize)
	check(err)

	// Allocate memory for struct

	scanner.Scan()
	line := []rune(scanner.Text())

	b.runes.Clear()
	for _, r := range line {
		b.runes.Add(r)
	}

	// Read extra newline
	scanner.Scan()
}

// Help routine taking care of the initialization.
func (b *Board) initializeGeneral() {

	b.rowPositions = make([][]int, b.size)
	b.colPositions = make([][]int, b.size)
	b.boxPositions = make([][]int, b.size)

	b.finalValue = make([]rune, b.size*b.size)
	b.candidates = make([]set.RuneSet, b.size*b.size)
	b.finalValuesLeft = b.size*b.size

	// Read second line with valid characters
	for pos := 0; pos < b.size*b.size; pos++ {
		b.candidates[pos].Copy(b.runes)
	}

	for i := 0; i < b.size; i++ {
		b.rowPositions[i] = make([]int, 0, b.size)
		b.colPositions[i] = make([]int, 0, b.size)
		b.boxPositions[i] = make([]int, 0, b.size)
	}

	// Initialize the board positions for rows, columns and boxs
	for pos := 0; pos < b.size*b.size; pos++ {
		row := b.PositionToRow(pos)
		col := b.PositionToCol(pos)
		box := b.PositionToBox(pos)

		b.rowPositions[row] = append(b.rowPositions[row], pos)
		b.colPositions[col] = append(b.colPositions[col], pos)
		b.boxPositions[box] = append(b.boxPositions[box], pos)
	}
}

// Read the initial board from input-file.
// The board is a b.size x b.size grid of characters,
// that should have been read by the function readHeader.
func (b *Board) readBoard(scanner *bufio.Scanner) {

	for row := 0; row < b.size; row++ {
		scanner.Scan()
		line := scanner.Text()
		for col := 0; col < b.size; col++ {

			r := rune(line[col])

			pos := b.RowColToPosition(row, col)
			if b.runes.Contains(r) {
				b.SetPosition(pos, r)
			} else {
				b.finalValue[pos] = '-'
			}
		}
	}
}

// Read initial sudoku board and set up the Board struct
func (b *Board) read(fileName string) {
	fid, err := os.Open(fileName)

	check(err)

	scanner := bufio.NewScanner(fid)
	scanner.Split(bufio.ScanLines)

	b.readHeader(scanner)
	b.initializeGeneral()
	b.readBoard(scanner)
}

// Set a position of the board to a value indexToRune[index]
func (b *Board) SetPosition(pos int, r rune) {

	b.finalValue[pos] = r
	b.finalValuesLeft--

	row := b.PositionToRow(pos)
	col := b.PositionToCol(pos)
	box := b.PositionToBox(pos)

	b.candidates[pos].Clear()

	for _, i := range b.rowPositions[row] {
		b.candidates[i].Remove(r)
	}

	for _, i := range b.colPositions[col] {
		b.candidates[i].Remove(r)
	}

	for _, i := range b.boxPositions[box] {
		b.candidates[i].Remove(r)
	}
}
