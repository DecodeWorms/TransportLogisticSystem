package storage

import (
	"Tsystem/config"
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func NewConnection(ctx context.Context, config config.Config, db *gorm.DB, err error) *Client {
	log.Println("Connecting to DB...")
	uri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s", config.DatabaseHost, config.DatabaseUsername, config.DatabaseName, config.DatabasePort)

	db, err = gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		fmt.Println(err)

	}
	log.Println("Connected to DB...")
	return &Client{
		db: db,
	}
	// client := Client{
	// 	db: db,
	// }

	// return &client
}
