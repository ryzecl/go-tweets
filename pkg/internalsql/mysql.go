package internalsql

import (
	"database/sql"
	"fmt"
	"go-tweets/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, "Asia%2FJakarta")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	log.Println("Database connected successfully")
	return db, nil
}
