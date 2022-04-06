package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/TwiN/go-color"
)

var userIn = "[" + color.Blue + "Q" + color.Reset + "] "
var error = "[" + color.Red + "E" + color.Reset + "] "
var info = "[" + color.Green + "I" + color.Reset + "] "

func main() {
	dir := input(userIn + "Enter the directory that the commit bot will use: ")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Print(error + "That dir does not exist")
		os.Exit(0)
	}

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		fmt.Println(info + "Directory: " + dir)
		var cmd string
		var run bool = true
		var commitMsg string
		for run {
			cmd = input(userIn + "Enter command: ")
			if cmd == "quit" {
				fmt.Println(info + "Exiting")
				run = false
			} else if cmd == "commit" {
				commitMsg = input(userIn + "Print the message for commit (can not be nothing): ")
				commit(commitMsg)
			} else if cmd == "help" {
				fmt.Println(info + "commit - commit to a git repo (muse already be initalized)")
				fmt.Println(info + "quit - exists the program")
				fmt.Println(info + "help - displays this page")
				fmt.Println(info + "regular commit - commits every minute (delay can be specified)")
			} else if cmd == "regular commit" {
				times := input(userIn + "Number of times to commit: ")
				delay := input(userIn + "delay between commits (minutes): ")
				for i := 0; i < strToInt(times); i++ {
					commit("Repeat commit")
					fmt.Println(info + " Waiting " + delay + " minutes unitl next commit")
					time.Sleep(time.Duration(strToInt(delay)) * time.Minute)

				}

			} else {
				fmt.Println(error + "Unkown command: " + cmd + "; type help for list of commands")
			}
		}
	}
}
func strToInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err == nil {
		fmt.Println(err)
	}
	return intVal
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
	fmt.Println(info + "Pushed")
}
func commit(commitMsg string) {
	if commitMsg == "" {
		fmt.Println(error + "commit msg cannot be none")
		return
	}
	cmd := exec.Command("git", "commit", "-a", "-m", commitMsg)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(outb.String())
	push()
}
func input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	return line
}
