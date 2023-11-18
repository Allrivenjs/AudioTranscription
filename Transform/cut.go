package Transform

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func CutAudio(filepath string, outputPrefix string, partSeconds int) error {
	// Obtener información del archivo de audio (duración)
	// Aquí deberás usar una biblioteca de procesamiento de audio en Go para obtener la duración.
	duration, err := GetAudioDuration(filepath)

	if err != nil {
		fmt.Println("Error al obtener la duración del archivo:", err)
	}

	cuts := partSizeInSeconds(duration, partSeconds)

	_, err = makeFilesForCuts(cuts, filepath, outputPrefix)
	if err != nil {
		return err
	}
	return nil
}

func getNameOfFilepath(filepath string) string {
	split := strings.Split(filepath, "/")
	return strings.Replace(split[len(split)-1], ".mp3", "", -1)
}

// makeFilesForCuts crea los archivos de audio cortados
// cuts: los puntos de corte
// filepath: la ruta del archivo de audio
// outputPrefix: el prefijo del archivo de salida
// Ejemplo: si el archivo de audio se llama "audio.mp3" y el prefijo es "parte", el archivo de salida se llamará "parte_0.mp3"
// Esta func es comandada con una gorutina
func makeFilesForCuts(cuts []string, filepath string, outputPrefix string) (string, error) {
	//crear una carpeta temporal
	temp, err := Temp(getNameOfFilepath(filepath), "parts")
	if err != nil {
		fmt.Println("Error al crear la carpeta temporal:", err)
		return "", err
	}
	var wg sync.WaitGroup
	// Dividir el archivo de audio en partes utilizando el comando Python
	for i, cut := range cuts {
		wg.Add(1)
		go func(i int, cut string) {
			defer wg.Done()
			outputFile := strings.Join([]string{temp, fmt.Sprintf(BaseOutputParts, outputPrefix, i)}, "")
			//"python", filepath, "--path", filepath, "--output", outputFile, cut
			cmd := exec.Command("bash", "-c", fmt.Sprintf(CutAudioCMD, filepath, cut, outputFile))
			fmt.Println("Ejecutando comando:", cmd.String())
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println("Error al ejecutar el comando:", err)
			}
		}(i, cut)
	}
	return temp, nil
}

func GetAudioDuration(filepath string) (float64, error) {
	// Ejecutar el comando ffmpeg
	//"ffmpeg -i %s 2>&1 | grep Duration | awk '{print $2}' | tr -d ,"
	cmd := exec.Command("bash", "-c", fmt.Sprintf(getDuration, filepath))
	durationStr, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// Limpiar la cadena de duración de caracteres no deseados
	durationStrc := strings.TrimSpace(string(durationStr))
	// Parsear la duración a un valor de tiempo
	duration, err := time.Parse("15:04:05.00", durationStrc)
	if err != nil {
		return 0, err
	}

	// Obtener la duración en segundos
	seconds := float64(duration.Second() + duration.Minute()*60 + duration.Hour()*3600)
	return seconds, nil
}

func partSizeInSeconds(duration float64, partSizeInSeconds int) []string {
	// Calcular el número de partes basado en el tamaño deseado
	numParts := int(duration / float64(partSizeInSeconds))
	// Calcular los puntos de corte
	var cuts []string
	for i := 0; i < numParts; i++ {
		start := i * partSizeInSeconds
		end := (i + 1) * partSizeInSeconds
		if i == numParts-1 {
			end = int(duration)
		}
		cut := fmt.Sprintf("%d -t %d", start, end)
		cuts = append(cuts, cut)
	}
	return cuts
}
