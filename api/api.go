package api

import (
	"github.com/daniyalibrahim/notybot/api/handlers"
	"github.com/daniyalibrahim/notybot/cron"
	proc "github.com/daniyalibrahim/notybot/internal/processing"
	"log"
	"time"
)

func job() {
	log.Println("Execute at: ", time.Now())
}

func GetMuxAPI() {
	log.Print("Initializing Rest Endpoints...")

	jobChan := make(chan proc.Job, 100)

	go proc.Worker(jobChan)
	go cron.CronJob("0 * * * *", job)

	handlers.Webhook(jobChan)

}
