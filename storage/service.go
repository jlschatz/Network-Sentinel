package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"sentinel/models"
)

type Store interface {
	NewDatabaseConnection() error
	CloseDatabaseConnection() error
	Upsert(ip, mac string) error
	Find(mac string) (models.Entity, error)
	FindAll() ([]models.Entity, error)
	ValidateEnvironmentVariables() error
	SetTestEnvironmentVariables() error
}

type store struct {
	db       *mongo.Database
	entities []models.Entity
}

func NewStore() Store {
	s := &store{}
	if err := s.SetTestEnvironmentVariables(); err != nil {
		log.Fatal(err)
	}
	if err := s.ValidateEnvironmentVariables(); err != nil {
		log.Fatal(err)
	}
	return s
}

func (s *store) NewDatabaseConnection() error {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO")))
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
		return err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println(err.Error())
		return err
	}

	s.db = client.Database(os.Getenv("MONGO_DATABASE"))

	return err
}

func (s *store) CloseDatabaseConnection() error {
	return s.db.Client().Disconnect(context.TODO())
}

func (s store) SetTestEnvironmentVariables() error {
	if err := os.Setenv("MONGO", "mongodb://localhost:27017"); err != nil {
		return err
	}
	if err := os.Setenv("MONGO_DATABASE", "test"); err != nil {
		return err
	}
	if err := os.Setenv("COLLECTION", "Network_Sentinel"); err != nil {
		return err
	}
	return nil
}

func (s store) ValidateEnvironmentVariables() error {
	if os.Getenv("MONGO") == "" {
		return errors.New("MONGO env variable not set")
	}
	if os.Getenv("MONGO_DATABASE") == "" {
		return errors.New("MONGO_DATABASE env variable not set")
	}
	if os.Getenv("COLLECTION") == "" {
		return errors.New("COLLECTION env variable not set")
	}

	return nil
}
