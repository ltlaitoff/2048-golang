package core

func Move(action string) bool {
	if GameEnd(&board) {
		return true
	}

	MoveCells(&board, action)
	RandomCell(&board)

	return false
}

func Reset() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			board[i][j] = 0
		}
	}

	RandomCell(&board)
}

func Map(callback func(value int64)) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			callback(board[i][j])
		}
	}
}

func State() (Board, Score) {
	return board, score

}

func Init() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			board[i][j] = 0
		}
	}

	// board[0][3] = 4
	// board[0][1] = 2
	// board[0][2] = 2
	// board[0][3] = 0
}
