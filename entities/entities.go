package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `bson:"name"`
	Email        string        `bson:"email"`
	PasswordHash string        `bson:"password_hash"`
	CreatedAt    time.Time     `bson:"created_at"`
}

type Session struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	UserID      bson.ObjectID `bson:"user_id"`
	ActiveRunId bson.ObjectID `bson:"run_id,omitempty"`
	CreatedAt   time.Time     `bson:"created_at"`
	ExpiredAt   time.Time     `bson:"expired_at"`
}

type Run struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	UserID     bson.ObjectID `bson:"user_id"`
	Score      int           `bson:"score"`
	Board      [4][4]int64   `bson:"board"`
	UserAgent  string        `bson:"user_agent"`
	CreatedAt  time.Time     `bson:"created_at"`
	IsFinished bool          `bson:"is_finished,omitempty"`
	FinishedAt *time.Time    `bson:"finished_at,omitempty"`
	Duration   int64         `bson:"duration_ms,omitempty"`
}

type RunHistoryRecord struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	RunID     bson.ObjectID `bson:"run_id"`
	Movement  MovementType  `bson:"movement"`
	Duration  int64         `bson:"movement_duration_ms"`
	CreatedAt time.Time     `bson:"created_at"`
}

type MovementType string

const (
	MovementUp    MovementType = "UP"
	MovementDown  MovementType = "DOWN"
	MovementLeft  MovementType = "LEFT"
	MovementRight MovementType = "RIGHT"
)
