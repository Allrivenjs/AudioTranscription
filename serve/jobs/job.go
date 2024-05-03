package jobs

import (
	"AudioTranscription/serve/cloudflare"
	"AudioTranscription/serve/db"
	"AudioTranscription/serve/models"
	"AudioTranscription/serve/services"
	"AudioTranscription/serve/storage"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"sort"
	"sync"
	"time"
)

var Instance *jobManager

type Job interface {
	Register(typeQueue string, payload string) error
	Completed(job *models.JobModel) error
	Run() error
	VerifyJobs() *[]*models.JobModel
	ExecTranscription(job *models.JobModel)
}

type jobManager struct {
	db *gorm.DB
}

func (j *jobManager) Register(typeQueue string, payload string) error {
	fmt.Println("Registering job")
	job := &models.JobModel{
		Queue:      typeQueue,
		Payload:    payload,
		Attempts:   0,
		Completed:  false,
		UnResolved: false,
		CreatedAt:  uint(time.Now().UnixNano()),
	}
	register := j.db.Save(job)
	err := register.Error
	if err != nil {
		return err

	}
	fmt.Println("Job registered", typeQueue)
	return nil
}

func (j *jobManager) Completed(job *models.JobModel) error {
	job.Completed = true
	persis := j.db.Save(job)
	err := persis.Error
	if err != nil {
		return err

	}
	return nil
}

func (j *jobManager) UnResolved(job *models.JobModel) error {
	job.UnResolved = true
	persis := j.db.Save(job)
	err := persis.Error
	if err != nil {
		return err
	}
	return nil
}

var wg sync.WaitGroup

func (j *jobManager) Run() error {
	// get all jobs
	jobs := j.VerifyJobs()
	if jobs == nil {
		return nil
	}
	wg.Add(len(*jobs))
	for _, job := range *jobs {
		go j.ExecTranscription(job)
	}
	wg.Wait()
	return nil
}

func (j *jobManager) ExecTranscription(job *models.JobModel) {
	defer wg.Done()
	// process job
	// convert payload to json
	var transcription models.Transcription
	err := json.Unmarshal([]byte(job.Payload), &transcription)
	if err != nil {
		fmt.Printf("Error unmarshalling json: %s", err.Error())
		j.failed(job)
		return
	}
	// verify if existed the folder
	_ = storage.CreateFolderTemp()

	path := fmt.Sprintf("%s%s%s", storage.GetPathCurrent(), storage.GetBaseRoute(), storage.GetBaseTemp())
	var files map[int]string
	locate := fmt.Sprintf("%s%s", storage.GetPathCurrent(), transcription.LocateFile)
	files, err = services.CuterAudio(locate, path)
	if err != nil {
		fmt.Printf("Error cutting audio: %s", err.Error())
		j.failed(job)
		return
	}
	numFiles := len(files)
	transcriptions := make([]chan cloudflare.Response, numFiles)

	keys := make([]int, 0, len(files))
	for k := range files {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var wg2 sync.WaitGroup
	for _, key := range keys {
		file := files[key]
		wg2.Add(1)
		transcriptions[key] = make(chan cloudflare.Response, 1)
		go func(key int, file string, ch chan cloudflare.Response, wg *sync.WaitGroup) {
			time.Sleep(5 * time.Microsecond)
			resp := cloudflare.CloudflareAI(key, file)
			fmt.Println("enviando respuesta")
			ch <- resp
			fmt.Println("respuesta enviada")
			fmt.Println("Done processing file: ", file)
			wg.Done()
		}(key, file, transcriptions[key], &wg2)
	}
	fmt.Println("Waiting for all files to be processed")
	wg2.Wait()
	fmt.Println("All files processed")
	transcriptionText := ""
	for _, key := range keys {
		response := <-transcriptions[key]
		transcriptionText = fmt.Sprintf("%s %s", transcriptionText, response.Result.Text)
	}

	fmt.Println("Transcription: ", transcriptionText)

	// save transcription
	transcription.Transcription = transcriptionText
	if len(transcriptionText) > 250 {
		transcription.SortTranscription = transcriptionText[:250]
	} else {
		transcription.SortTranscription = transcriptionText
	}
	err = j.db.Save(&transcription).Error
	if err != nil {
		fmt.Printf("Error saving transcription: %s", err.Error())
		j.failed(job)
		return
	}

	//mark job as completed
	err = j.Completed(job)
	if err != nil {
		fmt.Printf("Error marking job as completed: %s", err.Error())
		j.failed(job)
		return
	}
}

func (j *jobManager) failed(job *models.JobModel) {
	job.Attempts++
	if job.Attempts > 3 {
		// mark job as completed
		err := j.UnResolved(job)
		if err != nil {
			fmt.Printf("Error marking job as unresolved: %s", err.Error())
		}
		fmt.Printf("Job failed after 3 attempts: %s", job.ID)
	}
	err := j.db.Save(job).Error
	if err != nil {
		fmt.Printf("Error saving job: %s", err.Error())
	}
}

func (j *jobManager) VerifyJobs() *[]*models.JobModel {
	// get all jobs
	var jobs []*models.JobModel
	j.db.Find(&jobs, "completed = ?", false)
	if len(jobs) == 0 {
		return nil
	}
	return &jobs
}

func NewJobManager(coon db.Connection) Job {
	manager := &jobManager{db: coon.DB()}
	Instance = manager
	return manager
}

func GetInstance() Job {
	return Instance
}
