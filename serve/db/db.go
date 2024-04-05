package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Connection interface {
	Close()
	DB() *gorm.DB
	RegisterModels(...interface{})
	Migrate()
}

type conn struct {
	db *gorm.DB
}

var modelsRegistered []interface{}

func NewConnection() Connection {
	log.Println("creating connection")
	var c conn
	db, err := gorm.Open(mysql.Open(getDNS()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	c.db = db
	return &c
}

func (c *conn) RegisterModels(models ...interface{}) {
	log.Printf("Registering %d model", len(models))
	// print log name for each model
	for _, model := range models {
		log.Printf("Registering model %T", model)
	}
	modelsRegistered = append(modelsRegistered, models...)
}

func (c *conn) Migrate() {
	log.Println("Migrating models")
	err := c.db.AutoMigrate(modelsRegistered...)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *conn) Close() {
	log.Println("Closing connection")
	db, err := c.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *conn) DB() *gorm.DB {
	return c.db
}

func getDNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQLUSER"),
		os.Getenv("MYSQLPASSWORD"),
		os.Getenv("MYSQLHOST"),
		os.Getenv("MYSQLPORT"),
		os.Getenv("MYSQLDATABASE"),
	)
}
