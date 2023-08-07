package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	helpers "github.com/vudyaabbhinav/pokedexcli/helpers"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	helpers.PrintPrompt()
	for reader.Scan() {
		text := helpers.CleanInput(reader.Text())

		if strings.EqualFold("exit", text) {
			return
		}

		if command, ok := helpers.Commands[text]; ok {
			command.Callback()
		} else {
			helpers.PrintUnknown(text)
		}
		helpers.PrintPrompt()
	}
	fmt.Println()
}
