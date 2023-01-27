package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"sentinel/models"
)

type Store interface {
	NewDatabaseConnection() error
	CloseDatabaseConnection() error
	Insert(ip, mac string) error
	Find(mac string) (models.Entity, error)
	Update(m models.Entity) error
}

type store struct {
	db *mongo.Database
}

func NewStore() Store {
	return &store{}
}

func (s *store) NewDatabaseConnection() error {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO")))
	if err != nil {
		return err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}

	s.db = client.Database(os.Getenv("MONGO_DATABASE"))

	return err
}

func (s *store) CloseDatabaseConnection() error {
	return s.db.Client().Disconnect(context.TODO())
}
