package repository

import (
	"errors"
	"go-memory/internal/database"
)

type Repository interface {
	SetKey(key string, value interface{}) error
	GetKey(key string) (interface{}, error)
	DeleteKey(key string) error
	RegexSearchKeys(pattern string) ([]string, error)
}

type MemoryRepository struct {
	db *database.Database
}

func NewMemoryRepository(db *database.Database) *MemoryRepository {
	return &MemoryRepository{db: db}
}

func (m *MemoryRepository) SetKey(key string, value interface{}) error {
	err := m.db.Set(key, value)
	if err != nil {
		return errors.New("error while setting key : " + err.Error())
	}
	return nil
}

func (m *MemoryRepository) GetKey(key string) (interface{}, error) {
	value, err := m.db.Get(key)
	if err != nil {
		return "", errors.New("error while getting key : " + err.Error())
	}
	return value, nil
}

func (m *MemoryRepository) DeleteKey(key string) error {
	err := m.db.Delete(key)
	if err != nil {
		return errors.New("error while deleting key : " + err.Error())
	}
	return nil
}
func (m *MemoryRepository) RegexSearchKeys(pattern string) ([]string, error) {
	search, err := m.db.RegexSearch(pattern)
	if err != nil {
		return nil, errors.New("error while getting key : " + err.Error())
	}
	return search, nil
}
