package logic

import (
	"context"
	"errors"
	"github.com/go-rod/rod"
	"log"
	"time"
)

type CrawlTargetTypes int

const (
	TargetDiv CrawlTargetTypes = iota
	TargetSpan
	TargetButton
	TargetInput
	TargetLink
	TargetImage
	TargetTable
	TargetForm
	TargetSelect
	TargetScript
)

func NewCrawl(name, ttype, txpath string, timeout time.Duration, targs ...string) *Step {
	return &Step{
		ID:           "",
		Name:         name,
		TargetType:   ttype,
		TargetXPath:  txpath,
		TargetAction: "",
		InputArgs:    targs,
		Timeout:      timeout,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}

func PerformCrawlWithCancel(page *rod.Page, url string) {
	ctx, cancel := context.WithCancel(context.Background())
	pageWithCancel := page.Context(ctx)
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()
	pageWithCancel.MustNavigate(url)
}

func PerformCrawlDetectTimeout(page *rod.Page, url string, timeout time.Duration) {
	err := rod.Try(func() {
		page.Timeout(timeout).MustNavigate(url)
	})
	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("timeout error")
	} else if err != nil {
		log.Println("other types of error")
	}
}
