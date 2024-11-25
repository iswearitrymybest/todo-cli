package main

import (
	"os"
	"testing"
	"todo-cli/internal/storage"
	"todo-cli/internal/todo"
)

func TestIntegration(t *testing.T) {
	fileName := "integration_test.json"
	defer os.Remove(fileName)

	store := storage.NewStorage[todo.Todos](fileName)
	todos := todo.Todos{}

	todos.Add("Task 1")
	todos.Add("Task 2")
	store.Save(todos)

	var loadedTodos todo.Todos
	store.Load(&loadedTodos)

	if len(loadedTodos.Tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(loadedTodos.Tasks))
	}

	if loadedTodos.Tasks[0].Title != "Task 1" || loadedTodos.Tasks[1].Title != "Task 2" {
		t.Errorf("tasks not loaded correctly")
	}
}
