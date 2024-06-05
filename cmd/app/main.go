package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
				key := parts[1]
				value := parts[2]
				// Placeholder for set function
				fmt.Printf("Setting %s to %s\n", key, value)
			}
		case "get":
			if len(parts) != 2 {
				fmt.Println("Usage: get <key>")
			} else {
				key := parts[1]
				// Placeholder for get function
				fmt.Printf("Getting value for %s\n", key)
			}
		case "del":
			if len(parts) != 2 {
				fmt.Println("Usage: del <key>")
			} else {
				key := parts[1]
				// Placeholder for delete function
				fmt.Printf("Deleting %s\n", key)
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
