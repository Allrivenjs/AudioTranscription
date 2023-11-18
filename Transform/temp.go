package Transform

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Temp crea una carpeta temporal en el directorio actual
// y devuelve la ruta de la carpeta temporal
// Ejemplo de uso:
// tempPath, err := Temp("audio", "parts")
//
//	if err != nil {
//		fmt.Println("Error al crear la carpeta temporal:", err)
//	}
//
// fmt.Println(tempPath)
// donde "audio" es el nombre de la carpeta temporal
// y "parts" es el nombre de la subcarpeta
// La salida será algo como esto:
// /home/usuario/go/src/AudioTranscription/temp/02_01_2006_1612345678_audio
func Temp(name string, t string) (string, error) {
	// Obtener la ubicación del directorio temporal del sistema
	tempDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error al obtener el directorio actual:", err)
		return "", err
	}
	tempDir = strings.Join([]string{tempDir, "temp"}, "/")
	// Crear un nombre único para la carpeta temporal
	folderName := uniqueName(name)
	t = strings.Replace(t, " ", "_", -1)
	// Combinar la ubicación del directorio temporal, el nombre de la carpeta y "audio"
	folderPath := filepath.Join(tempDir, folderName, t)

	// Crear la carpeta temporal (incluyendo subcarpetas)
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Crear un archivo dentro de la carpeta temporal (por ejemplo, "output.txt")
	// Reemplaza "output.txt" con el nombre de tu archivo de audio
	outputPath := filepath.Join(folderPath, ".gitignore")
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			fmt.Println("Error al cerrar el archivo:", err)
		}
	}(outputFile)
	outputPath = folderPath
	return outputPath, nil
}

func uniqueName(name string) string {
	currentTime := time.Now()
	return fmt.Sprintf(
		"%s_%s_%s",
		replace(currentTime.Format("02-01-2006")),
		replace(strconv.FormatInt(currentTime.Unix(), 10)),
		name)
}

func replace(text string) string {
	return strings.Replace(text, "-", "_", -1)
}
