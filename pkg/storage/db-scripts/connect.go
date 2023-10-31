package scripts

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var CONN *pgx.Conn

const connectMsg string = "---------------------------------------------------------------------------------------------\nConnected to DB\n---------------------------------------------------------------------------------------------"

func ConnectDB() *pgx.Conn {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("SQLURI")

	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		log.Println(err)
		return nil
	}
	CONN = conn

	fmt.Println(connectMsg)
	return conn
}
