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
func watchdog (tg Target) {
	defer wg.Done()

	// Make request
	req, err := http.NewRequest(tg.Method, tg.Url, nil)
	if err != nil {
		fmt.Println("Something went wrong! Please check a", tg)
		return
	}

	// Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		mu.Lock()
		result[tg.Method + " " + tg.Url] = 0
		mu.Unlock()
		return
	}
	defer res.Body.Close()

	// Put result
	mu.Lock()
	result[tg.Method + " " + tg.Url] = res.StatusCode
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

	// Unmarshal to json
	err = json.Unmarshal(readFile, &config)
	if err != nil {
		panic("Something went wrong! Please check error\n" + err.Error())
	}

	// Watchdog work
	for _, tg := range config.TargetLists {
		wg.Add(1)
		go watchdog(tg)
	}
	wg.Wait()

	// Output result
	for i, j := range result {
		fmt.Println("URL:", i)
		fmt.Println("Result:", j)
	}
}