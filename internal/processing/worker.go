package processing

import (
	"log"
)

func Worker(jobChan <-chan Job) {
	for job := range jobChan {
		go process(job)
	}
}

func process(job Job) {
	log.Printf("job: %v", job)
	switch job.Action {
	case "queued":
		log.Printf("action: %v", job.Action)
	case "completed":
		log.Printf("action: %v", job.Action)
	case "in_progress":
		log.Printf("action: %v", job.Action)
	default:
		log.Printf("Not processed")
		return
	}

	log.Print("Processed")
}
