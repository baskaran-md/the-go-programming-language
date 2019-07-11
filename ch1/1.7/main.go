/*
1.7: Fetch URL and print in stdout using io.Copy
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching url %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("URL: %s\n", url)
		if _, err := io.Copy(os.Stdout, response.Body); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading the response body, %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
