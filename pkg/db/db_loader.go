package db

import (
	"context"
	"joosum-scheduler/pkg/config"
	"log"
	"time"
)

func StartMongoDB() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := GetMongoClient(ctx, config.GetEnvConfig("mongoDB"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	dbName := config.GetEnvConfig("dbName")

	// Collection load
	InitUserCollection(client, dbName)
	InitLinkCollection(client, dbName)
	InitLinkBookCollection(client, dbName)
	InitInactiveUserCollection(client, dbName)
	InitTagCollection(client, dbName)
	InitNotificationCollection(client, dbName)
}

func CloseMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := DisconnectMongoClient(ctx)
	if err != nil {
		log.Fatalf("Failed to disconnect to MongoDB: %v", err)
	}
}
