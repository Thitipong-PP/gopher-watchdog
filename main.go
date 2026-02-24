package main

import (
	"fmt"
	"net/http"
)

func main() {
	urlLists := []string{"https://google.com", "https://this-web-does-not-exist.com", "https://github.com"}

	for _, url := range urlLists {
		res, err := http.Get(url)

		if err != nil {
			fmt.Println("Cannot get response from ", url)
			fmt.Println("Error :", err.Error())
			continue
		}

		fmt.Println("URL:", url)
		fmt.Println("Response code:", res.StatusCode)
		res.Body.Close()
	}
}