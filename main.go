package main

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY")
	fmt.Println("This is free software, and you are welcome to redistribute it under certain conditions")
	fmt.Println("")
	var email string
	var password string
	var screen string
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Print("\nScreen: ")
	fmt.Scan(&screen)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		password = string(bytePassword)
	}
	inloop := true
	var prompt string
	reader := bufio.NewReader(os.Stdin)
	var prompt_history []string
	for inloop {
		sh, err := get_shell(email, password, screen)
		if err != nil {
			fmt.Println("Connection Error")
			inloop = false
		} else {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			for i := 0; i < len(sh.Screen); i++ {
				if len(sh.Screen)-1 == i {
					fmt.Print(sh.Screen[i])
				} else {
					fmt.Println(sh.Screen[i])
				}
			}
			if !(sh.Active) {
				inloop = false
			} else {
				text, _ := reader.ReadString('\n')
				prompt = strings.TrimSpace(text)
				_ = append(prompt_history, prompt)
				if strings.Split(prompt, " ")[0] == "disconnect" {
					inloop = false
				} else if strings.Split(prompt, " ")[0] == "reconnect" {
					screen = strings.Split(prompt, " ")[1]
				} else {
					send_command(email, password, screen, prompt)
				}
			}
		}

	}
	fmt.Print("\n")
}
