package model

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/sirupsen/logrus"
	"time"
)

type User struct {
	ID        string     `json:"id"`
	FirstName string     `json:"title"`
	LastName  string     `json:"artist"`
	UserName  string     `json:"price"`
	Email     string     `json:"email"`
	Completed string     `json:"completed"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func CreatedUserTAble(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := db.CreateTable(&User{}, opts)

	if createError != nil {
		logrus.Printf("Error while creating table , Reason: %v\n", createError)
		return createError
	}

	logrus.Printf("table created")
	return nil
}

var dbConnection *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnection = db
}
