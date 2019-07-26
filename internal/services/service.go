package services

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/disturb16/graphql_golang/settings"
)

// Service struct definition containing all services functionalties
type Service struct {
	db *sql.DB
}

// New initialize service and database connections
func New(dbDriver string, config *settings.Config) (*Service, error) {

	s := &Service{}

	// setup database connection
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name)

	db, err := sql.Open(dbDriver, connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	s.db = db

	return s, nil
}
