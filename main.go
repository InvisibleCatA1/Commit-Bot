package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
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
		fmt.Println("Directory: " + dir)
		var cmd string
		var run bool = true
		var commitMsg string
		for run {
			fmt.Print("Enter command: ")
			fmt.Scanln(&cmd)
			if cmd == "quit" {
				fmt.Println("Exiting")
				run = false
			} else if cmd == "commit" {
				fmt.Print("Print the message for commit (can not be nothing): ")
				fmt.Scanln(&commitMsg)
				cmd := exec.Command("git", "commit", "-a", "-m", "\""+commitMsg+"\"")
				var outb, errb bytes.Buffer
				cmd.Stdout = &outb
				cmd.Stderr = &errb
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("out:", outb.String(), "err:", errb.String())

			}
		}
	}
}
