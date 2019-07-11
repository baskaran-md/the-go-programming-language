/*
1.10: Fetch URL program with writing response to file
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	ch := make(chan string)
	if len(os.Args) <= 1 {
		fmt.Println("URL input is required as command line args")
		os.Exit(1)
	}
	for i, requestURL := range os.Args[1:] {
		if !strings.HasPrefix(requestURL, "http://") && !strings.HasPrefix(requestURL, "https://") {
			requestURL = "http://" + requestURL
		}
		go fetch(i, requestURL, ch)
	}

	fmt.Printf("Time\tBytes\tStatus\tURL\n")

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	xtime := time.Since(start).Seconds()
	fmt.Printf("Total Execution Time: %vs\n", xtime)
}

func fetch(index int, requestURL string, ch chan<- string) {

	u, err := url.Parse(requestURL)
	if err != nil {
		ch <- fmt.Sprintf("Error parsing url %s: %v", requestURL, err)
		return
	}

	start := time.Now()
	response, err := http.Get(requestURL)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching url %s: %v", requestURL, err)
		return
	}

	file := "Result_" + strconv.Itoa(index) + "_" + u.Hostname()
	f, err := os.Create(file)
	if err != nil {
		ch <- fmt.Sprintf("Error creating file %s: %v", file, err)
		return
	}
	defer f.Close()

	nBytes, err := io.Copy(f, response.Body)
	defer response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error reading the response body, %s: %v", requestURL, err)
		return
	}
	xTime := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f\t%d\t%s\t#%d:%s", xTime, nBytes, response.Status, index, requestURL)
}
