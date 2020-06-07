package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"io/ioutil"
	"bytes"
	"github.com/andreasatle/sudoku/board"
	"log"
)

var candidates = flag.Bool("c", false, "Print Candidates")
var fileName = flag.String("f", "no-file", "Initial Sudoku (sdk-file)")

func main() {
	flag.Parse()

	inputBoard, err := ioutil.ReadFile(*fileName)

	//fid, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	b := board.NewBoard(inputBoard)

	input := bufio.NewScanner(os.Stdin)

	for b.CountFinalValuesLeft() != 0 {
		prompt(b, input)()
		if b.CheckInvalid() {
			panic("Sudoku is invalid")
		}
	}
	fmt.Println(b.FinalValuesToString())
}

func prompt(b *board.Board, input *bufio.Scanner) func() {

	exitFunc := func() {
		os.Exit(0)
	}
	var buf bytes.Buffer
	buf.WriteString("\n==========================\n")
	buf.WriteString(b.FinalValuesToString())
	if *candidates {
		buf.WriteString("\n--------------------------\n")
		buf.WriteString(b.CandidatesToString())
	}
	buf.WriteString("\n===================\n")
	buf.WriteString("Menu options: \n")
	buf.WriteString("===================\n")
	buf.WriteString("1) Hidden Singles\n")
	buf.WriteString("2) Naked Singles\n")
	buf.WriteString("3) Naked Pairs\n")
	buf.WriteString("4) Pointing Pairs\n")
	buf.WriteString("5) Hidden Pairs\n")
	buf.WriteString("6) X-Wings\n")
	buf.WriteString("0) Exit program\n")
	buf.WriteString("===================\n")
	buf.WriteString("Your coice: ")
	fmt.Print(buf.String())
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

