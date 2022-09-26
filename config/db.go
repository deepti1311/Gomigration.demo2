package config

import (
	"Gomigration.demo2/controller"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
	"os"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "pass",
		Addr:     "localhost:5432",
		Database: "Album_records",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		logrus.Printf("Failed to connect")
		os.Exit(100)

	}
	logrus.Printf("Connected to db")
	controller.CreatedAlbumTAble(db)
	controller.InitiateDB(db)
	return db
}
