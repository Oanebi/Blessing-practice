package main

import (
	"bufio"
	"fmt"
	"os"
	"practice-goreload/processor"
	"strings"
)

func main() {
	text := bufio.NewScanner(os.Stdin)
	for text.Scan() {
		line := text.Text()
		command := strings.Fields(line)
		if len(command) == 2 {
			inputfile, err := os.Open(command[0])
			if err != nil {
				fmt.Println("error opening the file", err)
			}
			defer inputfile.Close()

			outputfile, err := os.Create(command[1])
			if err != nil {
				fmt.Println("error creating file", err)
			}
			defer outputfile.Close()
			scanner := bufio.NewScanner(inputfile)
			writer := bufio.NewWriter(outputfile)

			for scanner.Scan() {
				result := scanner.Text()

				line := processor.Compiler(result)

				writer.WriteString(line + "\n")
			}
			writer.Flush()

		} else {
			fmt.Println("incomplete command")
		}
	}
}
