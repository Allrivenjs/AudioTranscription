package Transform

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestSpleeter(t *testing.T) {
	var wg sync.WaitGroup
	var te []Times
	tempDir, err := os.Getwd()
	tempDir = strings.Replace(tempDir, "/Transform", "", 1)
	paths := []string{
		tempDir + "/on_process_files/audioTest1.mp3",
		tempDir + "/on_process_files/audioTest2.mp3",
	}
	if err != nil {
		t.Error(err)
		return
	}
	//for 1 to 1000
	run := func(ii int) {
		wg.Add(1)
		for i := 0; i < 4; i++ {
			for _, path := range paths {
				fmt.Println("iteration: ", ii, "path: ", path, "i: ", i, " ---")
				func(path string) {
					ti := Times{
						Name:  getNameOfFilepath(path),
						Start: time.Now().UTC().Truncate(time.Millisecond),
					}
					defer func() {
						ti.End = time.Now().UTC().Truncate(time.Millisecond)
						ti.CalTime()
						te = append(te, ti)
					}()
					err := Spleeter(path)
					if err != nil {
						fmt.Println("Error al usar Spleeter", err)
					}
				}(path)
			}
		}
		defer wg.Done()
	}
	var wg2 sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			run(i)
		}()
		if i%2 == 0 {
			wg2.Wait()
		}
	}
	wg.Wait()
	for _, v := range te {
		fmt.Println(v.Name, ": ", v.Cal)
	}
	defer func(te []Times) {
		err = ExportData(tempDir+"/statistics/spleeter.csv", te)
		if err != nil {
			return
		}
	}(te)
}

func TestTimerSpleeter(t *testing.T) {
	callTime("/statistics/spleeter.csv")
}
