package core

import (
	"context"
	"fmt"
	"time"

	"github.com/ltlaitoff/2048/db"
	"github.com/ltlaitoff/2048/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// FindRunByID знаходить Run за його ID
func FindRunByID(runID string) (*entities.Run, error) {
	collection := db.Database.Database("2048").Collection("runs")

	objectID, err := bson.ObjectIDFromHex(runID)
	if err != nil {
		return nil, fmt.Errorf("invalid run_id: %w", err)
	}

	var run entities.Run
	filter := bson.M{"_id": objectID}
	err = collection.FindOne(context.Background(), filter).Decode(&run)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("run not found")
		}
		return nil, err
	}
	return &run, nil
}

// CreateNewRun створює новий Run для користувача
func CreateNewRun(userID bson.ObjectID, board Board, userAgent string) (*bson.ObjectID, error) {
	collection := db.Database.Database("2048").Collection("runs")

	runData := entities.Run{
		UserID:     userID,
		Score:      0,
		Board:      board,
		UserAgent:  userAgent,
		CreatedAt:  time.Now(),
		IsFinished: false,
	}

	ins, err := collection.InsertOne(context.Background(), runData)
	if err != nil {
		return nil, err
	}
	resultObjectId, ok := ins.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, fmt.Errorf("InsertedID is not an ObjectID")
	}
	return &resultObjectId, nil
}

// FindActiveRunForUser знаходить активний (не завершений) Run для користувача
func FindActiveRunForUser(userID bson.ObjectID) (*entities.Run, error) {
	collection := db.Database.Database("2048").Collection("runs")
	filter := bson.M{"user_id": userID, "is_finished": false}
	var run entities.Run
	err := collection.FindOne(context.Background(), filter).Decode(&run)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Немає активного Run
		}
		return nil, err
	}
	return &run, nil
}

// UpdateRun оновлює поле Board, Score, IsFinished для Run за ID
func UpdateRun(runID bson.ObjectID, board [4][4]int64, score int, isFinished bool) error {
	collection := db.Database.Database("2048").Collection("runs")
	filter := bson.M{"_id": runID}
	update := bson.M{"$set": bson.M{
		"board":       board,
		"score":       score,
		"is_finished": isFinished,
	}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

// UpdateSessionActiveRunId оновлює поле run_id у сесії за її ID
func UpdateSessionActiveRunId(sessionID, runID bson.ObjectID) error {
	collection := db.Database.Database("2048").Collection("sessions")
	filter := bson.M{"_id": sessionID}
	update := bson.M{"$set": bson.M{"run_id": runID}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}
