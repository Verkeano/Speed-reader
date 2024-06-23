package main

import (
	"bufio"
	"time"
	"fmt"
	"os"
	"log"
)

func main() {
	var FileToOpen string
	var wpm time.Duration
	fmt.Print("what file would you like to open ? ")
	fmt.Scanln(&FileToOpen)
	fmt.Print("at what speed (milliseconds between each word (havent done the math yet for words per minute))")
	fmt.Scanln(&wpm)
	file, err := os.Open(FileToOpen)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Print(scanner.Text())
		time.Sleep(wpm * time.Millisecond)
		fmt.Println("\033c")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
