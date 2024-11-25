package storage

import (
	"os"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	type TestData struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	fileName := "test_storage.json"
	defer os.Remove(fileName)

	store := NewStorage[TestData](fileName)

	data := TestData{Name: "John", Age: 30}
	if err := store.Save(data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var loadedData TestData
	_, err := store.Load(&loadedData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if loadedData.Name != "John" || loadedData.Age != 30 {
		t.Errorf("expected data {Name: John, Age: 30}, got {Name: %s, Age: %d}", loadedData.Name, loadedData.Age)
	}
}
