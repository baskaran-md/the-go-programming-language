/*
* Examples for file reading
 */
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
* Method 1: Read the content of the file line by line
* using bufio scanner till it reaches EOF
 */
func fileReadMethod1(fileList []string, count map[string]int) {

	fmt.Printf("== FILE READ: METHOD 1 ==\n")
	for _, file := range fileList {
		fmt.Printf("Reading file: %s\n", file)
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open the file: %v\n", err)
			continue
		}
		defer f.Close()
		input := bufio.NewScanner(f)
		for input.Scan() {
			line := input.Text()
			fmt.Printf("Line: %s\n", line)
			count[line]++
		}
	}
}

/*
* Method 2: Read the entire content of the file in one gulp
* and parsing them by using strings.Splint function
 */
func fileReadMethod2(fileList []string, count map[string]int) {

	fmt.Printf("== FILE READ: METHOD 2 ==\n")
	for _, file := range fileList {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
			continue
		}

		for index, line := range strings.Split(string(data), "\n") {
			fmt.Printf("Line#%d: %v\n", index, line)
			count[line]++
		}

	}

}

/*
* Reset the counter map by deleting all existing keys
 */
func resetCount(count map[string]int) {
	for key := range count {
		delete(count, key)
	}
}

/*
* The Main function begins
 */
func main() {
	fmt.Printf("Counting duplicate words:\n")

	fileList := os.Args[1:]
	count := make(map[string]int)

	fileReadMethod1(fileList, count)
	fmt.Printf("Result1:\n%v\n\n", count)

	resetCount(count)

	fileReadMethod2(fileList, count)
	fmt.Printf("Result2:\n%v\n\n", count)
}
