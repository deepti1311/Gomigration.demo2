package config

import (
	"Gomigration.demo2/model"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
	"os"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "pass",
		Addr:     "localhost:5432",
		Database: "User_records",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		logrus.Printf("Failed to connect")
		os.Exit(100)

	}
	logrus.Printf("Connected to db")
	model.CreatedUserTAble(db)
	model.InitiateDB(db)
	return db
}
