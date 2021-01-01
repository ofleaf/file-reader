package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./text.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(file)
	for {
		var buffer bytes.Buffer
		var l []byte
		var isPrefix bool
		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)

			if !isPrefix {
				break
			}

			if err != nil {
				break
			}

		}
		if err == io.EOF {
			break
		}

		line := buffer.String()

		fmt.Printf("Read %d characters\n", len(line))
		fmt.Println(line)
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return
}
