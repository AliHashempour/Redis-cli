package storage

import (
	"errors"
	"fmt"
	"go-memory/internal/database"
)

type Storage struct {
	AllDatabases    []*database.Database
	CurrentDatabase *database.Database
}

func NewStorage() *Storage {
	defaultDb := database.NewDatabase("default")
	return &Storage{
		AllDatabases:    []*database.Database{defaultDb},
		CurrentDatabase: defaultDb,
	}
}

func (s *Storage) AddDatabase(name string) (*database.Database, error) {
	if s.CheckDBExists(name) {
		return nil, errors.New(fmt.Sprintf("Database %s is already created", name))
	}

	newDb := database.NewDatabase(name)
	s.AllDatabases = append(s.AllDatabases, newDb)
	return newDb, nil
}

func (s *Storage) ListAllDatabases() []string {
	var names []string
	for _, db := range s.AllDatabases {
		names = append(names, db.Name)
	}
	return names
}

func (s *Storage) GetDb(name string) (*database.Database, bool) {
	for _, db := range s.AllDatabases {
		if db.Name == name {
			return db, true
		}
	}
	return nil, false
}

func (s *Storage) CheckDBExists(name string) bool {
	for _, db := range s.AllDatabases {
		if db.Name == name {
			return true
		}
	}
	return false
}
