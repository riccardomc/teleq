package database

import "github.com/jinzhu/gorm"

//Database represents a database
type Database struct {
	dialect string
	url     string
	client  *gorm.DB
}

//NewDatabase database
func NewDatabase(dialect, url string) *Database {
	return &Database{dialect, url, nil}
}

//Connect to database
func (db *Database) Connect() error {
	client, err := gorm.Open(db.dialect, db.url)
	if err != nil {
		return err
	}
	db.client = client
	return nil
}
