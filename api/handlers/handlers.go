package handlers

import (
	proc "github.com/daniyalibrahim/notybot/internal/processing"
	"log"
)

func Webhook(jobChan chan<- proc.Job) {
	var job proc.Job
	str := "asd"
	switch str {
	case "queued":
		//job = proc.NewJob()
		log.Printf(" queued!!")
	case "completed":
		//job = proc.NewJob(payload)
		log.Printf("completed !!")
	default:
		log.Printf("Not implemented!!")
		return
	}
	jobChan <- job
}
