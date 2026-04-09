package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	text := bufio.NewScanner(os.Stdin)
	for text.Scan() {
		line := text.Text()
		result := strings.Fields(line) // turning it to slice
		if len(result) == 2 {
			inputfile, err := os.Open(result[0])
			if err != nil {
				fmt.Println("error opening the inputfile")
				return
			}
			defer inputfile.Close()
			outputfile, err := os.Create(result[1])
			if err != nil {
				fmt.Println("error creating the outputfile")
				return
			}
			defer outputfile.Close()
			// scanner := bufio.NewScanner(inputfile)
			//  writer := bufio.NewWriter(outputfile)
			//  for scanner.Scan(){
			// 	line := scanner.Text()
			// 	new := HexBin(line)
			// 	new =
			//  }

			break
		} else {
			fmt.Println("Enter Input.txt and output.txt")
		}
	}

	// if len(os.Args) != 3 {
	// 	fmt.Println("Error: incomplete argument")
	// 	fmt.Println("usage: go run . <input.txt> <output.txt>")
	// 	return
	// }
	// inputfile := os.Args[1]
	// outputfile := os.Args[2]

	// inputfile, err := os.Open("input.txt")
	// if err != nil {
	// 	fmt.Println("error opening the inputfile")
	// 	return
	// }
	// defer inputfile.Close()

	// outputfile, err := os.Create("output.txt")
	// if err != nil {
	// 	fmt.Println("error creating the outputfile")
	// 	return
	// }
	// defer outputfile.Close()

	// scanner := bufio.NewScanner(inputfile)
	// writer := bufio.NewWriter(outputfile)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// }
	// writer.Flush()

}
