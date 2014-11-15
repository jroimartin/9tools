package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var cols = flag.Int("cols", 80, "Maximum number of columns")

func main() {
	flag.Parse()

	var r io.Reader
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		r = file
	} else {
		r = os.Stdin
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		i := 0
		for {
			if len(line[i:]) < *cols {
				fmt.Println(line[i:])
				break
			}
			idx := strings.LastIndex(string(line[i:i+*cols]), " ")
			if idx == -1 {
				fmt.Println(line[i : i+*cols])
				i += *cols
			} else {
				fmt.Println(line[i : i+idx])
				i += idx + 1
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
