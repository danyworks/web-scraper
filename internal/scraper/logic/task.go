package logic

import "time"

type ScrapingTask struct {
	Name          string        `json:"name"`
	CrawlsChained CrawlingChain `json:"crawls_chain"`
	Timeout       time.Duration `json:"timeout"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	Interval      time.Duration `json:"interval"`
}
