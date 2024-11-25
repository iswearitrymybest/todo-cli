package command

import (
	"flag"
	"strconv"
	"strings"
	"testing"
	"todo-cli/internal/todo"
)

func TestAddCommand(t *testing.T) {
	args := []string{"-add", "Test task"}
	todos := todo.Todos{}

	flagSet := flag.NewFlagSet("test", flag.ExitOnError)
	addFlag := flagSet.String("add", "", "Add a todo item")

	flagSet.Parse(args)

	if *addFlag != "" {
		todos.Add(*addFlag)
	}

	if len(todos.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(todos.Tasks))
	}

	if todos.Tasks[0].Title != "Test task" {
		t.Errorf("expected title 'Test task', got '%s'", todos.Tasks[0].Title)
	}
}

func TestInvalidEditCommand(t *testing.T) {

	args := []string{"-edit", "invalid_format"}
	todos := todo.Todos{}
	todos.Add("Existing task")

	flagSet := flag.NewFlagSet("test", flag.ExitOnError)
	editFlag := flagSet.String("edit", "", "Edit a todo by id and specify new title id:new_title")

	flagSet.Parse(args)

	if *editFlag != "" {
		parts := strings.SplitN(*editFlag, ":", 2)
		if len(parts) == 2 {
			id, err := strconv.Atoi(parts[0])
			if err == nil && id >= 0 && id < len(todos.Tasks) {
				todos.Edit(id, parts[1])
			}
		}
	}

	if len(todos.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(todos.Tasks))
	}

	if todos.Tasks[0].Title != "Existing task" {
		t.Errorf("expected title 'Existing task', got '%s'", todos.Tasks[0].Title)
	}
}
