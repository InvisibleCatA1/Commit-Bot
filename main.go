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
	fmt.Println(`
_________                        .__  __           __________        __   
\_   ___ \  ____   _____   _____ |__|/  |_         \______   \ _____/  |_ 
/    \  \/ /  _ \ /     \ /     \|  \   __\  ______ |    |  _//  _ \   __\
\     \___(  <_> )  Y Y  \  Y Y  \  ||  |   /_____/ |    |   (  <_> )  |  
 \_______/\____/|__|_ |_ /__|_|  /__||__|           |______ _/\____/|__|  
		                                                         `)
	fmt.Println(info + "https://github.com/InvisibleCatA1/Commit-Bot")
	fmt.Println(info + "by InvisibleCat")
	dir := input("Enter the directory that the commit bot will use: ")

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
			cmd = input("Enter command: ")
			if cmd == "quit" {
				fmt.Println(info + "Exiting")
				run = false
			} else if cmd == "commit" {
				commitMsg = input("Print the message for commit (can not be nothing): ")
				commit(commitMsg)
			} else if cmd == "help" {
				fmt.Println(info + "commit - commit to a git repo (muse already be initalized)")
				fmt.Println(info + "quit - exists the program")
				fmt.Println(info + "help - displays this page")
				fmt.Println(info + "regular commit - commits every minute (delay can be specified)")
				fmt.Println(info + "init - auto inits a repo")
			} else if cmd == "regular commit" {
				times := input("number of times to commit: ")
				delay := input("delay between commits (minutes): ")
				timesAsInt := strToInt(times)
				i := 0
				for i < timesAsInt {
					i += 1
					time.Sleep(time.Duration(strToInt(delay)) * time.Minute)
					commit("Repeat commit by - https://github.com/InvisibleCatA1/Commit-Bot")
					fmt.Println(info + " Waiting " + delay + " minutes unitl next commit")
					push()
				}
			} else if cmd == "init" {
				createReadme := input("Create a README.md (y/n): ")
				if createReadme == "y" {
					readMeText := input("default README text: ")
					readmeCmd := exec.Command("echo", readMeText, ">>", "README.md")
					var outb, errb bytes.Buffer
					readmeCmd.Stdout = &outb
					readmeCmd.Stderr = &errb
					err := readmeCmd.Run()
					if err != nil {
						log.Fatal(err)
					}

				}

				initCmd := exec.Command("git", "init")
				originUrl := input("Url of the repo (ex. https://github.com/InvisibleCatA1/Commit-Bot.git): ")
				if originUrl == "" {
					fmt.Println(error + "origin Url can not be nothing")
					return
				}
				addOrigin := exec.Command("git", "add", "origin", originUrl)
				err := initCmd.Run()
				if err != nil {
					log.Fatal(err)
				}
				err = addOrigin.Run()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(info + "Pushing...")
				cmd := exec.Command("git", "push", "-u", "origin", "main")
				var outb, errb bytes.Buffer
				cmd.Stdout = &outb
				cmd.Stderr = &errb
				err = cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(info + "Done")

			} else {
				fmt.Println(error + "Unkown command: " + cmd + "; type help for list of commands")
			}
		}
	}
}
func strToInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err == nil {
		fmt.Print()
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
	fmt.Print(userIn + prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	return line
}
