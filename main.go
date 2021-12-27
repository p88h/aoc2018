package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

type Aoc struct{}

func Invoke(name string, args ...interface{}) {
	aoc := Aoc{}
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(aoc).MethodByName(name).Call(inputs)
}

func main() {
	files, _ := filepath.Glob("*.txt")
	args := os.Args[1:]
	dayno := int64(0)
	for f := range files {
		day, _ := strconv.ParseInt(files[f][3:len(files[f])-4], 10, 0)
		if day > dayno {
			dayno = day
		}
	}
	if len(args) > 0 {
		dayno, _ = strconv.ParseInt(args[0], 10, 0)
	}
	file, err := os.Open(fmt.Sprintf("day%d.txt", dayno))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	Invoke(fmt.Sprintf("Day%d", dayno), scanner)
}
