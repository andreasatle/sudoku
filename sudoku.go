package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"sudoku/board"
)

var possibleValues = flag.Bool("posval", false, "Print Possible Values")
var fileName = flag.String("f", "no-file", "Initial Sudoku (sdk-file)")

func main() {
	flag.Parse()
	b := board.NewBoard(*fileName)

	input := bufio.NewScanner(os.Stdin)

	for b.CountFinalValuesLeft() != 0 {
		prompt(b, input)()
	}
	fmt.Println(b.FinalValuesToString())
}

func prompt(b *board.Board, input *bufio.Scanner) func() {

	exitFunc := func() {
		os.Exit(0)
	}
	str := "==========================\n"
	str += b.FinalValuesToString()
	if *possibleValues {
		str += "--------------------------\n"
		str += b.PossibleValuesToString()
	}
	str += "===================\n"
	str += "Menu options: \n"
	str += "===================\n"
	str += "1) Hidden Singles\n"
	str += "2) Naked Singles\n"
	str += "3) Naked Pairs\n"
	str += "4) Pointing Pairs\n"
	str += "5) Hidden Pairs\n"
	str += "6) X-Wings\n"
	str += "0) Exit program\n"
	str += "===================\n"
	str += "Your coice: "
	fmt.Print(str)
	input.Scan()
	switch choice := input.Text(); choice {
	case "1":
		fmt.Println("Try hidden singles")
		return b.HiddenSingles
	case "2":
		fmt.Println("Try naked singles")
		return b.NakedSingles
	case "3":
		fmt.Println("Try naked pairs")
		return b.NakedPairs
	case "4":
		fmt.Println("Try pointing pairs")
		return b.PointingPairs
	case "5":
		fmt.Println("Try hidden pairs")
		return b.HiddenPairs
	case "6":
		fmt.Println("Try X-Wings")
		return b.XWings
	case "0":
		return exitFunc
	}
	return func() {}
}
