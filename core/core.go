package core

import (
	"github.com/ltlaitoff/2048/entities"
)

func Move(action string, session *entities.Session, userAgent string) bool {
	var run *entities.Run
	var err error

	if session.ActiveRunId.IsZero() {
		runId, err := CreateNewRun(session.UserID, Board{}, userAgent)
		if err != nil {
			return false
		}

		session.ActiveRunId = *runId
		run, _ = FindRunByID(runId.Hex())

		_ = UpdateSessionActiveRunId(session.ID, *runId)
	} else {
		run, err = FindRunByID(session.ActiveRunId.Hex())
		if err != nil {
			return false
		}
	}

	if run.IsFinished {
		return false
	}

	board := (*Board)(&run.Board)

	if GameEnd(board) {
		run.IsFinished = true
		_ = UpdateRun(run.ID, run.Board, run.Score, true)
		return true
	}

	MoveCells(board, action)
	RandomCell(board)

	_ = UpdateRun(run.ID, run.Board, run.Score, false)

	return false
}

func Reset() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			board[i][j] = 0
		}
	}

	score = 0

	RandomCell(&board)
}

func Map(callback func(value int64)) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			callback(board[i][j])
		}
	}
}

func State(session *entities.Session) (Board, Score, bool) {
	if session != nil && !session.ActiveRunId.IsZero() {
		run, err := FindRunByID(session.ActiveRunId.Hex())

		if err == nil {
			b := (*Board)(&run.Board)

			return *b, Score(run.Score), run.IsFinished
		}
	}

	var empty Board

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			empty[i][j] = 0
		}
	}

	return empty, 0, false
}

func Init() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			board[i][j] = 0
		}
	}

	score = 0

	// board[0][3] = 4
	// board[0][1] = 2
	// board[0][2] = 2
	// board[0][3] = 0
}
