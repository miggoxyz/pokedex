package main

import "fmt"

func commandClear(cfg *config, args ...string) error {
	fmt.Print("\033[2J\033[H")
	return nil
}