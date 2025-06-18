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

type SignUpUserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func signUpUser(user SignUpUserBody) (*bson.ObjectID, error) {
	collection := db.Database.Database("2048").Collection("users")

	log.Println("Sign up user with name " + user.Name)

	var result entities.User

	// TODO: Add password hashing

	filter := bson.M{"email": user.Email}
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err == nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	if errors.Is(err, mongo.ErrNoDocuments) == false {
		log.Panic(err)
		return nil, fmt.Errorf("Something went wrong!")
	}

	userData := entities.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password,
		CreatedAt:    time.Now(),
	}

	insertionResult, err := collection.InsertOne(context.Background(), userData)

	if err != nil {
		return nil, err
	}

	userObjectId, ok := insertionResult.InsertedID.(bson.ObjectID)

	if !ok {
		return nil, fmt.Errorf("InsertedID is not an ObjectID")
	}

	return createNewUserSession(userObjectId)
}

func AuthSignUp(w http.ResponseWriter, r *http.Request) {
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

	var user SignUpUserBody
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	sessionId, err := signUpUser(user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "session_id", Value: sessionId.Hex(), Expires: expiration, Path: "/"}
	http.SetCookie(w, &cookie)
}
