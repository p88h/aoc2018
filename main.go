package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	Day5(scanner)
}
