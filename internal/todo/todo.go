package todo

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Status      bool   `json:"status"`
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos struct {
	Tasks  []Todo
	nextID int
}

func (t *Todos) Add(title string) {
	todo := Todo{
		ID:          t.nextID,
		Title:       title,
		Status:      false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	t.Tasks = append(t.Tasks, todo)
	t.nextID++
}

func (t *Todos) Complete(id int) error {
	task, index, err := t.findByID(id)
	if err != nil {
		return err
	}

	if !task.Status {
		completionTime := time.Now()
		t.Tasks[index].CompletedAt = &completionTime
	}

	t.Tasks[index].Status = !task.Status
	return nil
}

func (t *Todos) Edit(id int, title string) error {
	_, index, err := t.findByID(id)
	if err != nil {
		return err
	}
	t.Tasks[index].Title = title
	return nil
}

func (t *Todos) Delete(id int) error {
	_, index, err := t.findByID(id)
	if err != nil {
		return err
	}
	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)
	return nil
}

func (t *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("ID", "Title", "Status", "Created At", "Completed At")

	for _, task := range t.Tasks {
		completed := "✖"
		if task.Status {
			completed = "✔"
		}
		completedAt := "N/A"
		if task.CompletedAt != nil {
			completedAt = task.CompletedAt.Format(time.RFC3339)
		}
		table.AddRow(strconv.Itoa(task.ID), task.Title, completed, task.CreatedAt.Format(time.RFC3339), completedAt)
	}

	table.Render()
}
func (t *Todos) findByID(id int) (Todo, int, error) {
	for i, task := range t.Tasks {
		if task.ID == id {
			return task, i, nil
		}
	}
	return Todo{}, -1, errors.New("task not found")
}

func (t *Todos) UpdateNextID() {
	maxID := 0
	for _, task := range t.Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	t.nextID = maxID + 1
}
