package dbcon

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

func GetPsqlConnection(config DBConfig) error {
	once.Do(func() {
		domainDataSource := fmt.Sprintf(
			"%s://%s:%s@%s:%d/%s?sslmode=%s",
			config.Driver,
			config.User,
			config.Password,
			config.Server,
			config.Port,
			config.DBName,
			config.SSLMode,
		)

		db, err = sql.Open("postgres", domainDataSource)
	})
	if err != nil {
		return err
	}

	return nil
}

// GetConnection return an unique instance of sql.DB
func GetConnection() *sql.DB {
	return db
}

// Close pool connection
func Close() {
	db.Close()
}
