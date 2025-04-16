package core

import ()

func GameEnd(board *Board) bool {
	for y := 0; y < SIZE; y++ {
		for x := 0; x < SIZE; x++ {
			if board[x][y] == 0 {
				return false
			}

			// Previous row
			if y != 0 {
				if board[y][x] == board[y-1][x] {
					return false
				}
			}

			// Current row
			if x != 0 && board[y][x] == board[y][x-1] {
				return false
			}

			if x != SIZE-1 && board[y][x] == board[y][x+1] {
				return false
			}

			// Next row
			if y != SIZE-1 {
				if board[y][x] == board[y+1][x] {
					return false
				}
			}
		}
	}

	return true
}
