package solution

func isValidSudoku(board [][]byte) bool {

	row := make([]map[byte]int, 9)
	col := make([]map[byte]int, 9)
	box := make([]map[byte]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// 行の数字の出現Mapを作成
			if row[i] == nil {
				row[i] = make(map[byte]int)
			}
			row[i][board[i][j]]++

			// 列の数字の出現Mapを作成
			if col[i] == nil {
				col[i] = make(map[byte]int)
			}
			col[i][board[j][i]]++

			if i <= 2 && j <= 2 {
				if box[0] == nil {
					box[0] = make(map[byte]int)
				}
				box[0][board[i][j]]++
			}

			if i <= 2 && (3 <= j && j <= 5) {
				if box[1] == nil {
					box[1] = make(map[byte]int)
				}
				box[1][board[i][j]]++
			}

			if i <= 2 && (6 <= j && j <= 8) {
				if box[2] == nil {
					box[2] = make(map[byte]int)
				}
				box[2][board[i][j]]++
			}

			if (3 <= i && i <= 5) && j <= 2 {
				if box[3] == nil {
					box[3] = make(map[byte]int)
				}
				box[3][board[i][j]]++
			}
			if (3 <= i && i <= 5) && (3 <= j && j <= 5) {
				if box[4] == nil {
					box[4] = make(map[byte]int)
				}
				box[4][board[i][j]]++
			}
			if (3 <= i && i <= 5) && (6 <= j && j <= 8) {
				if box[5] == nil {
					box[5] = make(map[byte]int)
				}
				box[5][board[i][j]]++
			}
			if (6 <= i && i <= 8) && (j <= 2) {
				if box[6] == nil {
					box[6] = make(map[byte]int)
				}
				box[6][board[i][j]]++
			}
			if (6 <= i && i <= 8) && (3 <= j && j <= 5) {
				if box[7] == nil {
					box[7] = make(map[byte]int)
				}
				box[7][board[i][j]]++
			}
			if (6 <= i && i <= 8) && (6 <= j && j <= 8) {
				if box[8] == nil {
					box[8] = make(map[byte]int)
				}
				box[8][board[i][j]]++
			}

		}
	}

	for _, v := range []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'} {
		for i := range 9 {
			if v, ok := row[i][v]; ok {
				if v > 1 {
					return false
				}
			}
			if v, ok := box[i][v]; ok {
				if v > 1 {
					return false
				}
			}
			if v, ok := col[i][v]; ok {
				if v > 1 {
					return false
				}
			}
		}
	}
	return true
}
