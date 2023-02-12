package logic

import (
	"github.com/daniyalibrahim/notybot/internal/scraper/common"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
	"log"
	"time"
)

type Step struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	TargetType   string        `json:"target_type"`
	TargetXPath  string        `json:"target_xpath"`
	TargetAction string        `json:"target_action"`
	InputArgs    []string      `json:"input_arg"`
	Timeout      time.Duration `json:"timeout"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

func (c *Step) New(timeout time.Duration, name, ttype, xPath, action string, args ...string) *Step {
	return &Step{
		Name:         name,
		TargetType:   ttype,
		TargetXPath:  xPath,
		TargetAction: action,
		InputArgs:    args,
		Timeout:      timeout,
	}
}

func (C *Step) PerformCrawl(page *rod.Page) {
	log.Print("Performing crawl...")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in PerformCrawl", r)
		}
	}()
	el, err := page.Timeout(C.Timeout).ElementX(C.TargetXPath)
	common.HandleError(err)
	switch C.TargetType {
	case "button":
		log.Printf("clicking %s", C.Name)
		err = el.Timeout(6*time.Second).Click(proto.InputMouseButtonLeft, 2)
		common.HandleError(err)
	case "input":
		log.Printf("Input %v in %v", C.InputArgs[0], C.Name)
		el.Timeout(6 * time.Second).MustInput(C.InputArgs[0])
		common.HandleError(err)
	case "select":
		log.Printf("Select option %v in %v\n", C.InputArgs[0], C.Name)
		err := el.Timeout(6*time.Second).Select(C.InputArgs, true, rod.SelectorTypeRegex)
		common.HandleError(err)
		//el.Select(C.CrawlTargetArgs, false, rod.SelectorTypeCSSSector)
	case "file":
		el.MustSetFiles("./tmp/input-example.txt")
	case "text":
		text, err := el.Text()
		common.HandleError(err)
		log.Printf("Text %v in %v\n", text, C.Name)
	}
	switch C.TargetAction {
	case "search":
		time.Sleep(2 * time.Second)
		el.MustType(input.ArrowDown, input.Enter)
	case "switch":
		el.MustType(input.Tab)
	case "fbclear":
		err = el.Timeout(6*time.Second).Click(proto.InputMouseButtonLeft, 2)
		common.HandleError(err)
	}
	log.Print("Step performed")

}
