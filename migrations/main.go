package main

import (
	"fmt"
	"os"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("fail to load .env file")
	}
}

func runMigrations(conn *pg.DB) error {
	// migrate database before starting services
	migrations.SetTableName("migrations")

	// init migrations table
	_, _, err := migrations.Run(conn, "init")
	if err != nil {
		return err
	}

	// up migrations
	old, new, err := migrations.Run(conn, "up")
	if err != nil {
		return err
	}

	fmt.Printf("\ndb version: %d -> %d\n", old, new)

	return nil
}

func main() {
	conn := pg.Connect(&pg.Options{
		Addr:     "db:"+os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
	})
	defer conn.Close()

	// migrate database before starting app
	err := runMigrations(conn)
	if err != nil {
		panic(err.Error())
	}
}