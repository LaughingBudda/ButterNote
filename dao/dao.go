package dao

import (
	"context"
	"fmt"
	"log"
	"github.com/LaughingBudda/ButterNote/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb+srv://butterKing:fingerLickinGood@butternote-2m1oj.mongodb.net/test?retryWrites=true&w=majority"
//const CONNECTIONSTRING = "mongodb://localhost:27017"

// DBNAME Database name
const DBNAME = "butternote"

// COLLNAME Collection name
const COLLNAME = "notes"

var db *mongo.Database

// Connect establish a connection to database
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

// InsertManyValues inserts many items from byte slice
func InsertManyValues(people []models.Note) {
	var ppl []interface{}
	for _, p := range people {
		ppl = append(ppl, p)
	}
	_, err := db.Collection(COLLNAME).InsertMany(context.Background(), ppl)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertOneValue inserts one item from Note model
func InsertOneValue(person models.Note) {
	fmt.Println(person)
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllPeople returns all people from DB
func GetAllPeople() []models.Note {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var elements []models.Note
	var elem models.Note
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}

// DeletePerson deletes an existing person
func DeletePerson(person models.Note) {
	_, err := db.Collection(COLLNAME).DeleteOne(context.Background(), person, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdatePerson updates an existing person
func UpdatePerson(person models.Note, personID string) {
	doc := db.Collection(COLLNAME).FindOneAndUpdate(
		context.Background(),
		bson.D{{"id", personID}},
		bson.D{{
			"$set", 
			bson.D{{"time", person.Timestamp}, 
			{"content", person.Content}, 
			{"contactinfo.id", person.ID}, 
			{"contactinfo.name", person.Name}}}},
	)
	fmt.Println(doc)
}
