/*
1.4 Dup program to print all the files of duplicate lines
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	file_map := make(map[string]map[string]int)

	for _, file := range files {
		// Reading the entire data of each file
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in reading file %v", err)
			continue
		}
		// Looping over each line of the file data
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if file_map[line] == nil {
				file_map[line] = make(map[string]int)
			}
			file_map[line][file]++
		}
	}

	// Output in Format#1
    fmt.Printf("OUTPUT IN FORMAT-1:\n")
	for line, count := range counts {
		fmt.Printf("%d\t%s\n\t\t%v\n", count, line, file_map[line])
	}

	// Output in Format#2
    fmt.Printf("\n\nOUTPUT IN FORMAT-2:\n")
	for line, count := range counts {
		fmt.Printf("%d\t%s\n", count, line)
		for file, fc := range file_map[line] {
			fmt.Printf("\t--> %d:%s\n", fc, file)
		}
		fmt.Printf("\n")
	}
}
