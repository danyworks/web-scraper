package api

import (
	"github.com/daniyalibrahim/notybot/api/handlers"
	"github.com/daniyalibrahim/notybot/cron"
	proc "github.com/daniyalibrahim/notybot/internal/processing"
	"log"
	"net/http"
	"time"
)

func job() {
	log.Println("Execute at: ", time.Now())
}

func GetMuxAPI() *http.ServeMux {
	log.Print("Initializing Rest Endpoints...")
	mux := http.NewServeMux()

	jobChan := make(chan proc.Job, 100)

	go proc.Worker(jobChan)
	go cron.CronJob("0 * * * *", job)

	mux.Handle("/", handlers.Webhook(jobChan))
	mux.HandleFunc("/webhook", handlers.Webhook(jobChan))
	mux.HandleFunc("/health", handlers.Health)

	return mux
}
