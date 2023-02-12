package logic

import (
	"github.com/go-rod/rod"
	"time"
)

type ListOfCrawls []Step

type CrawlingChain struct {
	Crawls    ListOfCrawls  `json:"crawls"`
	Timeout   time.Duration `json:"timeout"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (l *ListOfCrawls) Add(element Step) {
	*l = append(*l, element)
}

func (l *ListOfCrawls) Get(index int) Step {
	return (*l)[index]
}
func (l *ListOfCrawls) Len() int {
	return len(*l)
}
func (l *ListOfCrawls) Process(page *rod.Page) {
	for _, crawl := range *l {
		crawl.PerformCrawl(page)
	}
}
