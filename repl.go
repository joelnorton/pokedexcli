package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	mv := strings.Split(text, " ")
	var output []string
	for _, w := range mv {
		if w != "" {
			output = append(output, w)
		}
	}
	return output
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		o := cleanInput(scanner.Text())
		c, ok := cfg.cliCommands[o[0]]
		if ok {
			if len(o) == 1 {
				o = append(o, "")
			}
			err := c.callback(cfg, o[1])
			if err != nil {
				fmt.Printf("error was found: %v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
