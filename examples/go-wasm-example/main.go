package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello from Go WebAssembly!")
	fmt.Println("Type 'help' for available commands.")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		switch {
		case line == "help":
			fmt.Println("Available commands:")
			fmt.Println("  echo <text>  - print text")
			fmt.Println("  quit         - exit")
			fmt.Println("  help         - show this help")
		case strings.HasPrefix(line, "echo "):
			fmt.Println(strings.TrimPrefix(line, "echo "))
		case line == "quit" || line == "exit":
			fmt.Println("Goodbye!")
			return
		case line == "":
			// ignore empty input
		default:
			fmt.Fprintf(os.Stderr, "unknown command: %s\n", line)
		}
	}
}
