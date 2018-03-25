/*
1.8: Fetch program to add http:// if its missing
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
			fmt.Println(os.Stderr, "Error fetching url %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("URL: %s\n", url)
		_, err = io.Copy(os.Stdout, response.Body)
		defer response.Body.Close()
		if err != nil {
			fmt.Println(os.Stderr, "Error reading the response body: %v\n", url, err)
			os.Exit(1)
		}
	}
}
