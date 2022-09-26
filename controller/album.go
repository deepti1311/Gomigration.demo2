package controller

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/sirupsen/logrus"
	"time"
)

type album struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Artist    string     `json:"artist"`
	Price     string     `json:"price"`
	Completed string     `json:"completed"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func CreatedAlbumTAble(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := db.CreateTable(&album{}, opts)

	if createError != nil {
		logrus.Printf("Error while creating album table , Reason: %v\n", createError)
		return createError
	}

	logrus.Printf("album table created")
	return nil
}

var dbConnection *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnection = db
}
