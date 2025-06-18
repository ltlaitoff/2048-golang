package core

import (
	"time"

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

	score := run.Score

	// duration: різниця між CreatedAt поточного run і CreatedAt останнього запису runs_history
	duration := int64(0)
	history, err := GetLastRunHistoryRecord(run.ID)
	if err == nil && history != nil {
		duration = time.Now().Sub(history.CreatedAt).Milliseconds()
	}

	MoveCells(board, &score, action)
	_ = AddRunHistoryRecord(run.ID, action, duration)
	RandomCell(board)

	_ = UpdateRun(run.ID, run.Board, score, false)

	return false
}

func Reset(session *entities.Session) {
	if session != nil {
		runId, err := CreateNewRun(session.UserID, Board{}, "Reset agent")
		if err == nil {
			session.ActiveRunId = *runId
			_ = UpdateSessionActiveRunId(session.ID, *runId)
		}
	}
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
