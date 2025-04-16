package core

import (
	"log/slog"
	"math/rand"

	"github.com/ltlaitoff/2048/pkg/assert"
)

func getEmptyIndexes(board *Board) [][2]int {
	res := make([][2]int, 0, 16)

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board[i][j] != 0 {
				continue
			}

			res = append(res, [2]int{i, j})
		}
	}

	return res
}

func RandomCell(board *Board) {
	emptyIndexes := getEmptyIndexes(board)

	if len(emptyIndexes) == 0 {
		slog.Debug("Not add random cell because all cells is not empty")
		return
	}

	index := rand.Intn(len(emptyIndexes))
	i := emptyIndexes[index][0]
	j := emptyIndexes[index][1]

	assert.Assert(board[i][j] == 0, "Board cell on random add not equal to 0")

	isFour := rand.Intn(10)

	if isFour == 9 {
		board[i][j] = 4
	} else {
		board[i][j] = 2
	}
}
