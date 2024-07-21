package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var FileToOpen string
	var wpm int
	var ReadingSpeed time.Duration
	fmt.Print("what file would you like to open ? ")
	fmt.Scanln(&FileToOpen)
	fmt.Print("at what WPM (words per minute) ")
	fmt.Scanln(&wpm)
	ReadingSpeed = time.Duration(60000 / wpm)
	file, err := os.Open(FileToOpen)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Print(scanner.Text())
		fmt.Print("\n")
		time.Sleep(ReadingSpeed * time.Millisecond)
		fmt.Println("\033c")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
