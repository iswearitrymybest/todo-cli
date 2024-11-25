package storage

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) (T, error) {
	fileData, err := os.ReadFile(s.FileName)
	if os.IsNotExist(err) {
		return *data, nil
	}
	if err != nil {
		return *data, err
	}
	if err := json.Unmarshal(fileData, data); err != nil {
		return *data, err
	}
	return *data, nil
}
