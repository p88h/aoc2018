package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	Day10(scanner)
}
