package storage

import (
	"auth/pkg/storage/postgres"
)

func ConnectDB() {
	postgres.Postgres()
}
