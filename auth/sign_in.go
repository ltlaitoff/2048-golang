package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/ltlaitoff/2048/db"
	"github.com/ltlaitoff/2048/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SignInUserBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func signInUser(user SignInUserBody) (*bson.ObjectID, error) {
	collection := db.Database.Database("2048").Collection("users")

	log.Println("Sign in user " + user.Email)

	var result entities.User

	// TODO: Add password hashing

	filter := bson.M{"email": user.Email, "password_hash": user.Password}
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("Invalid credentials")
		}

		return nil, err
	}

	return createNewUserSession(result.ID)
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read body
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user SignInUserBody
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	sessionId, err := signInUser(user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "session_id", Value: sessionId.Hex(), Expires: expiration}
	http.SetCookie(w, &cookie)
}
