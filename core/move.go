package core

import (
	"fmt"
	"log"
	"log/slog"
)

const (
	LeftAction   = "left"
	RightAction  = "right"
	TopAction    = "top"
	BottomAction = "bottom"
)

// TODO: Rename
type MoveData struct {
	XStart, XStep, XEnd int
	YStart, YStep, YEnd int
	NextX, NextY        int
}

func moveWithoutMerge(data MoveData) {
	// DEV: Think :)  
	for k := 0; k <= SIZE - 1; k++ {
		for i := data.YStart; i < data.YEnd; i += data.YStep {
			for j := data.XStart; j < data.XEnd; j += data.XStep {
				if board[i][j] == 0 {
					continue
				}

				newI := i + data.NextY
				newJ := j + data.NextX

				slog.Debug(fmt.Sprintf("Remove zeros i = %d, j = %d, newI = %d, newJ = %d, current = %d, new = %d", i, j, newI, newJ, board[i][j], board[newI][newJ]))

				if board[newI][newJ] == 0 {
					board[newI][newJ] = board[i][j]
					board[i][j] = 0
				}
			}
		}
	}
}

func merge(data MoveData) {
	for i := data.YStart; i < data.YEnd; i += data.YStep {
		for j := data.XStart; j < data.XEnd; j += data.XStep {
			if board[i][j] == 0 {
				continue
			}

			newI := i + data.NextY
			newJ := j + data.NextX

			slog.Debug(fmt.Sprintf("Merge i = %d, j = %d, newI = %d, newJ = %d, current = %d, new = %d", i, j, newI, newJ, board[i][j], board[newI][newJ]))

			if board[newI][newJ] == board[i][j] {
				board[newI][newJ] = board[i][j] + board[newI][newJ]
				board[i][j] = 0
			}
		}
	}
}

// DEV: Refactor
func moveData(action string) MoveData {
	switch action {
	case "RIGHT":
		return MoveData{XStart: 0, XStep: 1, XEnd: SIZE - 1, YStart: 0, YStep: 1, YEnd: SIZE, NextX: 1, NextY: 0}
	case "LEFT":
		return MoveData{XStart: 1, XStep: 1, XEnd: SIZE, YStart: 0, YStep: 1, YEnd: SIZE, NextX: -1, NextY: 0}
	case "DOWN":
		return MoveData{XStart: 0, XStep: 1, XEnd: SIZE, YStart: 0, YStep: 1, YEnd: SIZE - 1, NextX: 0, NextY: 1}
	case "UP":
		return MoveData{XStart: 0, XStep: 1, XEnd: SIZE, YStart: 1, YStep: 1, YEnd: SIZE, NextX: 0, NextY: -1}
	default:
		log.Fatal("Not valid action")
	}

	log.Fatal("Not valid action")
	return MoveData{}
}

func Move(action string) {
	data := moveData(action)

	slog.Debug(fmt.Sprintf("Move %s", action))

	moveWithoutMerge(data)
	merge(data)
	moveWithoutMerge(data)
}

