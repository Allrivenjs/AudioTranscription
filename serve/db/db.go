package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strconv"
	"time"
)

type Connection interface {
	Close()
	DB() *mongo.Database
}

type conn struct {
	client *mongo.Client
}

func NewConnection() Connection {
	var c conn
	var err error
	url := getURL()
	clientOptions := options.Client().ApplyURI(url)
	c.client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Panicln(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = c.client.Connect(ctx)
	if err != nil {
		log.Panicln(err.Error())
	}
	return &c
}

func (c *conn) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := c.client.Disconnect(ctx); err != nil {
		log.Println("Error al desconectar:", err)
	}
}

func (c *conn) DB() *mongo.Database {
	return c.client.Database(os.Getenv("DATABASE_NAME"))
}

func getURL() string {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Println("Error al cargar el puerto de la base de datos desde la variable de entorno:", err)
		port = 27017
	}
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		port,
		os.Getenv("DATABASE_NAME"))
}
