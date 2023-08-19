package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	pokeapi "github.com/vudyaabhinav/pokedexcli/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

const CliName string = "pokedex"

var Commands map[string]cliCommand = map[string]cliCommand{
	"help": {
		Name:        "Help",
		Description: "Displays help information about pokedex.",
		Callback:    displayHelp,
	},
	"clear": {
		Name:        "Clear",
		Description: "Clears the terminal screen",
		Callback:    clearScreen,
	},
	"map": {
		Name:        "Map",
		Description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map will display the next 20 locations, and so on. The idea is that the map command will let us explore the world of Pokemon.",
		Callback:    getMapLocations,
	},
	"mapb": {
		Name:        "Map Back",
		Description: "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
		Callback:    mapb,
	},
}

func displayHelp(cfg *pokeapi.config) error {
	fmt.Printf(
		"Welcome to %v! These are the available commands: \n",
		CliName,
	)
	fmt.Println("map     - Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map will display the next 20 locations, and so on. The idea is that the map command will let us explore the world of Pokemon.")
	fmt.Println("mapb    - Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.")
	fmt.Println("clear   - Clear the terminal screen")
	fmt.Println("exit    - Exits Pokedex ")
	return nil
}

func clearScreen(cfg *pokeapi.config) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

func getMapLocations(cfg *pokeapi.config) error {
	// locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	locationsResp, err := pokeapi.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func mapb(cfg *pokeapi.config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	// locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	locationResp, err := pokeapi.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
