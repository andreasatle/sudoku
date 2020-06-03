package board

import "testing"

func Test(t *testing.T) {
	n := 9
	b := new(Board)

	if len(b.possibleValues) != n*n {
		t.Errorf("Init Test failed, %d.", len(b.possibleValues))
	}

	for i, e := range b.boardSet {
		if e.Len() != n {
			t.Errorf("Wrong set length %d for index %d", e.Len(), i)
		}
	}
}
