package logic

import (
	"fmt"
	"github.com/daniyalibrahim/notybot/internal/scraper/common"
	"github.com/go-rod/rod"
	"log"
	"strings"
	"time"
)

func ScanPage(page *rod.Page, all bool) ScannerResults {
	var elements ScannerResults
	log.Printf("Scanning page...")
	if all {
		elements = ScanPageElements(page, common.HTML_TAGS...)
	}
	//PrintReport(page)
	log.Printf("Page scanned.")
	return elements
}

func LogElement(ele HTMLElement) {
	log.Printf("Element %v: name: %v visible: %v properties %v", ele.Index, ele.TagName, ele.VisibleText, ele.Properties)
}

func ScanPageElements(page *rod.Page, htmlTags ...string) ScannerResults {
	log.Printf("Scanning page for %v tags...", len(htmlTags))
	var results ScannerResults

	// find all text input elements
	for key, tag := range htmlTags {
		log.Printf("Scanning for all %v elements in body...", tag)
		for idx, row := range page.MustElements(tag) {
			var ele = HTMLElement{}
			ele.Index = idx
			ele.TagName = tag
			text, err := row.Text()
			if err != nil {
				log.Printf("Error getting text for %v at index %v", tag, key)
			}
			ele.VisibleText = text
			ele.Properties = ScanElementForProperties(row, common.HTML_PROPERTIES...)
			LogElement(ele)
			// log complete ele
			results.Data = append(results.Data, ele)
		}
	}

	log.Printf("Page elements scanned.")
	return results
}

func ScanElementForProperties(selector *rod.Element, properties ...string) []HTMLProperties {
	defer func() {
		if r := recover(); r != nil {
			//log.Printf("Recovered in Properties Scanner", r)
		}
	}()
	var result = []HTMLProperties{}
	for _, prop := range properties {
		proper, err := selector.Timeout(1 * time.Second).Property(prop)
		common.HandleError(err)
		// check if string is empty
		if !proper.Nil() {
			eleProps := HTMLProperties{prop, proper.String()}
			result = append(result, eleProps)
			//log.Printf("HTML Tag has property with key %v and %v value", prop, proper.String())
		}
	}
	log.Printf("%v Properties found", len(result))
	return result
}

func ScanPageXPATHForSelector(page *rod.Page, targetEl string) string {
	var xpath string
	log.Printf("Getting Xpath for %v...", targetEl)

	log.Printf("Xpath found.")
	return xpath
}

func PrintReport(page *rod.Page) {
	log.Printf("Printing report...")
	el := page.MustElement("#broken-image-dimensions.passed")
	for _, row := range el.MustParents("table").First().MustElements("tr:nth-child(n+2)") {
		cells := row.MustElements("td")
		key := cells[0].MustProperty("textContent")
		if strings.HasPrefix(key.String(), "User Agent") {
			fmt.Printf("\t\t%s: %t\n\n", key, !strings.Contains(cells[1].MustProperty("textContent").String(), "HeadlessChrome/"))
		} else if strings.HasPrefix(key.String(), "Hairline Feature") {
			// Detects support for hidpi/retina hairlines, which are CSS borders with less than 1px in width, for being physically 1px on hidpi screens.
			// Not all the machine suppports it.
			continue
		} else {
			fmt.Printf("\t\t%s: %s\n\n", key, cells[1].MustProperty("textContent"))
		}
	}
	log.Printf("Report printed.")
}
