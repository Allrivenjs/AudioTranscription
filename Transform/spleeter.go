package Transform

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Spleeter(pathTemp string) error {
	//get name of file
	name := getNameOfFilepath(pathTemp)
	outputpath := strings.Replace(pathTemp, "/"+name+".mp3", "", -1)
	// Comando Spleeter
	fmt.Printf(fmt.Sprintf(SpleeterCMD, outputpath, pathTemp))
	cmd := exec.Command("bash", "-c", fmt.Sprintf(SpleeterCMD, outputpath, pathTemp))

	// Redirigir la salida estándar y estándar de error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Ejecutar el comando
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Spleeter completado con éxito")
	return nil
}
