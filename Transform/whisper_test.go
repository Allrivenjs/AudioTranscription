package Transform

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestWhisper(t *testing.T) {
	var wg sync.WaitGroup
	var te []Times
	tempDir, err := os.Getwd()
	tempDir = strings.Replace(tempDir, "/Transform", "", 1)
	paths := []string{
		tempDir + "/on_process_files/audioTest1/vocals.wav",
		tempDir + "/on_process_files/audioTest2/vocals.wav",
	}
	if err != nil {
		t.Error(err)
		return
	}
	for _, path := range paths {
		wg.Add(1)
		go func(path string) {
			fmt.Println("processing: ", path)
			defer wg.Done()
			ti := Times{
				Name:  getNameOfFilepath(path),
				Start: time.Now().UTC().Truncate(time.Millisecond),
			}
			defer func() {
				ti.End = time.Now().UTC().Truncate(time.Millisecond)
				ti.CalTime()
				te = append(te, ti)
			}()
			err := Whisper(path)
			if err != nil {
				fmt.Println("Error al usar Whisper", err)
			}
		}(path)
	}
	wg.Wait()
	for _, t := range te {
		fmt.Println(t.Name, ": ", t.Cal)
	}
	defer func() {
		err = ExportData(tempDir+"/statistics/whisper.csv", te)
		if err != nil {
			t.Error(err)
		}
	}()
}

func TestTimerWhisper(t *testing.T) {
	callTime("/statistics/whisper.csv")
}
