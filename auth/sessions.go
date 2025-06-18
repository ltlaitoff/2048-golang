package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ltlaitoff/2048/db"
	"github.com/ltlaitoff/2048/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func createNewUserSession(userId bson.ObjectID) (*bson.ObjectID, error) {
	collection := db.Database.Database("2048").Collection("sessions")

	log.Println("Create session for " + userId.String())

	sessionData := entities.Session{
		UserID:    userId,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(365 * 24 * time.Hour),
	}

	session, err := collection.InsertOne(context.Background(), sessionData)

	if err != nil {
		return nil, err
	}

	resultObjectId, ok := session.InsertedID.(bson.ObjectID)

	if !ok {
		return nil, fmt.Errorf("InsertedID is not an ObjectID")
	}

	return &resultObjectId, nil
}
