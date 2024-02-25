package main

import (
	"fmt"
	"time"
)

// srtTimestamp returns a string with the time formatted as SRT
func srtTimestamp(t time.Duration) string {
	return fmt.Sprintf("%02d:%02d:%02d,%03d",
		t/time.Hour,
		(t%time.Hour)/time.Minute,
		(t%time.Minute)/time.Second,
		(t%time.Second)/time.Millisecond,
	)
}
