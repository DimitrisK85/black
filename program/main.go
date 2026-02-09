package main

import (
	"black"
	"fmt"
	"os"
)

func main() {
	valid_args := black.ArgCheck(os.Args)
	if !valid_args {
		return
	}
	grid := black.UnsolvedSudoku(os.Args, valid_args)
	if black.Solver(grid) {
		black.Print_Box(grid)
	} else {
		fmt.Println("This sudoku is unsolvable. RIP")
	}
}
