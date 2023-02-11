package handlers

import (
	"encoding/json"
	"github.com/daniyalibrahim/notybot/internal/models"
	proc "github.com/daniyalibrahim/notybot/internal/processing"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ok := models.NewResponse("success", "healthy", http.StatusOK)
	json.NewEncoder(w).Encode(ok)
}

func Webhook(jobChan chan<- proc.Job) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		resp := &models.Response{}
		//req.Body = http.MaxBytesReader(w, req.Body, 1048576)
		decoder := json.NewDecoder(req.Body)
		payload := &proc.JobPayload{}
		err := decoder.Decode(payload)
		if err != nil {
			response := ParseErrors(w, err)
			json.NewEncoder(w).Encode(response)
			return
		}
		resp.Status = "success"
		resp.Message = payload.Action
		w.Header().Set("Content-Type", "application/json")

		var job proc.Job

		switch payload.Action {
		case "queued":
			job = proc.NewJob(payload)
		case "completed":
			job = proc.NewJob(payload)
		case "in_progress":
			job = proc.NewJob(payload)
			resp := models.NewResponse("success", payload.Action, http.StatusOK)
			json.NewEncoder(w).Encode(resp)
		default:
			resp := models.NewResponse("success", payload.Action, http.StatusOK)
			log.Printf("Not implemented!!")
			json.NewEncoder(w).Encode(resp)
			return
		}

		jobChan <- job
		json.NewEncoder(w).Encode(resp)
	}
}
