/*
1.10: Fetch URL program with writing response to file
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// AlexaURL URL to fetch the million websites
	AlexaURL = "http://downloads.majesticseo.com/majestic_million.csv"

	// SampleAlexaURL Sample URL containing few URLs for testing
	SampleAlexaURL = "https://gist.githubusercontent.com/baskaran-md/0ca2b3b5dfce82d70f4d516ae439532b/raw/a0703a9940151c12e92fc5f3d5a372018ffb8258/majestic_million_sample.csv"
	// Timeout in seconds for fetching the URL (client side timeout)
	Timeout = 120
)

func main() {

	start := time.Now()
	ch := make(chan string)

	var exportToFile = flag.Bool("export", false, "Option to write results to file. Defaults to false")
	var limitURL = flag.Int("limit", 100, "Option to limit the number of URLs to fetch.")
	var dryRun = flag.Bool("dryrun", false, "Option to dryrun against Test URL having smaller list of URLs")

	flag.Parse()

	fmt.Printf("Write to file: %t\n", *exportToFile)
	fmt.Printf("URL limitURL to read: %d\n", *limitURL)
	var urlToUse string
	if *dryRun == false {
		urlToUse = AlexaURL
	} else {
		urlToUse = SampleAlexaURL
	}

	fmt.Printf("Fetching the list from AlexaURL: %v\n", urlToUse)
	response, err := http.Get(urlToUse)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching the URL list from AlexaURL")
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading the response from AlexaURL")
		os.Exit(1)
	}

	allRequestURLs := getURLList(string(data), *limitURL)

	//fmt.Printf("All Request URLS:\n%v\n", allRequestURLs)
	for i, requestURL := range allRequestURLs {
		if !strings.HasPrefix(requestURL, "http://") && !strings.HasPrefix(requestURL, "https://") {
			requestURL = "http://" + requestURL
		}
		go fetch(i, requestURL, ch, *exportToFile)
	}

	fmt.Printf("Time\tBytes\tStatus\tURL\n")

	for range allRequestURLs {
		fmt.Println(<-ch)
	}

	xtime := time.Since(start).Seconds()
	fmt.Printf("Total Execution Time: %vs\n", xtime)
}

func getURLList(data string, limitURL int) []string {
	var allRequestURLs []string
	for index, line := range strings.Split(data, "\n") {
		if index == 0 {
			// skip header
			continue
		}
		url := strings.Split(line, ",")[2]
		//fmt.Printf("Parsing Field: %v:%s\n", fields, fields[2])
		allRequestURLs = append(allRequestURLs, url)
		if index >= limitURL {
			break
		}
	}
	return allRequestURLs
}

func fetch(index int, requestURL string, ch chan<- string, exportToFile bool) {

	u, err := url.Parse(requestURL)
	if err != nil {
		ch <- fmt.Sprintf("-\t-\t-\t#%d:%s: Error Parsing URL :: %v", index, requestURL, err)
		return
	}

	start := time.Now()
	client := &http.Client{
		Timeout: time.Duration(Timeout) * time.Second,
	}
	response, err := client.Get(requestURL)
	if err != nil {
		ch <- fmt.Sprintf("-\t-\t-\t#%d:%s: Error Fetchig URL :: %v", index, requestURL, err)
		return
	}

	var nbytes int64
	if exportToFile == true {
		file := "Result_" + strconv.Itoa(index) + "_" + u.Hostname()
		nbytes = writeResponseToFile(file, response, ch)
	} else {
		nbytes, err = io.Copy(ioutil.Discard, response.Body)
		if err != nil {
			ch <- fmt.Sprintf("-\t-\t-\t#%d:%s: Error reading response.Body :: %v", index, requestURL, err)
			return
		}
	}
	//nbytes := len(response.ContentLength)

	xtime := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f\t%d\t%s\t#%d:%s", xtime, nbytes, response.Status, index, requestURL)
}

func writeResponseToFile(file string, response *http.Response, ch chan<- string) int64 {
	f, err := os.Create(file)
	if err != nil {
		ch <- fmt.Sprintf("-\t-\t-\t#%v: Error creating file :: %v", file, err)
		return 0
	}
	defer f.Close()

	nbytes, err := io.Copy(f, response.Body)
	defer response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("-\t-\t-\t#%v: Error reading response.Body: %v", file, err)
		return 0
	}
	return nbytes
}
