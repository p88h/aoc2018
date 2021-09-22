package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	Day7(scanner)
}
