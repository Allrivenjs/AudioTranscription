package main

import (
	"fmt"
	"os"
	"os/exec"
)

func sh(c string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// audioToWav converts audio to wav format -> https://ffmpeg.org/documentation.html
func audioToWav(src, dst string) error {
	out, err := sh(fmt.Sprintf("ffmpeg -i %s -format s16le -ar 16000 -ac 1 -acodec pcm_s16le %s", src, dst))
	if err != nil {
		return fmt.Errorf("Error converting audio to wav: %s -> %s", out, err)
	}
	return nil
}

// cutSilence cuts silence from the audio file -> https://github.com/dunossauro/videomaker-helper
func cutSilence(src, dst string) error {
	out, err := sh(fmt.Sprintf("vmh cut-silences %s %s", src, dst))
	if err != nil {
		return fmt.Errorf("error: %w out: %s", err, out)
	}
	return nil
}
