package service

import (
	"encoding/json"
	"github.com/daniyalibrahim/notybot/internal/scraper/logic"
	"github.com/daniyalibrahim/notybot/internal/scraper/pkg/browser"
	"log"
)

type TaskOutput interface {
	GetResults() []byte
}

func ScanPage() []byte {
	log.Printf("[SERVICE] Scanning Page for Elements")
	driver := browser.InitBrowser(false, false, false, browser.IphoneX, browser.WithExtensions, browser.WithHead)

	page := driver.MustPage("https://google.com")
	results := logic.ScanPage(page, true)
	// encode to json
	//struct to json
	json, err := json.Marshal(results)
	if err != nil {
		log.Println(err)
	}
	driver.MustClose()
	return json
}
