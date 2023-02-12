package common

import (
	"bufio"
	"log"
	"os"
)

const (
	ROD_PROXY_PATH         = "127.0.0.1:8080"
	ROD_USER_DATA_DIR      = "/tmp/rod-data"
	SCREENSHOTS_PNG_PATH   = "/tmp/imgs"
	SCREENSHOTS_PDF_PATH   = "/tmp/pdfs"
	CHROME_EXTENSIONS_PATH = "/tmp/exts"
	RemoteBrowserURL       = "http://localhost:7320"
	PageCount              = 4
	EXT_CHROME_PATH        = "D:\\go-workspace\\toyscraper\\scraper\\pkg\\plugins\\cookie-editor-chrome-1.10.1"
)

var (
	HTML_TAGS       = readHtmlTags()
	HTML_PROPERTIES = readHTMLTagProperties()
	SVG_TAGS        = readSvgTags()
)

func readHTMLTagProperties() []string {
	var htmlTags []string
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "\\scraper\\common\\html5TagProperties.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		htmlTags = append(htmlTags, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return htmlTags
}
func readSvgTags() []string {
	var htmlTags []string
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "\\scraper\\common\\svgTags.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		htmlTags = append(htmlTags, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return htmlTags
}

func readHtmlTags() []string {
	var htmlTags []string
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "\\scraper\\common\\html5tags.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		htmlTags = append(htmlTags, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return htmlTags
}
