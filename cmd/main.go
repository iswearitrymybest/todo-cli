package main

import (
	"fmt"
	"todo-cli/internal/command"
	"todo-cli/internal/storage"
	"todo-cli/internal/todo"
)

func main() {
	todocli := todo.Todos{}
	store := storage.NewStorage[todo.Todos]("todos.json")

	if _, err := store.Load(&todocli); err != nil {
		fmt.Println("Warning: could not load todos from file:", err)
	} else {
		todocli.UpdateNextID()
	}

	cmdFlags := command.NewCmdFlags()

	if cmdFlags.Help {
		cmdFlags.ExecuteCMD(&todocli)
		return
	}

	cmdFlags.ExecuteCMD(&todocli)

	if err := store.Save(todocli); err != nil {
		fmt.Println("Error: could not save todos to file:", err)
	}

	if !cmdFlags.List && !cmdFlags.Help {
		fmt.Println("\nUpdated Todo List:")
		todocli.Print()
	}
}
