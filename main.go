package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

// Wait Group
var wg sync.WaitGroup

// Mutex lock
var mu sync.Mutex

// Map
var result = make(map[string]int)

// Target struct
type Target struct {
	Url string `json:"url"`
	Method string `json:"method"`
}

// Config struct
type Config struct {
	TargetLists []Target `json:"targets"`
}

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
	var config Config;

	// Read file
	readFile, err := os.ReadFile("config.json")
	if err != nil {
		panic("Something went wrong! Please check error\n" + err.Error())
	}

	err = json.Unmarshal(readFile, &config)
	if err != nil {
		panic("Something went wrong! Please check error\n" + err.Error())
	}

	for _, tg := range config.TargetLists {
		wg.Add(1)
		go watchdog(tg.Url)
	}

	wg.Wait()

	for i, j := range result {
		fmt.Println("URL:", i)
		fmt.Println("Result:", j)
	}
}