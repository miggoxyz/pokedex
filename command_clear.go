package main

import "fmt"

func commandClear(cfg *config) error {
	fmt.Print("\033[2J\033[H")
	return nil
}