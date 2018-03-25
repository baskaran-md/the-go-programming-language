/*
1.10: Fetch URL program with writing response to file
*/
package main

import (
	"fmt"
	"io"
	"net/http"
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
	for i, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		go fetch(i, url, ch)
	}

	fmt.Printf("Time\tBytes\tStatus\tURL\n")

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	xtime := time.Since(start).Seconds()
	fmt.Printf("Total Execution Time: %vs\n", xtime)
}

func fetch(index int, url string, ch chan<- string) {

	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching url %s: %v", url, err)
		return
	}

	file := "result_" + strconv.Itoa(index)
	f, err := os.Create(file)
	if err != nil {
		ch <- fmt.Sprintf("Error creating file: %v", file, err)
		return
	}
	defer f.Close()

	nbytes, err := io.Copy(f, response.Body)
	defer response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error reading the response body: %v", url, err)
		return
	}
	xtime := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f\t%d\t%s\t#%d:%s", xtime, nbytes, response.Status, index, url)
}
