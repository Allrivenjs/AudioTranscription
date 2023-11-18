package Transform

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Whisper(pathTemp string) error {
	name := getNameOfFilepath(pathTemp)
	outputpath := strings.Replace(pathTemp, name, "", -1)
	// Comando Whisper
	cmd := exec.Command("bash", "-c", "/home/allrivenjs/anaconda3/bin/python3.11 -m "+fmt.Sprintf(WhispetCMD, pathTemp, outputpath))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("Whisper completado con éxito")
	return nil
}
