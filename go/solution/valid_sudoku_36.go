package solution

func isValidSudoku(board [][]byte) bool {

	row := make([]map[byte]int, 9)
	col := make([]map[byte]int, 9)
	box := make([]map[byte]int, 9)

	for i := range 9 {
		row[i] = make(map[byte]int)
		col[i] = make(map[byte]int)
		box[i] = make(map[byte]int)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			v := board[i][j]
			boxIdx := byte((i/3)*3 + j/3)
			if (row[i][v] > 0) || col[j][v] > 0 || box[boxIdx][v] > 0 {
				return false
			}

			row[i][v]++
			col[j][v]++
			box[boxIdx][v]++
		}
	}
	return true
}
