package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Task struct {
	name string
	done bool
	
}

var todos []Task

func add() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter task: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	words := strings.Split(strings.TrimSpace(input), ",")
	task := Task{ words[0], false }
	todos = append(todos, task)
	fmt.Println(todos)
}

func command() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter command: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	
	words := strings.Split(strings.TrimSpace(input), ",")
	cmd := words[0]
	
	fmt.Println("Command:", cmd)

	switch cmd {
	case "a", "add":
		add()
	default:
		fmt.Println("unknown command:", cmd)
	}
	command()
}

func main() {
	command()
}
