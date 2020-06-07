package board

import (
	"github.com/andreasatle/set"
)

// Count the occurancies of each value in the block
func (b *Board) countValues(positions []int) map[rune]int {
	counter := make(map[rune]int)
	for _, pos := range positions {
		for _, r := range b.candidates[pos].ToSlice() {
			counter[r]++
		}
	}
	return counter
}

// Filter out all values that do not occur exactly once
func filterCounter(counter map[rune]int, num int) {
	for key, value := range counter {
		if value != num {
			delete(counter, key)
		}
	}
}

// When only one possible value, move it to final value
func (b *Board) NakedSingles() {
	for i := 0; i < b.size*b.size; i++ {
		if b.candidates[i].Size() == 1 {
			r := b.candidates[i].ToSlice()[0]
			b.SetPosition(i, r)
		}
	}
}

// When a value is possible at only one position in a block,
// then that value is the final value.
func (b *Board) HiddenSingles() {

	// Set the final position for the remaining values
	setRemainingValues := func(positions []int, counter map[rune]int) {
		for key := range counter {
			for _, pos := range positions {
				if b.candidates[pos].Contains(key) {
					b.SetPosition(pos, key)
				}
			}
		}
	}

	// One case of hidden singles (row, column or box)
	hiddenSinglesBlock := func(positions []int) {

		counter := b.countValues(positions)
		filterCounter(counter, 1)
		setRemainingValues(positions, counter)

	}

	// Loop over all possible rows, columns and boxes
	for i := 0; i < b.size; i++ {
		hiddenSinglesBlock(b.rowPositions[i])
		hiddenSinglesBlock(b.colPositions[i])
		hiddenSinglesBlock(b.boxPositions[i])
	}
}

func (b *Board) NakedPairs() {

	// Find all positions in the block with two possible values
	tuplePositions := func(positions []int) *set.IntSet {
		tuplePos := set.NewIntSet()
		for _, pos := range positions {
			if b.candidates[pos].Size() == 2 {
				tuplePos.Add(pos)
			}
		}
		return tuplePos
	}

	// Keep all tuples that appears twice and its values
	tuplePairs := func(tuplePos *set.IntSet) (*set.IntSet, *set.RuneSet) {

		tuplePairs := set.NewIntSet()
		values := set.NewRuneSet()

		// Loop in upper triangular pattern detect pairs of equal tuples
		for _, posI := range tuplePos.ToSlice() {
			for _, posJ := range tuplePos.ToSlice() {
				pVposI := &b.candidates[posI]
				pVposJ := &b.candidates[posJ]
				if posI < posJ && pVposI.Equal(pVposJ) {
					tuplePairs.Add(posI)
					tuplePairs.Add(posJ)
					values = values.Union(pVposI).Union(pVposJ)
				}
			}
		}
		return tuplePairs, values
	}

	// Remove the values that appears in tuplePairs (from the other positions)
	removeValues := func(positions []int, tuplePairs *set.IntSet, values *set.RuneSet) {
		positionSet := set.FromIntSlice(positions).Difference(tuplePairs)
		for _, pos := range positionSet.ToSlice() {
			posVal := &b.candidates[pos]
			for _, value := range values.ToSlice() {
				posVal.Remove(value)
			}
		}
	}

	// One case of naked pairs (row, column or box)
	nakedPairsBlock := func(positions []int) {
		tuplePos := tuplePositions(positions)
		tuplePairs, values := tuplePairs(tuplePos)
		removeValues(positions, tuplePairs, values)
	}

	for i := 0; i < b.size; i++ {
		nakedPairsBlock(b.rowPositions[i])
		nakedPairsBlock(b.colPositions[i])
		nakedPairsBlock(b.boxPositions[i])
	}
}

func (b *Board) PointingPairs() {

	// Create three disjoint sets of positions
	createPosSubsets := func(positions1 []int, positions2 []int) (*set.IntSet, *set.IntSet, *set.IntSet) {
		posSet1 := set.FromIntSlice(positions1)
		posSet2 := set.FromIntSlice(positions2)
		posIsect := posSet1.Intersection(posSet2)
		return posIsect, posSet1.Difference(posIsect), posSet2.Difference(posIsect)
	}

	// Create the set with all possible values of the positions in posSet
	createValSet := func(posSet *set.IntSet) *set.RuneSet {
		posSlice := posSet.ToSlice()
		runeSet := set.NewRuneSet()
		for _, pos := range posSlice {
			runeSet = runeSet.Union(&b.candidates[pos])
		}
		return runeSet
	}

	// Filter the values in valIsec-valSet from the values in positions in posSet
	filterVals := func(valIsect *set.RuneSet, valSet *set.RuneSet, posSet *set.IntSet) {
		vals := valIsect.Difference(valSet)
		for _, key := range vals.ToSlice() {
			for _, pos := range posSet.ToSlice() {
				b.candidates[pos].Remove(key)
			}
		}
	}

	// One box and one row or column
	pointingPairsBox := func(positions1 []int, positions2 []int) {

		posIsect, posSubset1, posSubset2 := createPosSubsets(positions1, positions2)
		if posIsect.Empty() {
			return
		}
		valIsectSet := createValSet(posIsect)
		valSet1 := createValSet(posSubset1)
		valSet2 := createValSet(posSubset2)
		filterVals(valIsectSet, valSet1, posSubset2)
		filterVals(valIsectSet, valSet2, posSubset1)
	}

	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			pointingPairsBox(b.boxPositions[i], b.rowPositions[j])
			pointingPairsBox(b.boxPositions[i], b.colPositions[j])
		}
	}
}

