package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Store struct {
	Session           *mongo.Client
	DatabaseName      string
	ConnectionTimeout int
}

func NewStore(connectionString, dbName string, timeout int) (*Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+connectionString))
	if err != nil {
		return nil, err
	}

	ctxPing, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	if err = client.Ping(ctxPing, readpref.Primary()); err != nil {
		return nil, err
	}

	return &Store{
		Session:           client,
		DatabaseName:      dbName,
		ConnectionTimeout: timeout,
	}, nil
}

func (s *Store) GetCollection(collectionName string) *mongo.Collection {
	return s.Session.Database(s.DatabaseName).Collection(collectionName)
}
