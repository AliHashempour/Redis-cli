package main

import (
	"bufio"
	"fmt"
	"go-memory/cmd/utils"
	"go-memory/handler"
	"go-memory/internal/repository"
	"go-memory/internal/storage"
	"os"
)

func main() {
	store := storage.NewStorage()
	db := store.CurrentDatabase
	commandHandler := handler.NewCommandHandler(repository.NewMemoryRepository(db))
	processor := utils.NewCommandProcessor(store, commandHandler)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("In-app Key-Value Store")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		command := scanner.Text()
		response := processor.ProcessCommand(command)
		if response == "Exiting..." {
			fmt.Println(response)
			break
		} else if response != "" {
			fmt.Println(response)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %s\n", err)
	}
}
