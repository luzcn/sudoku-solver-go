package main

//const N = 9

func Solver(board [][]int) {
	solve(board, 0, 0)
}

func solve(board [][]int, i int, j int) bool {
	// the length should be 9
	if i >= 9 {
		return true
	}

	if j >= 9 {
		return solve(board, i+1, 0)
	}

	for v := 1; v <= 9; v++ {
		if isValid(board, i, j, v) {
			board[i][j] = v
			if solve(board, i, j+1) {
				return true
			}

			board[i][j] = 0
		}
	}

	return false
}

func isValid(board [][]int, x int, y int, v int) bool {

	var m int = len(board)
	var n int = len(board[0])

	// check row
	for i := 0; i < m; i++ {
		if board[i][y] == v {
			return false;
		}
	}

	// check column
	for j := 0; j < n; j++ {
		if board[x][j] == v {
			return false;
		}
	}

	// check 3*3 squre
	squareX := (x / 3) * 3;
	squareY := (y / 3) * 3;
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if board[i][y] != 0 && board[squareX+i][squareY+j] == v {
				return false;
			}
		}
	}

	return true
}

// ["9,3,4,2,0,1,6,0,7","2,0,6,9,0,7,4,3,1","4,6,3,9,0,2 0 0 0","5,0,0,3,9,6,8,7,4","4,9,7,5,2,8,3,1,6","3,1,7,4,5,2,9 0 0","3,9,7,4,5 0 0 0 0","5,0,4,2,1,9,3 0 0","3,4,9,7,1,5,2,6,8"]
