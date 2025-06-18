package auth

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ltlaitoff/2048/db"
	"github.com/ltlaitoff/2048/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SignInUserBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignInUser(user SignInUserBody) error {
	collection := db.Database.Database("2048").Collection("users")

	log.Println("Sign in user " + user.Email)

	var result entities.User

	// TODO: Add password hashing

	filter := bson.M{"email": user.Email, "password_hash": user.Password}
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("Invalid credentials")
		}

		log.Panic(err)
	}

	// TODO: Create session and add this to cookies

	return nil
}
