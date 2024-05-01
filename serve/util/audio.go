package util

import (
	"fmt"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"os"
	"os/exec"
)

// sh executes shell command.
func sh(c string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Env = os.Environ()
	o, err := cmd.CombinedOutput()
	return string(o), err
}

func Sh(c string) (string, error) {
	return sh(c)
}

// AudioToWav converts audio to wav for transcribe.
func AudioToWav(src, dst string) error {
	err := ffmpeg_go.Input(src).Output(dst).OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

// CutSilences cuts silences from audio.
func CutSilences(src, dst string, startTime, endTime int) error {
	err := ffmpeg_go.Input(src, ffmpeg_go.KwArgs{"ss": startTime}).Output(dst, ffmpeg_go.KwArgs{"t": endTime}).OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return fmt.Errorf("error: %w out: %s", err, dst)
	}
	return nil
}
