// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

package main

func main() {

	board := [][]byte{
		[]byte{'.', '.', '.', '9', '.', '.', '.', '.', '.'},
		[]byte{'.', '3', '.', '.', '.', '.', '.', '.', '1'},
		[]byte{'1', '.', '.', '.', '.', '.', '3', '.', '.'},
		[]byte{'.', '9', '.', '.', '.', '.', '.', '7', '.'},
		[]byte{'8', '.', '.', '8', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '2', '.', '6', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	println(isValidSudoku(board)) // false
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !checkRow(board, i) {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if !checkColumn(board, i) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !checkBox(board, i, j) {
				return false
			}
		}
	}

	return true

}

func checkRow(board [][]byte, row int) bool {
	seen := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		if board[row][i] != '.' {
			if seen[board[row][i]] {
				return false
			}
			seen[board[row][i]] = true
		}
	}
	return true
}

func checkColumn(board [][]byte, column int) bool {
	seen := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		if board[i][column] != '.' {
			if seen[board[i][column]] {
				return false
			}
			seen[board[i][column]] = true
		}
	}
	return true
}

func checkBox(board [][]byte, row, column int) bool {
	seen := make(map[byte]bool)
	for i := row; i < row+3; i++ {
		for j := column; j < column+3; j++ {
			if board[i][j] != '.' {
				if seen[board[i][j]] {
					return false
				}
				seen[board[i][j]] = true
			}
		}
	}
	return true
}
