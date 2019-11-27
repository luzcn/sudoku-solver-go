package solver

import (
	"fmt"
	"log"
)

// ["53..7....","6..195...", ".98....6.","8...6...3", "4..8.3..1","7...2...6", ".6....28.","...419..5", "....8..79"]
//board := [][]byte{
//{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
//{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
//{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
//{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
//{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
//{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
//{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
//{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
//{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

func Solver(board [][] byte) bool {

	log.Println(fmt.Sprintf("%s", board))
	return solve(board, 0, 0)
}

func solve(board [][]byte, row int, column int) bool {
	if column == 9 {
		return solve(board, row+1, 0)
	}

	// the length should be 9
	if row == 9 {
		return true
	}

	if board[row][column] != '.' {
		return solve(board, row, column+1)
	}

	for _, v := range []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'} {
		if isValid(board, row, column, v) {
			board[row][column] = v
			if solve(board, row, column+1) {
				return true
			}
			board[row][column] = '.'
		}
	}
	return false
}

func isValid(board [][]byte, x int, y int, v byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][y] == v || board[x][i] == v {
			return false
		}
	}

	// check 3*3 square
	squareX := (x / 3) * 3
	squareY := (y / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[squareX+i][squareY+j] == v {
				return false
			}
		}
	}

	return true
}
