package storage

import "authTest/pkg/storage/postgres"

func ConnectDB() {
	postgres.Postgres()
}
