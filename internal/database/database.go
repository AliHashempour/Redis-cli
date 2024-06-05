package database

import (
	"errors"
	"sync"
)

type Database struct {
	Name string
	Data map[string]interface{}
	lock sync.RWMutex
}

func NewDatabase(dbname string) *Database {
	return &Database{
		Name: dbname,
		Data: make(map[string]interface{}),
	}
}

func (db *Database) Set(key string, value interface{}) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.Data[key] = value
	return nil
}

func (db *Database) Get(key string) (interface{}, error) {

	value, exists := db.Data[key]
	if !exists {
		return nil, errors.New("key not found")
	}
	return value, nil
}

func (db *Database) Delete(key string) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	if _, exists := db.Data[key]; !exists {
		return errors.New("key not found")
	}
	delete(db.Data, key)
	return nil
}
