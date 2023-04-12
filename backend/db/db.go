package db

import (
	"context"
	"log"

	sqlite "go.elastic.co/apm/module/apmgormv2/v2/driver/sqlite"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID       uint
	FullName string
	Point    int
}

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tennisATP.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("EXIT PROBLEM DB", err)
	}
}

func GetDB(ctx context.Context) *gorm.DB {
	db := DB
	db = db.WithContext(ctx)

	return db
}
