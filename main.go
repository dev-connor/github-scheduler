package main

import (
	"joosum-scheduler/app/notification"
	"joosum-scheduler/pkg/config"
	"joosum-scheduler/pkg/db"
	"log"
	"os"
)

func main() {
	config.EnvConfig()
	db.StartMongoDB()

	var notificationType string
	if len(os.Args) > 1 {
		notificationType = os.Args[1]
	}

	if notificationType == "unread" {
		notification.SendUnreadLink()
	} else if notificationType == "unclassified" {
		notification.SendUnclassifiedLink()
	} else {
		log.Fatal("invalid notificationType")
	}

	db.CloseMongoDB()
}
