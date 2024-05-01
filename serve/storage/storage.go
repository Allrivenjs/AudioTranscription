package storage

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
	"strings"
)

type Storage interface {
	CreateFile(path string, file *multipart.FileHeader) (string, error)
	CreateAudioFile(path string, file *multipart.FileHeader) (string, error)
}

var baseRoute = fmt.Sprintf("%s", "/storage/app/")
var baseTemp = fmt.Sprintf("%s", "temp/")

type storage struct {
	baseRoute string
	ctx       *fiber.Ctx
}

var pathCurrent = fmt.Sprintf("%s", func() string {
	ospath, err := os.Getwd()
	if err != nil {
		panic("Error getting current working directory")
	}
	return ospath
}())

func GetPathCurrent() string {
	return pathCurrent
}

func GetBaseRoute() string {
	return baseRoute
}

func GetBaseTemp() string {
	return baseTemp
}

func CreateFolderTemp() error {
	err := CreateFolder("temp")
	if err != nil {
		return err
	}
	return nil
}

func CreateFolder(path string) error {
	finalPath := fmt.Sprintf("%s%s%s", GetPathCurrent(), baseRoute, path)
	for _, err := os.Stat(finalPath); os.IsNotExist(err); {
		err := os.Mkdir(finalPath, 0755)
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
	err := CreateFolder("audio")
	newPath, err := s.CreateFile(fmt.Sprintf("%s/%s", "audio", path), file)
	if err != nil {
		return "", err
	}
	return newPath, nil
}

func GetFile(path string) (*os.File, error) {
	newPath := fmt.Sprintf("%s", path)
	file, err := os.Open(newPath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetFileInformation(path string) (os.FileInfo, error) {
	newPath := fmt.Sprintf("%s/%s", pathCurrent, path)
	file, err := os.Stat(newPath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func DeleteFile(path string) error {
	newPath := fmt.Sprintf("%s", path)
	err := os.Remove(newPath)
	if err != nil {
		return err
	}
	return nil

}

func NewStorage(ctx *fiber.Ctx) Storage {
	return &storage{baseRoute: baseRoute, ctx: ctx}
}
