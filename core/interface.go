package core

func Up() {
	Move("UP")
	addRandomCell()
}

func Left() {
	Move("LEFT")
	addRandomCell()
}

func Right() {
	Move("RIGHT")
	addRandomCell()
}

func Down() {
	Move("DOWN")
	addRandomCell()
}

func Reset() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			board[i][j] = 0
		}
	}
	
	// Init()
	// addRandomCell()
}

func Map(callback func(value int64)) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			callback(board[i][j])
		}
	}
}

func Init() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			board[i][j] = 0
		}
	}

	// board[0][0] = 2
	// board[0][1] = 2
	// board[0][2] = 2
	// board[0][3] = 2
	//
	// board[0][0] = 2
	// board[1][0] = 2
	// board[2][0] = 2
	// board[3][0] = 2
}
