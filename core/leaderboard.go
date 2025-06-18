package core

import (
	"context"
	"log"
	"time"

	"github.com/ltlaitoff/2048/db"
	"github.com/ltlaitoff/2048/entities"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type LeaderboardRow struct {
	Name  string
	Score int64
	Date  time.Time
}

// GetLeaderboard повертає топ-10 останніх Run за CreatedAt (останній зверху)
func GetLeaderboard() ([]LeaderboardRow, error) {
	collection := db.Database.Database("2048").Collection("runs")
	filter := bson.M{"is_finished": true}
	opts := options.Find().SetSort(bson.D{{"score", -1}, {"created_at", -1}}).SetLimit(10)
	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		log.Println("Leaderboard: find error:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var runs []entities.Run
	if err := cursor.All(context.Background(), &runs); err != nil {
		log.Println("Leaderboard: cursor.All error:", err)
		return nil, err
	}
	log.Printf("Leaderboard: found %d runs\n", len(runs))
	if len(runs) == 0 {
		log.Println("Leaderboard: no finished runs found")
	}

	rows := make([]LeaderboardRow, 0, len(runs))
	for _, run := range runs {
		user, _ := getUserByID(run.UserID)
		rows = append(rows, LeaderboardRow{
			Name:  user.Name,
			Score: run.Score,
			Date:  run.CreatedAt,
		})
	}
	return rows, nil
}

func getUserByID(id bson.ObjectID) (*entities.User, error) {
	collection := db.Database.Database("2048").Collection("users")
	var user entities.User
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return &entities.User{Name: "Unknown"}, err
	}
	return &user, nil
}