// Same as pointing pair (or am I wrong?!)
func (b *Board) ClaimingPairs() {
	b.PointingPairs()
}

// This is broke!
func (b *Board) XWings() {

	createIsect := func(counterI, counterJ map[rune]int) *set.RuneSet {
		setI := set.NewRuneSet()
		setJ := set.NewRuneSet()
		for key := range counterI {
			setI.Add(key)
		}
		for key := range counterJ {
			setJ.Add(key)
		}
		return setI.Intersection(setJ)
	}

	createRowColSet := func(value rune, positions []int, rowSet, colSet *set.IntSet) {
		for _, pos := range positions {
			if b.candidates[pos].Contains(value) {
				rowSet.Add(b.PositionToRow(pos))
				colSet.Add(b.PositionToCol(pos))
			}
		}
	}
	// Given that value is found twice in both positionsI and positionsJ
	filterValues := func(value rune, positionsI, positionsJ []int, isRow bool) {
		rowSet := set.NewIntSet()
		colSet := set.NewIntSet()
		createRowColSet(value, positionsI, rowSet, colSet)
		createRowColSet(value, positionsJ, rowSet, colSet)
		if rowSet.Size() != 2 || colSet.Size() != 2 {
			return
		}
		posSet := set.NewIntSet()
		if isRow {
			for _, i := range colSet.ToSlice() {
				posSet = posSet.Union(set.FromIntSlice(b.colPositions[i]))
			}
			for _, i := range rowSet.ToSlice() {
				posSet = posSet.Difference(set.FromIntSlice(b.rowPositions[i]))
			}
		} else {
			for _, i := range rowSet.ToSlice() {
				posSet = posSet.Union(set.FromIntSlice(b.rowPositions[i]))
			}
			for _, i := range colSet.ToSlice() {
				posSet = posSet.Difference(set.FromIntSlice(b.colPositions[i]))
			}
		}

		for _, pos := range posSet.ToSlice() {
			b.candidates[pos].Remove(value)
		}
	}

	xWingsBlock := func(i int, j int, isRow bool) {
		var positionsI, positionsJ []int

		// Get the board positions for either row or column i and j
		if isRow { // i and j are row indices; k, l are column indices
			positionsI = b.rowPositions[i]
			positionsJ = b.rowPositions[j]
		} else { // i and j are column indices; k, l are row indices
			positionsI = b.colPositions[i]
			positionsJ = b.colPositions[j]
		}

		// Count occurrencies of each value in the blocks
		counterI := b.countValues(positionsI)
		counterJ := b.countValues(positionsJ)

		// Only keep values that appears exactly once in block I
		filterCounter(counterI, 2)
		filterCounter(counterJ, 2)

		valueCandidate := createIsect(counterI, counterJ)

		if valueCandidate.Empty() {
			return
		}

		for _, value := range valueCandidate.ToSlice() {
			filterValues(value, positionsI, positionsJ, isRow)
		}
	}

	for i := 0; i < b.size; i++ {
		for j := i + 1; j < b.size; j++ {
			xWingsBlock(i, j, true)  // Investigate rows first
			xWingsBlock(i, j, false) // Investigate columns first
		}
	}
}

func (b *Board) HiddenPairs() {
	getValues := func(positions []int) []rune {
		counter := b.countValues(positions)
		filterCounter(counter, 2)
		values := make([]rune, 0, len(counter))
		for key := range counter {
			values = append(values, key)
		}
		return values
	}
	processPair := func(valI, valJ rune, positions []int) {
		posSet := set.NewIntSet()
		for _, pos := range positions {
			if b.candidates[pos].Contains(valI) && b.candidates[pos].Contains(valJ) {
				posSet.Add(pos)
			}
		}
		if posSet.Size() == 2 {
			for _, pos := range posSet.ToSlice() {
				b.candidates[pos].Clear()
				b.candidates[pos].Add(valI)
				b.candidates[pos].Add(valJ)
			}
		}
		//fmt.Println(valI,valJ,posSet.ToString())

	}
	hiddenPairsBlock := func(positions []int) {
		values := getValues(positions)
		if len(values) < 2 {
			return
		}
		for i, valI := range values {
			for j, valJ := range values {
				if i >= j { // Avoid duplicates
					continue
				}
				processPair(valI, valJ, positions)
			}
		}
	}

	for i := 0; i < b.size; i++ {
		hiddenPairsBlock(b.rowPositions[i])
		hiddenPairsBlock(b.colPositions[i])
		hiddenPairsBlock(b.boxPositions[i])
	}
}

func (b *Board) NakedTriplets() {
	panic("NakeTriplets Not Implemented")
}

func (b *Board) NakedQuads() {
	panic("NakedQuads Not Implemented")
}

func (b *Board) Coloring() {
	panic("Coloring Not Implemented")
}

func (b *Board) BruteForce() {
	panic("BruteForce Not Implemented")
}
