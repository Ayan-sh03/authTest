package storage

import "yuj/pkg/storage/postgres"

func ConnectDB() {
	postgres.Postgres()
}
