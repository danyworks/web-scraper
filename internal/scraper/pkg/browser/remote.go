package browser

import (
	"github.com/daniyalibrahim/notybot/internal/scraper/common"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"log"
)

func InitRemoteBrowser(device Devices) *rod.Browser {
	la, err := launcher.NewManaged(common.RemoteBrowserURL)
	if err != nil {
		log.Panicf("Error while connecting to remote browser %v", err)
	}
	// You can also set any flag remotely before you launch the remote driver.
	// Available flags: https://peter.sh/experiments/chromium-command-line-switches
	la.Set("disable-gpu").Delete("disable-gpu")
	la.Set("rod", "show,slow=1s,trace")
	la.Set("load-extension", "D:\\go-workspace\\toyscraper\\scraper\\pkg\\plugins\\cookie-editor-chrome-1.10.1")
	// Launch with headful mode
	la.Headless(false).XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16")
	driver := rod.New().Client(la.MustClient()).MustConnect()
	driver = SetBrowserDisplaySize(device, driver)
	// You may want to start a server to watch the screenshots of the remote driver.
	launcher.Open(driver.ServeMonitor(""))
	log.Printf("Remote driver initialized...")
	return driver
}
