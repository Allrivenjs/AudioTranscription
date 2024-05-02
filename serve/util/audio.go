package util

import (
	"fmt"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

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
	err := ffmpeg_go.Input(src, ffmpeg_go.KwArgs{"ss": startTime}).Output(dst, ffmpeg_go.KwArgs{"to": endTime}).Silent(true).Run()
	if err != nil {
		return fmt.Errorf("error: %w out: %s", err, dst)
	}
	return nil
}

func SplitAudio(src, dst string, duration int) error {
	err := ffmpeg_go.Input(src).
		Output(dst, ffmpeg_go.KwArgs{"f": "segment", "segment_time": duration, "c": "copy"}).Silent(false).Silent(true).
		Run()
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}
