package main

import (
	"bufio"
	"fmt"
	"go-memory/handler"
	"go-memory/internal/repository"
	"go-memory/internal/storage"
	"os"
	"strings"
)

func main() {

	store := storage.NewStorage()
	db := store.CurrentDatabase

	commandHandler := handler.NewCommandHandler(repository.NewMemoryRepository(db))

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("In-app Key-Value Store")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break // Exit the loop on EOF or read error.
		}

		command := scanner.Text()
		parts := strings.Fields(command)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "set":
			if len(parts) != 3 {
				fmt.Println("Usage: set <key> <value>")
			} else {
				fmt.Println(commandHandler.SetCommand(parts[1], parts[2]))
			}
		case "get":
			if len(parts) != 2 {
				fmt.Println("Usage: get <key>")
			} else {
				fmt.Println(commandHandler.GetCommand(parts[1]))
			}
		case "del":
			if len(parts) != 2 {
				fmt.Println("Usage: del <key>")
			} else {
				fmt.Println(commandHandler.DeleteCommand(parts[1]))
			}
		case "keys":
			if len(parts) != 2 {
				fmt.Println("Usage: keys <regex>")
			} else {
				regex := parts[1]
				// Placeholder for keys function
				fmt.Printf("Searching keys with pattern %s\n", regex)
			}
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %s\n", err)
	}
}
