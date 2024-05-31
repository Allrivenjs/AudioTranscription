package services

import (
	"AudioTranscription/serve/storage"
	"AudioTranscription/serve/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// CuterAudioDuration is the duration of the audio to cut.
	CuterAudioDuration = 30
)

// CuterAudio cuts audio every 15 seconds.
func CuterAudio(src, dst string) (map[int]string, error) {
	list := make(map[int]string)
	// get the duration of the audio
	file, err := storage.GetFile(src)
	if err != nil {
		return list, err
	}
	// remove .wav extension
	fileInfo, err := file.Stat()
	if err != nil {
		return list, err
	}
	newName := fileInfo.Name()[:len(fileInfo.Name())-4]
	// cut the audio every 15 seconds
	_ = storage.CreateFolderIntoTemp(newName)
	name := fmt.Sprintf("%soutput_%%d_file.wav", fmt.Sprintf("%s%s/", dst, newName))
	err = util.SplitAudio(src, name, CuterAudioDuration)
	if err != nil {
		return list, err
	}
	pathSave := fmt.Sprintf("%s%s/", dst, newName)
	var files []os.DirEntry
	files, err = storage.GetFilesOnDir(pathSave)
	if err != nil {
		fmt.Println("Error getting files on dir")
		return nil, err
	}
	for _, file := range files {
		key := strings.Split(file.Name(), "_")[1]
		num, _ := strconv.Atoi(key)
		list[num] = fmt.Sprintf("%s%s", pathSave, file.Name())
	}
	return list, nil
}
