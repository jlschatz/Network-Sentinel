package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"sentinel/models"
)

func (s store) Find(mac string) (models.Entity, error) {
	c := s.db.Collection(os.Getenv("COLLECTION"))

	filter := bson.D{{"mac", mac}}

	var result models.Entity

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return result, err
	}

	fmt.Printf("Found a single document: %v\n", result)

	return result, nil
}

func (s store) Insert(ip, mac string) error {
	s.NewDatabaseConnection()
	c := s.db.Collection(os.Getenv("COLLECTION"))

	insertResult, err := c.InsertOne(context.TODO(), models.Entity{IP: ip, MAC: mac})
	if err != nil {
		log.Println(err)
		return s.CloseDatabaseConnection()
	}

	fmt.Printf("Inserted a single document: %v\n", insertResult.InsertedID)
	return s.CloseDatabaseConnection()

}

func (s store) Update(m models.Entity) error {
	c := s.db.Collection(os.Getenv("COLLECTION"))

	filter := bson.D{{"mac", m.MAC}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "ip", Value: m.IP},
	}}}

	entity := &models.Entity{}

	return c.FindOneAndUpdate(context.TODO(), filter, update).Decode(entity)
}
