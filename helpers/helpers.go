package helpers

import (
	"fmt"
	"strings"
)

const CliName string = "pokedex"

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
