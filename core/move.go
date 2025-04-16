package core

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/ltlaitoff/2048/pkg/ternary"
)

// TODO: Rename
type MoveData struct {
	XStart, XStep, XEnd int
	YStart, YStep, YEnd int
	NextX, NextY        int
	reversedX           bool
	reversedY           bool
}

func moveWithoutMerge(board *Board, data MoveData) {
	for k := 0; k <= SIZE-1; k++ {
		for y := data.YStart; ternary.Ternary(data.reversedY, y > data.YEnd, y < data.YEnd); y += data.YStep {
			for x := data.XStart; ternary.Ternary(data.reversedX, x > data.XEnd, x < data.XEnd); x += data.XStep {
				if board[y][x] == 0 {
					continue
				}

				newI := y + data.NextY
				newJ := x + data.NextX

				if board[newI][newJ] == 0 {
					board[newI][newJ] = board[y][x]
					board[y][x] = 0
				}
			}
		}
	}
}

func moveWithMerge(board *Board, data MoveData) {
	for y := data.YStart; ternary.Ternary(data.reversedY, y > data.YEnd, y < data.YEnd); y += data.YStep {
		for x := data.XStart; ternary.Ternary(data.reversedX, x > data.XEnd, x < data.XEnd); x += data.XStep {
			if board[y][x] == 0 {
				continue
			}

			newI := y + data.NextY
			newJ := x + data.NextX

			if board[newI][newJ] == board[y][x] {
				newValue := board[y][x] + board[newI][newJ]

				board[newI][newJ] = newValue
				score += Score(newValue)

				board[y][x] = 0
			}
		}
	}
}

// TODO: Refactor
func dataForMove(action string) MoveData {
	switch action {
	case "RIGHT":
		return MoveData{XStart: SIZE - 2, XStep: -1, XEnd: -1, YStart: 0, YStep: 1, YEnd: SIZE, NextX: 1, NextY: 0, reversedX: true}
	case "BOTTOM":
		return MoveData{XStart: 0, XStep: 1, XEnd: SIZE, YStart: SIZE - 2, YStep: -1, YEnd: -1, NextX: 0, NextY: 1, reversedY: true}
	case "LEFT":
		return MoveData{XStart: 1, XStep: 1, XEnd: SIZE, YStart: 0, YStep: 1, YEnd: SIZE, NextX: -1, NextY: 0}
	case "TOP":
		return MoveData{XStart: 0, XStep: 1, XEnd: SIZE, YStart: 1, YStep: 1, YEnd: SIZE, NextX: 0, NextY: -1}
	default:
		log.Fatal("Not valid action")
	}

	log.Fatal("Not valid action")
	return MoveData{}
}

func MoveCells(board *Board, action string) {
	data := dataForMove(action)

	slog.Debug(fmt.Sprintf("Move %s", action))

	moveWithoutMerge(board, data)
	moveWithMerge(board, data)
	moveWithoutMerge(board, data)
}
