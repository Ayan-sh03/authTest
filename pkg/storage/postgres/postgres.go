package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var CONN *pgx.Conn

const connectMsg string = "---------------------------------------------------------------------------------------------\nConnected to DB\n---------------------------------------------------------------------------------------------"

func Postgres() *pgx.Conn {
	ctx := context.Background()
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
