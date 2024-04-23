package storage

import (
	"fmt"
	"os"
)

type Storage interface {
	CreateFolder(path string) error
	CreateFile(path string, body []byte) (string, error)
	CreateAudioFile(path string, body []byte) (string, error)
	GetFile(path string) ([]byte, error)
}

const baseRoute = "../storage/app/"

type storage struct {
	baseRoute string
}

func (s storage) CreateFolder(path string) error {
	for _, err := os.Stat(fmt.Sprintf("%s%s", s.baseRoute, path)); os.IsNotExist(err); {
		err := os.Mkdir(fmt.Sprintf("%s%s", s.baseRoute, path), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s storage) CreateFile(path string, body []byte) (string, error) {
	newPath := fmt.Sprintf("%s%s", s.baseRoute, path)
	err := os.WriteFile(newPath, body, 0644)
	if err != nil {
		return "", err
	}
	return newPath, nil
}

func (s storage) CreateAudioFile(path string, body []byte) (string, error) {
	file, err := s.CreateFile(fmt.Sprintf("%s/%s", "audio", path), body)
	if err != nil {
		return "", err
	}
	return file, nil
}

func (s storage) GetFile(path string) ([]byte, error) {
	newPath := fmt.Sprintf("%s%s", s.baseRoute, path)
	file, err := os.ReadFile(newPath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewStorage() Storage {
	return &storage{baseRoute: baseRoute}
}
