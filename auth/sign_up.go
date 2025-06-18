package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
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

func SignUpUser(user SignUpUserBody) error {
	collection := db.Database.Database("2048").Collection("users")

	log.Println("Sign up user with name " + user.Name)

	var result entities.User

	// TODO: Add password hashing

	filter := bson.M{"email": user.Email}
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err == nil {
		return fmt.Errorf("Invalid credentials")
	}

	if errors.Is(err, mongo.ErrNoDocuments) == false {
		log.Panic(err)
		return fmt.Errorf("Something went wrong!")
	}

	newUser := entities.User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password,
		CreatedAt:    time.Now(),
	}

	_, err = collection.InsertOne(context.Background(), newUser)

	if err != nil {
		return err
	}

	// TODO: Create session and add this to cookies

	return nil
}
