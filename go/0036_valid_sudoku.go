package main

//time: O(1)
//space O(1)

func isValidSudoku(board [][]byte) bool {
	var rows, cols, squares [9]map[byte]bool

	for i := 0; i < 9; i++ {
		rows[i] = map[byte]bool{}
		cols[i] = map[byte]bool{}
		squares[i] = map[byte]bool{}
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			field := board[r][c]
			square := r/3*3 + c/3
			if rows[r][field] || cols[c][field] || squares[square][field] {
				return false
			}
			rows[r][field] = true
			cols[c][field] = true
			squares[square][field] = true
		}
	}

	return true
}
