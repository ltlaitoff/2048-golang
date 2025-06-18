package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ltlaitoff/2048/db"
	"github.com/ltlaitoff/2048/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func IsAuthorized(r *http.Request) (*bool, error) {
	cookie, err := r.Cookie("session_id")
	notAuthorized := false

	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return &notAuthorized, nil
		}

		return nil, err
	}

	cookieValue := cookie.Value

	session, err := FindSessionByID(cookieValue)

	if err != nil {
		return nil, err
	}

	if session.ExpiredAt.Before(time.Now()) {
		return &notAuthorized, nil
	}

	isAuthorized := true
	return &isAuthorized, nil
}

func FindSessionByID(sessionID string) (*entities.Session, error) {
	collection := db.Database.Database("2048").Collection("sessions")

	objectID, err := bson.ObjectIDFromHex(sessionID)
	if err != nil {
		return nil, fmt.Errorf("invalid session_id: %w", err)
	}

	var session entities.Session

	filter := bson.M{"_id": objectID}
	err = collection.FindOne(context.Background(), filter).Decode(&session)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("session not found")
		}
		return nil, err
	}

	return &session, nil
}
