package jobs

import (
	"AudioTranscription/serve/db"
	"fmt"
	"time"
)

type Dispatch interface {
	Listen()
}

type dispatchRepository struct {
	job Job
}

func (d dispatchRepository) Listen() {
	fmt.Println("Listening for jobs")
	for {
		// sleep for 1 second
		time.Sleep(1 * time.Second)
		err := d.job.Run()
		if err != nil {
			panic(err)
		}
	}
}

func Init(conn db.Connection) Dispatch {
	job := NewJobManager(conn)
	return &dispatchRepository{
		job: job,
	}
}
