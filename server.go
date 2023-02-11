package main

import (
	"github.com/daniyalibrahim/notybot/api"
	"log"
	"net/http"
	"time"
)

func main() {

	srv := &http.Server{
		Addr: ":8080",
		// Good practice to set timeouts
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.GetMuxAPI(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Error while running server %v", err)
	}

	log.Println("shutting down...")

}
