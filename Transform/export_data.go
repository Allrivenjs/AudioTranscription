package Transform

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func ExportData(filename string, data []Times) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir el encabezado CSV
	header := []string{"Name", "Start", "End", "Cal"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Escribir datos CSV
	for _, item := range data {
		record := []string{
			item.Name,
			item.Start.Format(time.RFC3339),
			item.End.Format(time.RFC3339),
			fmt.Sprintf("%v", item.Cal),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	fmt.Printf("Datos exportados a %s\n", filename)
	return nil
}
