package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	time.Sleep(time.Millisecond)
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func MyPrint(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
	time.Sleep(time.Millisecond)
}

// ******** Exercise! *********
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
var crawlUrl = map[string]bool{}
var mutex sync.Mutex

func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	defer wg.Done()
	mutex.Lock()
	_, ok := crawlUrl[url]
	if !ok {
		crawlUrl[url] = true
	}
	mutex.Unlock()

	if depth <= 0 || ok {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	//fmt.Println(urls)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg)
	}
	return
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func main() {
	// We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock.
	// We can also use defer to ensure the mutex will be unlocked as in the Value method.
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 10; i++ {
		go c.Inc("somekey") // 10 goroutines start at the same time
	}

	time.Sleep(2 * time.Millisecond) // we'll wait for 2ms as the Value method also waits for the lock (what's the priority of the lock acquirement)
	fmt.Println(c.Value("somekey"))
	fmt.Println("---")

	// Use sync.WaitGroup to wait for all goroutines finished.
	// (Go没有像Python中多线程的join那样直接的方法，我们需要手动设置一个计数器（即sync.WaitGroup），一般会在goroutine外计数加一，
	// 而在goroutine内使用`defer wg.Done()`，即函数返回之后计数减一)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go MyPrint(i, &wg)
	}
	wg.Wait()
	fmt.Println("---")

	// ******** Exercise! *********
	var wg2 sync.WaitGroup
	wg2.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &wg2)
	wg2.Wait()
}
