# Sudoku Solver

A command-line Sudoku solver written in Go that solves 9x9 Sudoku puzzles using a backtracking algorithm.

## Features

- Input Validation: Checks for valid Sudoku input format  
- Backtracking Algorithm: Efficiently solves Sudoku puzzles using recursion  
- Visual Output: Displays the solved Sudoku in an organized 3x3 grid format  
- Error Handling: Provides clear error messages for invalid inputs  

## Installation

Ensure you have Go installed on your system.

Clone or download the Sudoku solver files.

## Usage

Run the program with 9 arguments representing the Sudoku rows, where:

- Numbers 1-9 represent filled cells  
- Dots (.) represent empty cells  

## Syntax

### Input Format Rules

- Exactly 9 arguments must be provided  
- Each argument must be exactly 9 characters long  
- Valid characters: digits 1-9 and dots (.) for empty cells  
- Example valid row: `"1..7..2.4"`  

### Example

```bash
go run main.go "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79"
Expected Output
+-------+-------+-------+
| 5 3 4 | 6 7 8 | 9 1 2 |
| 6 7 2 | 1 9 5 | 3 4 8 |
| 1 9 8 | 3 4 2 | 5 6 7 |
+-------+-------+-------+
| 8 5 9 | 7 6 1 | 4 2 3 |
| 4 2 6 | 8 5 3 | 7 9 1 |
| 7 1 3 | 9 2 4 | 8 5 6 |
+-------+-------+-------+
| 9 6 1 | 5 3 7 | 2 8 4 |
| 2 8 7 | 4 1 9 | 6 3 5 |
| 3 4 5 | 2 8 6 | 1 7 9 |
+-------+-------+-------+
File Descriptions
main.go: Handles command-line arguments and orchestrates the solving process

sudoku/sudoku.go: Contains all Sudoku-related functions:

ArgCheck(): Validates input arguments

UnsolvedSudoku(): Parses input into a 2D grid

Solver(): Implements the backtracking algorithm

IsValid(): Checks if a number can be placed in a cell

Print_Box(): Displays the Sudoku grid

Algorithm
The solver uses a recursive backtracking algorithm:

Finds the next empty cell

Tries numbers 1-9 in that cell

Checks if the number is valid (no conflicts in row, column, or 3x3 box)

Recursively attempts to solve the rest of the puzzle

Backtracks if a dead end is reached

Error Messages
The program provides helpful error messages for:

Incorrect number of arguments

Invalid row length

Invalid characters in input

Unsolvable puzzles

Limitations
Only solves standard 9x9 Sudoku puzzles

Input must be provided via command-line arguments

No graphical user interface

Sudoku Solver - Examples
Valid Arguments (10 Examples)
Example 1: Easy Puzzle

go run main.go "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79"
Example 2: Medium Puzzle

go run main.go "..5.1.2.." "8....5.9." ".6.8.2.7." "9..4.3..1" ".2.....5." "4..9.6..2" ".3.1.9.8." ".8.4....3" "..7.6.4.."
Example 3: Hard Puzzle

go run main.go "8........" "..36....." ".7..9.2.." ".5...7..." "....457.." "...1...3." "..1....68" "..85...1." ".9....4.."
Example 4: Empty

go run main.go "........." "........." "........." "........." "........." "........." "........." "........." "........."
Example 5: Almost Full

go run main.go "12345678." "456789123" "789123456" "234567891" "567891234" "891234567" "345678912" "678912345" "912345678"
Example 6: Diagonal Pattern

go run main.go "1........" ".2......." "..3......" "...4....." "....5...." ".....6..." "......7.." ".......8." "........9"
Example 7: Center Focus

go run main.go "....5...." "........." "..7.3.8.." "...2.4..." ".6.....1." "...8.9..." "..1.6.5.." "........." "....7...."
Example 8: Border Pattern

go run main.go "123456789" "4.......6" "7.......3" "8.......2" "5.......5" "2.......8" "9.......1" "6.......4" "987654321"
Example 9: Checkerboard

go run main.go "1.2.3.4.5" ".6.7.8.9." "2.3.4.5.6" ".7.8.9.1." "3.4.5.6.7" ".8.9.1.2." "4.5.6.7.8" ".9.1.2.3." "5.6.7.8.9"
Example 10: Mixed Difficulty

go run main.go "..6..3.7." "8..5.1..9" ".4.7.6..2" "3.8...5.1" ".1.4.8.6." "7.5...9.4" "2..8.5.3." "5..9.2..8" ".3.1..4.."
Invalid Arguments (3 Examples)
Example 1: Wrong Number of Arguments

go run main.go "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1"
Error:

Error: Error: Invaid number of arguments. The number of arguments must be nine. Your number of inputs was 5.
Example 2: Invalid Row Length

go run main.go "53..7...." "6..195..." ".98....6" "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79"
Error:

Error: Error: Invaid arguments. Here is an example of a valid argument: '1..7..2.4'. Your argument was .98....6.
Example 3: Invalid Character

go run main.go "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..a9"
Error:

Error: Error: Invaid character in your argument. Your argument must only contain numbers from 1 to 9 and for empty spaces dots '.'. Your argument was a.