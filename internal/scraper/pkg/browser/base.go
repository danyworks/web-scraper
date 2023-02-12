package browser

import (
	"fmt"
	"github.com/daniyalibrahim/notybot/internal/scraper/common"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/cdp"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
	"log"
)

type Devices int
type BrowserOptions int

const (
	Desktop Devices = iota
	GalaxyFold
	IphoneX
	IPad
)
const (
	WithExtensions BrowserOptions = iota
	DevTools
	UserDataDir
	WithHead
	WithTrace
)

func LauncherString(options ...BrowserOptions) string {
	var opts string
	var runner = launcher.New()
	log.Printf("Initializing driver...")
	for _, option := range options {
		switch option {
		case WithExtensions:
			runner = runner.Set("load-extension", common.EXT_CHROME_PATH)
		case DevTools:
			runner = runner.Devtools(true)
		case UserDataDir:
			runner = runner.UserDataDir(common.ROD_USER_DATA_DIR)
		case WithHead:
			runner = runner.Headless(false)
		case WithTrace:
			runner = runner.Set("-rod", "show,slow=1s,trace")
		}
	}
	opts = runner.MustLaunch()
	return opts
}

func SetBrowserDisplaySize(device Devices, driver *rod.Browser) *rod.Browser {
	switch device {
	case Desktop:
		log.Printf("Initializing Desktop...")
		driver.DefaultDevice(devices.Clear)
		return driver
	case GalaxyFold:
		log.Printf("Initializing Galaxy Fold...")
		driver.DefaultDevice(devices.GalaxyFold)
		return driver
	case IphoneX:
		log.Printf("Initializing Iphone X...")
		driver.DefaultDevice(devices.IPhoneX)
		return driver
	case IPad:
		log.Printf("Initializing iPad...")
		driver.DefaultDevice(devices.IPad)
		return driver
	}
	driver.DefaultDevice(devices.Clear)
	return driver
}

func InitCDPTraffic() *cdp.Client {
	cdp := cdp.New().
		// Here we can customize how to log the requests, responses, and events transferred between Rod and the browser.
		Logger(utils.Log(func(args ...interface{}) {
			switch v := args[0].(type) {
			case *cdp.Request:
				fmt.Printf("id: %v\n", v.String())
			}
		})).
		Start(cdp.MustConnectWS(launcher.New().MustLaunch()))
	return cdp
}

func HijackRequests() *rod.HijackRouter {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	router := browser.HijackRequests()
	defer router.MustStop()

	router.MustAdd("*.js", func(ctx *rod.Hijack) {
		// Here we update the request's header. Rod gives functionality to
		// change or update all parts of the request. Refer to the documentation
		// for more information.
		ctx.Request.Req().Header.Set("My-Header", "test")

		// LoadResponse runs the default request to the destination of the request.
		// Not calling this will require you to mock the entire response.
		// This can be done with the SetXxx (Status, Header, Body) functions on the
		// ctx.Response struct.
		ctx.MustLoadResponse()

		// Here we append some code to every js file.
		// The code will update the document title to "hi"
		ctx.Response.SetBody(ctx.Response.Body() + "\n document.title = 'hi' ")
	})
	return router
	// Output: done
}
