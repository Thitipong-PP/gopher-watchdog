package main

import (
	"fmt"
	"net/http"
)

// Send request to url
func watchdog (url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Cannot get response from ", url)
		fmt.Println("Error :", err.Error())
		return
	}

	fmt.Println("URL:", url)
	fmt.Println("Response code:", res.StatusCode)
	res.Body.Close()
}

// Main program
func main() {
	// Url list for check
	urlLists := []string{"https://google.com", "https://this-web-does-not-exist.com", "https://github.com"}

	for _, url := range urlLists {
		watchdog(url)
	}
}