package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"path/filepath"
	"io/ioutil"
)

type Task struct {
	id uint
	name string
	done bool
	
}

var todos []Task
var id uint = 0

func save() error {
	dir, _ := filepath.Abs(".todos")
	file, _ := filepath.Abs(".todos/todos.txt")
	_ = os.MkdirAll(dir, 0755)
	var txt string
	for _, task := range todos {
		txt += fmt.Sprintf("%d,%s,%t\n", task.id, task.name, task.done)
	}
	_ = ioutil.WriteFile(file, []byte(txt), 0644)
	return nil
}

func load() {
	file, _ := filepath.Abs(".todos/todos.txt")
	data, _ := ioutil.ReadFile(file)
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		id, _ := strconv.ParseUint(fields[0], 10, 32)
		done, _ := strconv.ParseBool(fields[2])
		todos = append(todos, Task{uint(id), fields[1], done})
	}
}

func add() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter task: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	words := strings.Split(strings.TrimSpace(input), ",")
	id++
	task := Task{ id, words[0], false }
	todos = append(todos, task)
	fmt.Println(todos)
	save()
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
	case "a", "add": add()
	case "q", "quit":
		fmt.Println("Bye!")
		return
	default:
		fmt.Println("unknown command:", cmd)
	}
	command()
}

func main() {
	load()
	command()
}
