package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Wait Group
var wg sync.WaitGroup

// Mutex lock
var mu sync.Mutex

// Map
var result = make(map[string]int)

// Send request to url
func watchdog (url string) {
	defer wg.Done()

	res, err := http.Get(url)
	if err != nil {
		mu.Lock()
		result[url] = 0
		mu.Unlock()
		return
	}
	defer res.Body.Close()

	mu.Lock()
	result[url] = res.StatusCode
	mu.Unlock()
}

// Main program
func main() {
	// Url list for check
	urlLists := []string{"https://google.com", "https://this-web-does-not-exist.com", "https://github.com"}

	for _, url := range urlLists {
		wg.Add(1)
		go watchdog(url)
	}

	wg.Wait()

	for i, j := range result {
		fmt.Println("URL:", i)
		fmt.Println("Result:", j)
	}
}