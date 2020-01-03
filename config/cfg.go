package cfg

import (
	"github.com/jinzhu/gorm"

	// driver postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB struct type
type DB struct {
	*gorm.DB
}

// Connect database setup configuration
func Connect(connectionDb string) (*DB, error) {
	db, err := gorm.Open("postgres", connectionDb)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
