package black

import (
	"fmt"
)

func ArgCheck(args []string) bool {
	if len(args) != 10 { // Checking if the number of arguments is valid.
		length := len(args) - 1
		fmt.Printf("Error: Invaid number of arguments. The number of arguments must be nine. Your number of inputs was %d.\n", length)
		return false
	}
	for i := 1; i < len(args); i++ { // Checking if the size of the arguments is valid.
		if len(args[i]) != 9 {
			invalid := args[i]
			fmt.Printf("Error: Invaid arguments. Here is an example of a valid argument: '1..7..2.4'. Your argument was '%s'.\n", invalid)
			return false
		} else {
			for _, char := range args[i] { // Checking if the elements of the arguments are valid.
				if !(char > '0' && char <= '9') && char != '.' {
					invalid := char
					fmt.Printf("Error: Invaid character in your argument. Your argument must only contain numbers from 1 to 9 and for empty spaces dots '.'. Your argument was '%s'.\n", string(invalid))
					return false
				}
			}
		}
	}
	return true
}

func UnsolvedSudoku(args []string, valid_args bool) [][]int {
	grid := make([][]int, 9)
	if ArgCheck(args) {
		for i := 1; i < 10; i++ { // Since the arguments are valid, we are importing the arguments into the grid.
			row := make([]int, 9)
			for j, char := range args[i] {
				if char == '.' { // We changed dots to 0, which represent empty cells.
					row[j] = 0
				} else {
					row[j] = int(char - '0')
				}
			}
			grid[i-1] = row
		}
	}
	return grid
}

func Solver(grid [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if IsValid(grid, row, col, num) == true { // Calling the valid function to check if the number is valid to be a candidate.
						grid[row][col] = num
						if Solver(grid) { // Recursive to check if the number was indeed the right one, or blocks the rest of the sudoku.
							return true
						}
						grid[row][col] = 0 // Backtrack to check other valid numbers which were candidates.
					}
				}
				return false
			}
		}
	}
	return true
}

func IsValid(grid [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ { // Checking if there is a number in the same row twice.
		if grid[row][i] == num {
			return false
		}
	}
	for i := 0; i < 9; i++ { // Checking if there is a number in the same column twice.
		if grid[i][col] == num {
			return false
		}
	}
	startrow := (row / 3) * 3 // Check the 3x3 box.
	startcol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startrow+i][startcol+j] == num {
				return false
			}
		}
	}
	return true
}

func Print_Box(grid [][]int) { // Printing a nice box for the sudoku with the solution.
	fmt.Println("+-------+-------+-------+")
	for i := 0; i < 9; i++ {
		fmt.Print("| ")
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", grid[i][j])
			if j == 2 || j == 5 { // Printing the vertical lines for the 3x3 box.
				fmt.Print("| ")
			}
		}
		fmt.Println("|")

		if i == 2 || i == 5 { // Printing the horizontal lines for the 3x3 box.
			fmt.Println("+-------+-------+-------+")
		}
	}
	fmt.Println("+-------+-------+-------+")
}
