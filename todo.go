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
		_id, _ := strconv.ParseUint(fields[0], 10, 32)
		if uint(_id) > id { id = uint(_id) }
		done, _ := strconv.ParseBool(fields[2])
		todos = append(todos, Task{uint(_id), fields[1], done})
	}
}

func show(){
	for i, task := range todos {
		fmt.Printf("[ %02d # %03d ] %s\n", i, task.id, task.name)
	}
}
func add() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter task: ")
	input, _ := reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(input), ",")
	id++
	task := Task{ id, words[0], false }
	todos = append(todos, task)
	fmt.Println(todos)
	save()
	show()
}

func del() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter id: ")
	input, _ := reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(input), ",")
	_id, _ := strconv.ParseUint(words[0], 10, 64)
	var index int = -1
	for i, task := range todos {
		if uint(_id) == task.id {
			index = i
			break
		}
	}

	if(index != -1){
		var new_todos []Task
		if index == 0 {
			new_todos = todos[1:]
		}else {
			new_todos = append(todos[0:index], todos[index + 1:]...)
		}
		todos = new_todos
		save()
		show()
	}
}

func command() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter command: ")
	input, _ := reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(input), ",")
	cmd := words[0]
	
	fmt.Println("Command:", cmd)

	switch cmd {
	case "s", "show": show()
	case "a", "add": add()
	case "d", "del": del()
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
