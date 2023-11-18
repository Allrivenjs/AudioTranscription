package Transform

import (
	"fmt"
	"os"
	"os/exec"
)

func ConvertM4AToMP3(inputPath string, outputPath string) error {
	// Verificar si FFmpeg está instalado en el sistema
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		return fmt.Errorf("FFmpeg no está instalado en tu sistema")
	}

	// Ejecutar FFmpeg para convertir el archivo M4A a MP3
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-acodec", "libmp3lame", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("Conversión de %s a %s completada.\n", inputPath, outputPath)
	return nil
}
