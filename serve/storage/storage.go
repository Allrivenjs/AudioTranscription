package storage

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
	"strings"
)

type Storage interface {
	CreateFolder(path string) error
	CreateFile(path string, file *multipart.FileHeader) (string, error)
	CreateAudioFile(path string, file *multipart.FileHeader) (string, error)
	GetFile(path string) ([]byte, error)
}

var baseRoute = fmt.Sprintf("%s", "/storage/app/")

type storage struct {
	baseRoute string
	ctx       *fiber.Ctx
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

func (s storage) CreateFile(path string, file *multipart.FileHeader) (string, error) {
	fmt.Println("Current working directory")
	ospath, _ := os.Getwd()
	newPath := fmt.Sprintf("%s%s", s.baseRoute, strings.Replace(path, " ", "_", -1))
	err := s.ctx.SaveFile(file, fmt.Sprintf("%s/%s", ospath, newPath))
	if err != nil {
		return "", err

	}
	return newPath, nil
}

func (s storage) CreateAudioFile(path string, file *multipart.FileHeader) (string, error) {

	err := s.CreateFolder("audio")
	newPath, err := s.CreateFile(fmt.Sprintf("%s/%s", "audio", path), file)
	if err != nil {
		return "", err
	}
	return newPath, nil
}

func (s storage) GetFile(path string) ([]byte, error) {
	newPath := fmt.Sprintf("%s", path)
	file, err := os.ReadFile(newPath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewStorage(ctx *fiber.Ctx) Storage {
	return &storage{baseRoute: baseRoute, ctx: ctx}
}
