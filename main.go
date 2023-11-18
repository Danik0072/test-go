package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	file, err := os.Open("links.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)
		requestURL := scanner.Text()
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(requestURL)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}

			fmt.Printf("%s - %d\n", requestURL, res.StatusCode)
		}(requestURL)
	}
	wg.Wait()
}
