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
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("error reading input: %v\n", err)
			}
			return
		}
		o := cleanInput(scanner.Text())
		if len(o) == 0 {
			continue
		}
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
