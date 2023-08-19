package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	c "github.com/vudyaabbhinav/pokedexcli/commands"
	helpers "github.com/vudyaabbhinav/pokedexcli/helpers"
	pokeapi "github.com/vudyaabbhinav/pokedexcli/pokeapi"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	helpers.PrintPrompt()
	for reader.Scan() {
		text := helpers.CleanInput(reader.Text())

		if strings.EqualFold("exit", text) {
			return
		}

		if command, ok := c.Commands[text]; ok {
			pokeClient := pokeapi.NewClient(5 * time.Second)
			// cfg := &pokeapi.config{
			cfg := &pokeapi.config{
				// pokeapiClient: pokeClient,
			}
			command.Callback(cfg)
		} else {
			helpers.PrintUnknown(text)
		}
		helpers.PrintPrompt()
	}
	fmt.Println()
}
