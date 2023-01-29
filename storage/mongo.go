package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sentinel/models"
)

func (s store) FindAll() ([]models.Entity, error) {
	c := s.db.Collection(os.Getenv("COLLECTION"))
	findOptions := options.Find()

	var results []models.Entity

	cur, err := c.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Println(err)
		return results, err
	}

	for cur.Next(context.TODO()) {

		var elem models.Entity
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
			return results, err
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return results, err
	}

	cur.Close(context.TODO())

	return results, nil
}

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

func (s store) Upsert(ip, mac string) error {
	s.NewDatabaseConnection()
	c := s.db.Collection(os.Getenv("COLLECTION"))

	filter := bson.D{{"mac", mac}}
	update := bson.D{{"$set", bson.D{{"ip", ip}, {"mac", mac}}}}
	opts := options.Update().SetUpsert(true)

	result, err := c.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		s.CloseDatabaseConnection()
		return err
	}

	if result.ModifiedCount > 0 {
		fmt.Println(fmt.Printf("Updated entity: %v", mac))
	}
	if result.UpsertedCount > 0 {
		fmt.Println(fmt.Printf("Inserted entity: IP %v MAC %v", ip, mac))
	}
	return s.CloseDatabaseConnection()
}
