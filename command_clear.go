package main

import "fmt"

func commandClear() error {
	fmt.Print("\033[2J\033[H")
	return nil
}