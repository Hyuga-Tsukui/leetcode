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

			IsCurRowNumBlank := board[i][j] != '.'
			IsCurColNumBlank := board[j][i] != '.'
			boxIdx := byte((i/3)*3 + j/3)
			if (IsCurRowNumBlank && row[i][board[i][j]] > 0) || (IsCurColNumBlank && col[i][board[j][i]] > 0) || (IsCurRowNumBlank && box[boxIdx][board[i][j]] > 0) {
				return false
			}

			// 行の数字の出現Mapを作成
			row[i][board[i][j]]++

			// 列の数字の出現Mapを作成
			col[i][board[j][i]]++
			box[boxIdx][board[i][j]]++
		}
	}
	return true
}
