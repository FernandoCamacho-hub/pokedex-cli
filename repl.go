package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print(" > ")

        scanner.Scan()
        text := scanner.Text()
        cleaned := cleanInput(text)
        if len(cleaned) == 0 {
            continue
        }

        commandName := cleaned[0]
        availableCommands := getCommands()

        command, ok := availableCommands[commandName]
        if !ok {
            fmt.Println("Invalid command")
            continue
        }
        err := command.callback(cfg)
        if err != nil {
            fmt.Println(err)
        }
    }

}

type cliCommand struct {
    name string
    description string
    callback func(*config) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Prints the help menu",
            callback: callbackHelp,
        },
        "exit": {
            name: "exit",
            description: "Exits the pokedex",
            callback: callbackExit,
        },
        "map": {
            name: "map",
            description: "List some location areas",
            callback: callbackMap,
        },
        "mapb": {
            name: "mapb",
            description: "List some previous location areas",
            callback: callbackMapBack,
        },
    }
}

func cleanInput(str string) []string {
    lowered := strings.ToLower(str)
    words := strings.Fields(lowered)
    return words
}
