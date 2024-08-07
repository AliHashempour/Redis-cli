package storage

import (
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

	newDb := database.NewDatabase(name)
	s.AllDatabases = append(s.AllDatabases, newDb)
	return newDb, nil
}

func (s *Storage) UseDatabase(name string) (*database.Database, error) {
	db, exists := s.CheckDBExists(name)
	if exists {
		s.CurrentDatabase = db
		return db, nil
	}
	// Database not found, create a new one
	newDb, err := s.AddDatabase(name)
	if err != nil {
		return nil, err
	}
	s.CurrentDatabase = newDb
	return newDb, nil
}

func (s *Storage) CheckDBExists(name string) (*database.Database, bool) {
	for _, db := range s.AllDatabases {
		if db.Name == name {
			return db, true
		}
	}
	return nil, false
}

func (s *Storage) ListAllDatabases() []string {
	var names []string
	for _, db := range s.AllDatabases {
		names = append(names, db.Name)
	}
	return names
}
