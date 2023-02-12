package browser

import (
	"github.com/daniyalibrahim/notybot/internal/scraper/common"
	"github.com/go-rod/rod"
	"log"
)

func InitBrowser(incognito, hijack, remote bool, device Devices, options ...BrowserOptions) *rod.Browser {
	var driver *rod.Browser
	if remote {
		log.Printf("Initializing remote driver from %v...", common.RemoteBrowserURL)
		driver = InitRemoteBrowser(device)
	} else {
		log.Printf("Initializing local driver...")
		opts := LauncherString(options...)
		driver = rod.New().ControlURL(opts).MustConnect()
	}
	if incognito {
		log.Printf("with incognito")
		driver.MustIncognito()
	}
	if hijack {
		log.Printf("with hijack...")
		router := HijackRequests()
		go router.Run()
	}
	defer log.Printf("Browser initialized...")
	driver = SetBrowserDisplaySize(device, driver)
	return driver
}
