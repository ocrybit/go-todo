package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"strings"
	"strconv"
	"path/filepath"
	"io/ioutil"
)

type Task struct {
	id uint
	name string
	done bool
	done_at uint
}

var todos []Task
var id uint = 0

func save() error {
	dir, _ := filepath.Abs(".todos")
	file, _ := filepath.Abs(".todos/todos.txt")
	_ = os.MkdirAll(dir, 0755)
	var txt string
	for _, task := range todos {
		txt += fmt.Sprintf("%d,%s,%t,%d\n", task.id, task.name, task.done, task.done_at)
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
		var done_at uint = 0
		if len(fields) > 3 {
			_done_at, _ := strconv.ParseUint(fields[3], 10, 32)
			done_at = uint(_done_at)
		}
		todos = append(todos, Task{uint(_id), fields[1], done, done_at})
	}
}

func show(){
	undone := false
	for i, task := range todos {
		if !task.done {
			fmt.Printf("[ %02d # %03d ] %s\n", i, task.id, task.name)
		}else{
			undone = true
		}
	}
	if(undone){
		fmt.Println("-------------------------------------- [ done ]")
		for i, task := range todos {
			if task.done {
				done_at := time.Unix(int64(task.done_at), 0)
				fmt.Printf("[ %02d # %03d ] %s %02d/%02d \n", i, task.id, task.name, done_at.Month(), done_at.Day())
			}
		}
	}
}
func add() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter task: ")
	input, _ := reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(input), ",")
	id++
	task := Task{ id, words[0], false, 0 }
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

func complete() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter id: ")
	input, _ := reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(input), ",")
	_id, _ := strconv.ParseUint(words[0], 10, 64)
	var index int = -1
	for i, task := range todos {
		if uint(_id) == task.id {
			index = i
			if todos[i].done {
				todos[i].done_at = 0
			} else {
				todos[i].done_at = uint(time.Now().Unix())
			}
			todos[i].done = !todos[i].done
			break
		}
	}
	if(index != -1){
		save()
		show()
	}
}

func trash() {
	var new_todos []Task
	for _, task := range todos {
		if !task.done {
			new_todos = append(new_todos, task)
		}
	}
	todos = new_todos
	save()
	show()
}

func command() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nenter command: ")
	input, _ := reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(input), ",")
	cmd := words[0]
	
	switch cmd {
	case "s", "show": show()
	case "c", "complete": complete()
	case "t", "trash": trash()
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
	show()
	command()
}
