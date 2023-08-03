package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

const cliName string = "pokedex"

var commands map[string]cliCommand = map[string]cliCommand{
	"help":  {name: "Help", description: "Displays help information about pokedex.", callback: displayHelp},
	"clear": {name: "Clear", description: "Clears the terminal screen", callback: clearScreen},
}

func printPrompt() {
	fmt.Print(cliName, "> ")
}

func printUnknown(text string) {
	fmt.Print(text, " : Command not found")
}

func displayHelp() error {
	fmt.Printf(
		"Welcome to %v! These are the available commands: \n",
		cliName,
	)
	fmt.Println("help    - Show available commands")
	fmt.Println("clear   - Clear the terminal screen")
	fmt.Println("exit    - Exits Pokedex ")
	return nil
}

func clearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())

		if strings.EqualFold("exit", text) {
			return
		}

		if command, ok := commands[text]; ok {
			command.callback()
		} else {
			printUnknown(text)
		}
		printPrompt()
	}
	fmt.Println()
}
