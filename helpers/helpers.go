package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

const CliName string = "pokedex"

var Commands map[string]cliCommand = map[string]cliCommand{
	"help":  {Name: "Help", Description: "Displays help information about pokedex.", Callback: displayHelp},
	"clear": {Name: "Clear", Description: "Clears the terminal screen", Callback: clearScreen},
}

func displayHelp() error {
	fmt.Printf(
		"Welcome to %v! These are the available commands: \n",
		CliName,
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

func PrintPrompt() {
	fmt.Print(CliName, "> ")
}

func PrintUnknown(text string) {
	fmt.Print(text, " : Command not found")
}

func CleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
