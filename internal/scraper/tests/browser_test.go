package tests

import (
	"github.com/daniyalibrahim/notybot/internal/scraper/pkg/browser"
	"strings"
	"testing"
)

var (
	URLS = []string{
		"http://facebook.com",
		"http://instagram.com",
		"http://tiktok.com",
		"http://linkedin.com",
		"http://twitter.com",
	}
)

func TestLocalBrowser(t *testing.T) {
	keyword := "some"
	driver := browser.InitBrowser(false, false, false, browser.Desktop, browser.WithExtensions, browser.WithHead)
	defer driver.MustClose()
	page := driver.MustPage("https://www.wikipedia.org/")
	page.MustWaitLoad().MustScreenshot("a.png")
	el := page.MustElementX("/html/body/div[3]/form/fieldset/div/input")
	el.MustInput(keyword)
	text, err := el.Text()
	if err != nil {
		panic(err)
	}
	if !strings.Contains(text, keyword) {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", text, keyword)
	}
}
func TestRemoteBrowser(t *testing.T) {
	keyword := "some"
	driver := browser.InitBrowser(false, false, true, browser.Desktop, browser.WithExtensions, browser.WithHead)
	defer driver.MustClose()
	page := driver.MustPage("https://www.wikipedia.org/")
	page.MustWaitLoad().MustScreenshot("a.png")
	el := page.MustElementX("/html/body/div[3]/form/fieldset/div/input")
	el.MustInput(keyword)
	text, err := el.Text()
	if err != nil {
		panic(err)
	}
	if !strings.Contains(text, keyword) {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", text, keyword)
	}
}
