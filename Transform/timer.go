package Transform

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type Times struct {
	Name  string
	Start time.Time
	End   time.Time
	Cal   int64
}

func (t Times) CalTime() int64 {
	cal := t.End.Sub(t.Start).Milliseconds()
	t.Cal = cal
	return cal
}

func callTime(path string) {
	// Abrir el archivo CSV de entrada
	tempDir, err := os.Getwd()
	tempDir = strings.Replace(tempDir, "/Transform", "", 1)
	file, err := os.Open(tempDir + path)
	if err != nil {
		fmt.Println("Error al abrir el archivo CSV:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Crear un lector CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}

	// Crear un archivo CSV de salida
	outputFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error al crear el archivo de salida:", err)
		return
	}
	defer outputFile.Close()

	// Crear un escritor CSV para el archivo de salida
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Escribir el encabezado CSV en el archivo de salida
	header := []string{"Name", "Start", "End", "Cal"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error al escribir el encabezado CSV:", err)
		return
	}

	// Procesar los registros y calcular el tiempo en la columna "Cal"
	for _, record := range records {
		if len(record) != 4 {
			fmt.Println("Registro CSV incorrecto:", record, "len:", len(record))
			continue
		}

		name := record[0]
		startStr := record[1]
		endStr := record[2]

		// Validar que las cadenas de fecha y hora sean válidas antes de analizarlas
		startTime, err := time.Parse(time.RFC3339, startStr)
		if err != nil {
			fmt.Println("Error al analizar la hora de inicio:", err)
			continue
		}

		endTime, err := time.Parse(time.RFC3339, endStr)
		if err != nil {
			fmt.Println("Error al analizar la hora de finalización:", err)
			continue
		}

		duration := endTime.Sub(startTime).Milliseconds()

		// Crear un nuevo registro con el cálculo en "Cal"
		newRecord := Times{
			Name:  name,
			Start: startTime,
			End:   endTime,
			Cal:   duration,
		}

		// Escribir el registro en el archivo de salida
		outputRecord := []string{newRecord.Name, newRecord.Start.Format(time.RFC3339), newRecord.End.Format(time.RFC3339), fmt.Sprintf("%v", newRecord.Cal)}
		if err := writer.Write(outputRecord); err != nil {
			fmt.Println("Error al escribir el registro CSV de salida:", err)
		}
	}

	fmt.Println("Proceso completado. Resultados en output.csv")
}
