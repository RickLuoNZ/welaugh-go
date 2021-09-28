package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Client *gorm.DB
}

func New(connStr string) (*DB, error) {
	db, err := get(connStr)

	if err != nil {
		return nil, err
	}
	
	return &DB{
		Client: db,
	}, nil
}

func get(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(200)
	db.LogMode(true)

	return db, nil
}

func (d *DB) Close() error {
	return d.Client.Close()
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../realworld_test.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../realworld_test.db"); err != nil {
		return err
	}
	return nil
}
