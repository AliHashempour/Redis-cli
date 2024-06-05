package handler

import (
	"go-memory/internal/repository"
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
