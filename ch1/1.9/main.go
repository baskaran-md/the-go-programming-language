/*
1.9: Fetch program to print http status
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching url %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("URL: %s\nStatus: %s\n", url, response.Status)
		if _, err := io.Copy(os.Stdout, response.Body); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading the response body %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
