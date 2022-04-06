package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var dir string

	dir = input("Enter the directory that the commit bot will use: ")

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
			cmd = input("Enter command: ")
			if cmd == "quit" {
				fmt.Println("Exiting")
				run = false
			} else if cmd == "commit" {
				commitMsg = input("Print the message for commit (can not be nothing): ")
				cmd := exec.Command("git", "commit", "-a", "-m", "\""+commitMsg+"\"")
				var outb, errb bytes.Buffer
				cmd.Stdout = &outb
				cmd.Stderr = &errb
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("out:", outb.String())
				push()

			}
		}
	}
}
func push() {
	cmd := exec.Command("git", "push")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Push output: ", outb.String())
}
func input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	return line
}
