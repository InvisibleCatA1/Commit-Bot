package main

import (
	"fmt"
	"os"
)

func main() {
	var dir string

	fmt.Print("Enter the directory that the commit bot will use: ")
	fmt.Scanln(&dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Print("That dir does not exist")
		os.Exit(0)
	}

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		fmt.Print("Directory: " + dir)
	}
}
