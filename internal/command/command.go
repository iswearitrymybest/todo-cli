package command

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"todo-cli/internal/todo"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
	Help   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a todo item")
	flag.IntVar(&cf.Del, "del", -1, "Specify id of todo item to delete")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by id and specify new title id:new_title")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify id of todo item to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todo items")
	flag.BoolVar(&cf.Help, "help", false, "Show this help message")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) ExecuteCMD(t *todo.Todos) {
	switch {
	case cf.Help:
		fmt.Println(`Todo CLI - A simple command-line Todo application

		Usage:
		  -add "title"          Add a new todo item
		  -del id               Delete a todo item by ID
		  -edit "id:new_title"  Edit the title of a todo item by ID
		  -toggle id            Toggle the completion status of a todo item by ID
		  -list                 List all todo items
		  -help                 Show this help message
		
		Examples:
		  todo-cli -add "Buy groceries"
		  todo-cli -list
		  todo-cli -toggle 1
		  todo-cli -edit "1:New title"
		  todo-cli -del 1
		`)

	case cf.Add != "":
		t.Add(cf.Add)

	case cf.Del >= 0:
		if err := t.Delete(cf.Del); err != nil {
			fmt.Println("Error:", err)
		}

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit format. Use id:new_title")
			return
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid id")
			return
		}

		if err := t.Edit(id, parts[1]); err != nil {
			fmt.Println("Error:", err)
		}

	case cf.Toggle >= 0:
		if err := t.Complete(cf.Toggle); err != nil {
			fmt.Println("Error:", err)
		}

	case cf.List:
		t.Print()

	default:
		fmt.Println("Invalid command. Use -help to see available commands.")
	}
}
