package entities

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type User struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	CreatedAt time.Time     `bson:"created_at"`
}

type Session struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	UserID     bson.ObjectID `bson:"user_id"`
	Score      int           `bson:"score"`
	Board      []int         `bson:"board"`
	UserAgent  string        `bson:"user_agent"`
	CreatedAt  time.Time     `bson:"created_at"`
	IsFinished bool          `bson:"is_finished,omitempty"`
	FinishedAt *time.Time    `bson:"finished_at,omitempty"`
}

type SessionHistory struct {
	ID               bson.ObjectID `bson:"_id,omitempty"`
	SessionID        bson.ObjectID `bson:"session_id"`
	Movement         MovementType  `bson:"movement"`
	MovementDuration float64       `bson:"movement_duration_ms"`
	CreatedAt        time.Time     `bson:"created_at"`
}

type Leaderboard struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	UserID     bson.ObjectID `bson:"user_id"`
	Score      int           `bson:"score"`
	DurationMs float64       `bson:"duration_ms"`
	UserAgent  string        `bson:"user_agent,omitempty"`
	CreatedAt  time.Time     `bson:"created_at"`
}

type MovementType string

const (
	MovementUp    MovementType = "UP"
	MovementDown  MovementType = "DOWN"
	MovementLeft  MovementType = "LEFT"
	MovementRight MovementType = "RIGHT"
)
