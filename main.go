package main

import (
	"fmt"
	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

func main() {
	var modelpath = "models/ggml-small.bin"
	var samples []float32

	// load the model
	model, err := whisper.New(modelpath)
	if err != nil {
		panic(err)
	}
	defer model.Close()

	// process the samples
	context, err := model.NewContext()
	if err != nil {
		panic(err)
	}
	if err := context.Process(samples, nil, nil); err != nil {
		panic(err)
	}

	// Print out the results
	for {
		segment, err := context.NextSegment()
		if err != nil {
			break
		}
		fmt.Printf("[%6s->%6s] %s\n", segment.Start, segment.End, segment.Text)
	}
}
