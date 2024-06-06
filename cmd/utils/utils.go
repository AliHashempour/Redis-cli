package utils

import (
	"fmt"
	"go-memory/handler"
	"go-memory/internal/repository"
	"go-memory/internal/storage"
	"strings"
)

type CommandProcessor struct {
	Store          *storage.Storage
	CommandHandler *handler.CommandHandler
}

func NewCommandProcessor(store *storage.Storage, handler *handler.CommandHandler) *CommandProcessor {
	return &CommandProcessor{
		Store:          store,
		CommandHandler: handler,
	}
}

func (cp *CommandProcessor) ProcessCommand(input string) string {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return ""
	}

	switch parts[0] {
	case "set":
		return cp.processSet(parts)
	case "get":
		return cp.processGet(parts)
	case "del":
		return cp.processDel(parts)
	case "keys":
		return cp.processKeys(parts)
	case "list":
		return cp.CommandHandler.ListDatabasesCommand(cp.Store)
	case "use":
		return cp.processUse(parts)
	case "exit":
		return "Exiting..."
	default:
		return "Unknown command"
	}
}

func (cp *CommandProcessor) processSet(parts []string) string {
	if len(parts) != 3 {
		return "Usage: set <key> <value>"
	}
	return cp.CommandHandler.SetCommand(parts[1], parts[2])
}

func (cp *CommandProcessor) processGet(parts []string) string {
	if len(parts) != 2 {
		return "Usage: get <key>"
	}
	return fmt.Sprint(cp.CommandHandler.GetCommand(parts[1]))
}

func (cp *CommandProcessor) processDel(parts []string) string {
	if len(parts) != 2 {
		return "Usage: del <key>"
	}
	return cp.CommandHandler.DeleteCommand(parts[1])
}

func (cp *CommandProcessor) processKeys(parts []string) string {
	if len(parts) != 2 {
		return "Usage: keys <regex_pattern>"
	}
	return cp.CommandHandler.KeysCommand(parts[1])
}

func (cp *CommandProcessor) processUse(parts []string) string {
	if len(parts) < 2 {
		return "Usage: use <dbName>"
	}
	dbName := parts[1]
	_, err := cp.Store.UseDatabase(dbName)
	if err != nil {
		return fmt.Sprintf("Failed to switch database: %s", err)
	}
	cp.CommandHandler = handler.NewCommandHandler(repository.NewMemoryRepository(cp.Store.CurrentDatabase))
	return fmt.Sprintf("Switched to database: %s", dbName)
}
