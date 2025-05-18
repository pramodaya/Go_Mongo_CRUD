package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go_mongo_crud/traveluser"
)

var collection *mongo.Collection

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use MONGO_URI from env var or fallback to default
	mongoURI := "mongodb://mongodb:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("MongoDB connect error: %v", err)
	}

	collection = client.Database("travelapp").Collection("travelusers")
	log.Println("Connected to MongoDB")

	http.HandleFunc("/travelusers", travelUsersHandler)
	http.HandleFunc("/travelusers/", travelUserHandler)

	log.Println("API running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func travelUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var users []traveluser.TravelUser
		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			http.Error(w, "Error retrieving users", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {
			var u traveluser.TravelUser
			cursor.Decode(&u)
			users = append(users, u)
		}
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var u traveluser.TravelUser
		json.NewDecoder(r.Body).Decode(&u)
		u.ID = uuid.New().String()

		_, err := collection.InsertOne(context.TODO(), u)
		if err != nil {
			http.Error(w, "Insert failed", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(u)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func travelUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/travelusers/"):]

	switch r.Method {
	case http.MethodGet:
		var u traveluser.TravelUser
		err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&u)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(u)

	case http.MethodPut:
		var u traveluser.TravelUser
		json.NewDecoder(r.Body).Decode(&u)
		_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": u})
		if err != nil {
			http.Error(w, "Update failed", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(u)

	case http.MethodDelete:
		_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
		if err != nil {
			http.Error(w, "Delete failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
