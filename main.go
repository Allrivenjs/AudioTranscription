package main

import (
	"AudioTranscription/Transform"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os/exec"
)

func main() {
	app := fiber.New()
	// Verificar si FFmpeg está instalado en el sistema
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		panic(fmt.Errorf("FFmpeg no está instalado en tu sistema"))
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/create-temp-folder", func(c *fiber.Ctx) error {
		//filepath := "on_process_files/audioTest1.mp3"
		filepath := "on_process_files/"

		//err := Transform.CutAudio(filepath, "parts", 15)
		//err := Transform.Spleeter(filepath)
		path, err := Transform.GetAllFilesPath(filepath)
		if err != nil {
			fmt.Println("Error al obtener la duración del archivo:", err)
		}
		fmt.Println(path)
		return c.SendString("Hello, World 👋! ")
	})

	panic(app.Listen(":3000"))
}
