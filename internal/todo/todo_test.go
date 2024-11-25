package todo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	todos := Todos{}
	todos.Add("Test task 1")
	todos.Add("Test task 2")

	if len(todos.Tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(todos.Tasks))
	}

	if todos.Tasks[0].Title != "Test task 1" {
		t.Errorf("expected title 'Test task 1', got %s", todos.Tasks[0].Title)
	}

	if todos.Tasks[0].ID != 0 || todos.Tasks[1].ID != 1 {
		t.Errorf("expected IDs 0 and 1, got %d and %d", todos.Tasks[0].ID, todos.Tasks[1].ID)
	}
}

func TestComplete(t *testing.T) {
	todos := Todos{}
	todos.Add("Test task 1")

	err := todos.Complete(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !todos.Tasks[0].Status {
		t.Errorf("expected task to be marked as completed")
	}

	if todos.Tasks[0].CompletedAt == nil {
		t.Errorf("expected CompletedAt to be set")
	}
}

func TestDelete(t *testing.T) {
	todos := Todos{}
	todos.Add("Task 1")
	todos.Add("Task 2")

	err := todos.Delete(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(todos.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(todos.Tasks))
	}

	if todos.Tasks[0].Title != "Task 2" {
		t.Errorf("expected remaining task title to be 'Task 2', got %s", todos.Tasks[0].Title)
	}
}

func TestFindByID(t *testing.T) {
	todos := Todos{}
	todos.Add("Task 1")
	todos.Add("Task 2")

	_, _, err := todos.findByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, _, err = todos.findByID(999)
	if err == nil {
		t.Fatalf("expected error for non-existent ID")
	}
}

func TestUpdateNextID(t *testing.T) {
	todos := Todos{
		Tasks: []Todo{
			{ID: 10, Title: "Task 1"},
			{ID: 20, Title: "Task 2"},
		},
	}
	todos.UpdateNextID()

	if todos.nextID != 21 {
		t.Errorf("expected nextID to be 21, got %d", todos.nextID)
	}
}
