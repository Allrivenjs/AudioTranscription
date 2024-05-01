package services

import (
	"AudioTranscription/serve/storage"
	"AudioTranscription/serve/util"
	"fmt"
	"github.com/go-audio/wav"
	"strings"
)

const (
	// CuterAudioDuration is the duration of the audio to cut.
	CuterAudioDuration = 15
)

// CuterAudio cuts audio every 15 seconds.
func CuterAudio(src, dst string) ([]string, error) {
	var list []string
	// get the duration of the audio
	file, err := storage.GetFile(src)
	if err != nil {
		return list, err
	}
	// calculate the max duration of the audio
	maxDuration, err := wav.NewDecoder(file).Duration()
	// calculate the number of 15 seconds intervals
	intervals := int(maxDuration.Seconds() / CuterAudioDuration)
	// remove .wav extension
	newName := strings.Split(file.Name(), "/")[3]
	newName = newName[:len(newName)-4]
	// cut the audio every 15 seconds
	errorsNames := make(chan error, intervals)
	created := make(chan string, intervals)
	go func() {
		for i := 0; i < intervals; i++ {
			name := fmt.Sprintf("%s%d.wav", fmt.Sprintf("%s/%s", dst, newName), i)
			err := util.CutSilences(src, name, i*CuterAudioDuration, (i+1)*CuterAudioDuration)
			if err != nil {
				errorsNames <- err
				return
			}
			created <- name
		}
		// end the goroutine
		close(created)
	}()
	close(errorsNames)
	for n := range errorsNames {
		if n != nil {
			// remove the files created
			for name := range created {
				err := storage.DeleteFile(name)
				if err != nil {
					return list, err
				}
			}
			return list, err
		}
	}
	for name := range created {
		list = append(list, name)
	}

	return list, nil
}
