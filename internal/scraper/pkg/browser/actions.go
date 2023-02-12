package browser

import (
	"github.com/go-rod/rod"
	"sync"
)

var (
	pool = rod.NewPagePool(3)
)

func PagePool(browser *rod.Browser, job func(func() *rod.Page)) {

	// We create a pool that will hold at most 3 pages which means the max concurrency is 3
	// Create a page if needed
	create := func() *rod.Page {
		// We use MustIncognito to isolate pages with each other
		return browser.MustPage()
	}
	// Run jobs concurrently
	wg := sync.WaitGroup{}
	for range "...." {
		wg.Add(1)
		go func() {
			defer wg.Done()
			job(create)

		}()
	}
	wg.Wait()
	// cleanup pool
	pool.Cleanup(func(p *rod.Page) { p.MustClose() })
}

func Job(create func() *rod.Page) {
	page := pool.Get(create)
	defer pool.Put(page)
	//page.MustNavigate(URLS[0])

}

// run a function in a goroutine and return a channel which will be closed when the function exits
func RunAsync(f func()) chan struct{} {
	done := make(chan struct{})
	go func() {
		f()

	}()
	return done
}
