package utils

import (
	"encoding/json"
	"fmt"
	"go-memory/handler"
	"go-memory/internal/repository"
	"go-memory/internal/storage"
	"os"
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
		return cp.processList()
	case "use":
		return cp.processUse(parts)
	case "dump":
		return cp.DumpDatabase(parts)
	case "load":
		return cp.LoadDatabase(parts)
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

func (cp *CommandProcessor) processList() string {
	return cp.CommandHandler.ListDatabasesCommand(cp.Store)
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

func (cp *CommandProcessor) DumpDatabase(parts []string) string {
	if len(parts) != 3 {
		return "Usage: dump <db_name> <path>"
	}
	dbName := parts[1]
	filepath := parts[2]

	db, exists := cp.Store.CheckDBExists(dbName)
	if !exists {
		return fmt.Sprintf("Database %s does not exist.", dbName)
	}
	data, err := json.Marshal(db.Data)
	if err != nil {
		return fmt.Sprintf("Failed to serialize database data: %s", err)
	}
	// Write data to file
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return fmt.Sprintf("Failed to write to file: %s", err)
	}
	return "Database dumped successfully."
}

func (cp *CommandProcessor) LoadDatabase(parts []string) string {
	if len(parts) != 3 {
		return "Usage: load <path> <db_name>"
	}
	filepath := parts[1]
	dbName := parts[2]

	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Sprintf("Failed to read file: %s", err)
	}
	var deserializedData map[string]interface{}
	err = json.Unmarshal(data, &deserializedData)
	if err != nil {
		return fmt.Sprintf("Failed to deserialize data: %s", err)
	}
	db, _ := cp.Store.UseDatabase(dbName)
	db.Data = deserializedData

	return fmt.Sprintf("Database %s loaded successfully.", dbName)
}
