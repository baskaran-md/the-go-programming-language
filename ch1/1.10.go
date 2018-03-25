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
	for i, request_url := range os.Args[1:] {
		if !strings.HasPrefix(request_url, "http://") && !strings.HasPrefix(request_url, "https://") {
			request_url = "http://" + request_url
		}
		go fetch(i, request_url, ch)
	}

	fmt.Printf("Time\tBytes\tStatus\tURL\n")

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	xtime := time.Since(start).Seconds()
	fmt.Printf("Total Execution Time: %vs\n", xtime)
}

func fetch(index int, request_url string, ch chan<- string) {

	u, err := url.Parse(request_url)
	if err != nil {
		ch <- fmt.Sprintf("Error parsing url %s: %v", request_url, err)
		return
	}

	start := time.Now()
	response, err := http.Get(request_url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching url %s: %v", request_url, err)
		return
	}

	file := "Result_" + strconv.Itoa(index) + "_" + u.Hostname()
	f, err := os.Create(file)
	if err != nil {
		ch <- fmt.Sprintf("Error creating file: %v", file, err)
		return
	}
	defer f.Close()

	nbytes, err := io.Copy(f, response.Body)
	defer response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error reading the response body: %v", request_url, err)
		return
	}
	xtime := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f\t%d\t%s\t#%d:%s", xtime, nbytes, response.Status, index, request_url)
}
