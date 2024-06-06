package handler

import (
	"encoding/json"
	"fmt"
	"go-memory/internal/repository"
	"go-memory/internal/storage"
)

// CommandHandler holds a repository to perform database operations.
type CommandHandler struct {
	Repo repository.Repository
}

// NewCommandHandler creates a new handler with a given repository.
func NewCommandHandler(repo repository.Repository) *CommandHandler {
	return &CommandHandler{Repo: repo}
}

func (h *CommandHandler) SetCommand(key string, value interface{}) string {
	err := h.Repo.SetKey(key, value)
	if err != nil {
		return err.Error()
	}
	return "OK"
}

func (h *CommandHandler) GetCommand(key string) interface{} {
	value, err := h.Repo.GetKey(key)
	if err != nil {
		return err.Error()
	}
	return value
}

func (h *CommandHandler) DeleteCommand(key string) string {
	err := h.Repo.DeleteKey(key)
	if err != nil {
		return err.Error()
	}
	return "OK"
}

func (h *CommandHandler) ListDatabasesCommand(s *storage.Storage) string {
	databases := s.ListAllDatabases()
	jsonData, err := json.Marshal(databases)
	if err != nil {
		return "Error formatting database list"
	}
	return string(jsonData)
}

func (h *CommandHandler) UseDatabaseCommand(s *storage.Storage, dbName string) string {
	_, err := s.UseDatabase(dbName)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Switched to or created and using database: %s", dbName)
}

func (h *CommandHandler) KeysCommand(pattern string) string {
	keys, err := h.Repo.RegexSearchKeys(pattern)
	if err != nil {
		return err.Error()
	}

	// Convert the keys list to JSON format
	jsonData, err := json.Marshal(keys)
	if err != nil {
		return err.Error()
	}

	return string(jsonData) // Return JSON formatted string
}
